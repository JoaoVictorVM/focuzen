import { useCallback, useEffect, useState } from 'react';

// useFullscreen toggles the Fullscreen API on the document root, which fills the
// monitor the browser window is on (hiding the browser chrome). It tracks the
// real state via the fullscreenchange event so leaving with Esc stays in sync.
export function useFullscreen() {
  const [isFullscreen, setIsFullscreen] = useState(() => Boolean(document.fullscreenElement));

  useEffect(() => {
    const onChange = () => setIsFullscreen(Boolean(document.fullscreenElement));
    document.addEventListener('fullscreenchange', onChange);
    return () => document.removeEventListener('fullscreenchange', onChange);
  }, []);

  const toggleFullscreen = useCallback(() => {
    if (document.fullscreenElement) {
      void document.exitFullscreen().catch(() => {});
    } else {
      void document.documentElement.requestFullscreen().catch(() => {});
    }
  }, []);

  return { isFullscreen, toggleFullscreen } as const;
}
