import { useEffect, useState } from 'react';

// useNow is the single source of the current time, re-rendering consumers on a
// fixed interval so the clock and date stay in sync.
export function useNow(intervalMs = 1000) {
  const [now, setNow] = useState(() => new Date());

  useEffect(() => {
    const id = setInterval(() => setNow(new Date()), intervalMs);
    return () => clearInterval(id);
  }, [intervalMs]);

  return now;
}
