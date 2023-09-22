import { useState, useEffect } from 'react'
import { GoPTSClient } from 'go-pts-client';

export default function usePTS() : [GoPTSClient|null] {
  const [ptsClient, setPtsClient] = useState<GoPTSClient|null>(null);

  useEffect(() => {
    setPtsClient(
      new GoPTSClient({ url: `ws://${window.location.host}/connect`, debugging: true })
    );
  }, []);

  return [ptsClient];
}