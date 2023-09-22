import Controls from '@/components/controls'
import Header from '@/components/header'
import HeightDisplay from '@/components/height-display'
import Image from 'next/image'

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-start p-24">
      <Header />
      <div className="my-24">
        <HeightDisplay />
      </div>
      <Controls />
    </main>
  )
}
