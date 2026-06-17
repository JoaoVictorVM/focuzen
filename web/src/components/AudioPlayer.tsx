import type { RefObject } from 'react';

type AudioPlayerProps = {
  containerRef: RefObject<HTMLDivElement | null>;
};

// Mount point for the YouTube IFrame player. The API replaces the inner div with
// an iframe; the wrapper keeps it off-screen and inert so only the audio plays.
export function AudioPlayer({ containerRef }: AudioPlayerProps) {
  return (
    <div
      aria-hidden="true"
      className="pointer-events-none fixed bottom-0 left-0 -z-10 h-1 w-1 overflow-hidden opacity-0"
    >
      <div ref={containerRef} />
    </div>
  );
}
