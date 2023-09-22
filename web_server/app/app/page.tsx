"use client";
import Controls from '@/components/controls'
import Header from '@/components/header'
import HeightDisplay from '@/components/height-display'
import usePTS from '@/hooks/use_pts'
import { GoPTSClient } from 'go-pts-client'
import { useEffect, useState } from 'react'

export default function Home() {
  const [client]: [GoPTSClient|null] = usePTS();
  const [height, setHeight] = useState(72);

  useEffect(() => {
    if (client) {
      client.subscribeChannel('control', (data: any) => setHeight(data.height));
    }

    return () => {
      client?.unsubscribeChannel('control', {});
      client?.close();
    }
  }, [client]);

  if (client == null) {
    return (
      <main className="flex min-h-screen flex-col items-center justify-start p-24">
        <Header />
        <div className="my-24">
          <div className="text-2xl text-center text-gray-600">
            Connecting to motor controller...
          </div>
        </div>
      </main>
    )
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-start p-24">
      <Header />
      <div className="my-24">
        <HeightDisplay value={height} />
      </div>
      <Controls
        onUp={() => client.send("control", { payload: "up" })}
        onDown={() => client.send("control", { payload: "down" })}
      />
    </main>
  )
}
