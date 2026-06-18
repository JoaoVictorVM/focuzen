// Minimal type declarations for the YouTube IFrame Player API
// (https://developers.google.com/youtube/iframe_api_reference). We declare only
// what we use instead of depending on @types/youtube.
declare namespace YT {
  enum PlayerState {
    UNSTARTED = -1,
    ENDED = 0,
    PLAYING = 1,
    PAUSED = 2,
    BUFFERING = 3,
    CUED = 5,
  }

  type PlayerEvent = { target: Player };
  type OnStateChangeEvent = { target: Player; data: PlayerState };

  type PlayerVars = {
    autoplay?: 0 | 1;
    controls?: 0 | 1;
    disablekb?: 0 | 1;
    playsinline?: 0 | 1;
  };

  type PlayerOptions = {
    videoId?: string;
    playerVars?: PlayerVars;
    events?: {
      onReady?: (event: PlayerEvent) => void;
      onStateChange?: (event: OnStateChangeEvent) => void;
    };
  };

  class Player {
    constructor(element: HTMLElement, options: PlayerOptions);
    playVideo(): void;
    pauseVideo(): void;
    loadVideoById(videoId: string): void;
    setVolume(volume: number): void;
    getVolume(): number;
    mute(): void;
    unMute(): void;
    getPlayerState(): PlayerState;
    destroy(): void;
  }
}

interface Window {
  YT?: typeof YT;
  onYouTubeIframeAPIReady?: () => void;
}
