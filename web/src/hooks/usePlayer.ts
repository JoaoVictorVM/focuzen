import { useCallback, useEffect, useRef, useState } from 'react';

import { loadYouTubeIframeAPI } from '../lib/youtube-iframe';
import type { Video } from '../types/youtube';

const DEFAULT_VOLUME = 70;

// usePlayer owns a single hidden YouTube IFrame player. The video is hidden via
// CSS (see AudioPlayer) so only the audio is heard. It lazily creates the player
// on the first play and reuses it (loadVideoById) afterwards.
export function usePlayer() {
  const containerRef = useRef<HTMLDivElement>(null);
  const playerRef = useRef<YT.Player | null>(null);
  const [current, setCurrent] = useState<Video | null>(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [volume, setVolume] = useState(DEFAULT_VOLUME);

  // Keep the latest volume readable from onReady without re-creating the player.
  const volumeRef = useRef(volume);
  volumeRef.current = volume;

  useEffect(() => {
    return () => {
      playerRef.current?.destroy();
      playerRef.current = null;
    };
  }, []);

  const play = useCallback((video: Video) => {
    setCurrent(video);

    if (playerRef.current) {
      playerRef.current.loadVideoById(video.id);
      return;
    }

    void loadYouTubeIframeAPI().then((api) => {
      if (!containerRef.current || playerRef.current) {
        return;
      }
      playerRef.current = new api.Player(containerRef.current, {
        videoId: video.id,
        playerVars: { autoplay: 1, controls: 0, disablekb: 1, playsinline: 1 },
        events: {
          onReady: (event) => {
            event.target.setVolume(volumeRef.current);
            event.target.playVideo();
          },
          onStateChange: (event) => {
            setIsPlaying(event.data === api.PlayerState.PLAYING);
          },
        },
      });
    });
  }, []);

  const togglePlay = useCallback(() => {
    const player = playerRef.current;
    if (!player) {
      return;
    }
    if (isPlaying) {
      player.pauseVideo();
    } else {
      player.playVideo();
    }
  }, [isPlaying]);

  const changeVolume = useCallback((value: number) => {
    setVolume(value);
    playerRef.current?.setVolume(value);
  }, []);

  return { containerRef, current, isPlaying, volume, play, togglePlay, changeVolume } as const;
}
