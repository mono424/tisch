package main

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

const UP_HEIGHT = 90
const DOWN_HEIGHT = 10

type MessageType byte

const (
	TypeAliveRequest      MessageType = 0x01
	TypeAliveResponse     MessageType = 0x02
	TypeSetHeightRequest  MessageType = 0x03
	TypeGetHeightRequest  MessageType = 0x04
	TypeGetHeightResponse MessageType = 0x05
	TypeStopRequest       MessageType = 0x06
	TypeGetStatusRequest  MessageType = 0x07
	TypeGetStatusResponse MessageType = 0x08
	TypeMoveUpRequest     MessageType = 0x0A
	TypeMoveDownRequest   MessageType = 0x0B
	TypeUpdateHeightEvent MessageType = 0x0C
)

type Message struct {
	Type  MessageType
	Value byte
}

func receiver(c chan<- Message, p *serial.Port) {
	message := make([]byte, 3)
	buf := make([]byte, 1)

	for {
		// shift bytes to left
		message[0] = message[1]
		message[1] = message[2]

		// read new byte
		_, err := p.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		// append new byte
		message[2] = buf[0]

		// checksum
		if message[0]+message[1] != message[2] {
			continue
		}

		c <- Message{Type: MessageType(message[0]), Value: message[1]}
	}
}

func sender(c <-chan Message, p *serial.Port) {
	message := make([]byte, 3)

	for {
		m := <-c
		message[0] = byte(m.Type)
		message[1] = m.Value
		message[2] = message[0] + message[1]
		_, err := p.Write(message)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func heightPercentageToCentimeters(percentage int) int {
	factor := float64(percentage) / 100.0
	return int(68.0 + 50.0*factor)
}

type MotorController struct {
	in  chan Message
	out chan Message
}

func New(config *serial.Config) *MotorController {
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	var outgoing chan Message = make(chan Message)
	var incoming chan Message = make(chan Message)
	go sender(outgoing, s)
	go receiver(incoming, s)

	return &MotorController{out: outgoing, in: incoming}
}

func (m *MotorController) togglePosition() {
	time.Sleep(2000 * time.Millisecond)
	m.out <- Message{Type: TypeGetHeightRequest}
	heightRaw := <-m.in
	height := int(heightRaw.Value)
	if height >= 100 {
		m.setPosition(DOWN_HEIGHT)
	} else if height < 100 && height >= 1 {
		m.setPosition(UP_HEIGHT)
	}
}

func (m *MotorController) setPosition(position int) {
	log.Println("Setting desk to", position, "percent height")
	height := heightPercentageToCentimeters(position)
	log.Println("This corresponds to", height, "cm height")
	m.out <- Message{Type: MessageType(TypeSetHeightRequest), Value: byte(height)}
}
