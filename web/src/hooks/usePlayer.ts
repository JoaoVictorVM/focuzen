import { useCallback, useEffect, useRef, useState } from 'react';

import { loadYouTubeIframeAPI } from '../lib/youtube-iframe';
import type { Video } from '../types/youtube';

const DEFAULT_VOLUME = 70;

// usePlayer owns a single hidden YouTube IFrame player plus a local play queue.
// "Next" advances within this queue (the search results), not via YouTube
// recommendations — the API no longer exposes related videos (see ADR-0006).
export function usePlayer() {
  const containerRef = useRef<HTMLDivElement>(null);
  const playerRef = useRef<YT.Player | null>(null);
  const [queue, setQueue] = useState<Video[]>([]);
  const [index, setIndex] = useState(0);
  const [isPlaying, setIsPlaying] = useState(false);
  const [volume, setVolume] = useState(DEFAULT_VOLUME);
  const [repeat, setRepeat] = useState(false);

  const current = queue[index] ?? null;
  const hasNext = index < queue.length - 1;
  const hasPrevious = index > 0;

  // Keep the latest volume readable from onReady without re-creating the player.
  const volumeRef = useRef(volume);
  volumeRef.current = volume;

  // Read inside the once-created onStateChange handler to decide what to do when
  // a track ends.
  const repeatRef = useRef(repeat);
  repeatRef.current = repeat;
  const currentIdRef = useRef<string | null>(current?.id ?? null);
  currentIdRef.current = current?.id ?? null;

  // playNextRef always points at the current advance logic so the player's
  // onStateChange handler (created once) can auto-advance on ENDED.
  const playNextRef = useRef<() => void>(() => {});
  useEffect(() => {
    playNextRef.current = () => {
      if (index + 1 < queue.length) {
        const nextIndex = index + 1;
        setIndex(nextIndex);
        playerRef.current?.loadVideoById(queue[nextIndex].id);
      }
    };
  }, [index, queue]);

  useEffect(() => {
    return () => {
      playerRef.current?.destroy();
      playerRef.current = null;
    };
  }, []);

  const loadOrCreate = useCallback((videoId: string) => {
    if (playerRef.current) {
      playerRef.current.loadVideoById(videoId);
      return;
    }

    void loadYouTubeIframeAPI().then((api) => {
      if (!containerRef.current || playerRef.current) {
        return;
      }
      playerRef.current = new api.Player(containerRef.current, {
        videoId,
        playerVars: { autoplay: 1, controls: 0, disablekb: 1, playsinline: 1 },
        events: {
          onReady: (event) => {
            event.target.setVolume(volumeRef.current);
            event.target.playVideo();
          },
          onStateChange: (event) => {
            setIsPlaying(event.data === api.PlayerState.PLAYING);
            if (event.data === api.PlayerState.ENDED) {
              if (repeatRef.current && currentIdRef.current) {
                event.target.loadVideoById(currentIdRef.current);
              } else {
                playNextRef.current();
              }
            }
          },
        },
      });
    });
  }, []);

  const play = useCallback(
    (video: Video, list?: Video[]) => {
      const nextQueue = list && list.length > 0 ? list : [video];
      const found = nextQueue.findIndex((item) => item.id === video.id);
      const startIndex = found >= 0 ? found : 0;

      setQueue(nextQueue);
      setIndex(startIndex);
      loadOrCreate(nextQueue[startIndex].id);
    },
    [loadOrCreate],
  );

  const next = useCallback(() => {
    playNextRef.current();
  }, []);

  const previous = useCallback(() => {
    if (index > 0) {
      const previousIndex = index - 1;
      setIndex(previousIndex);
      playerRef.current?.loadVideoById(queue[previousIndex].id);
    }
  }, [index, queue]);

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

  const toggleRepeat = useCallback(() => {
    setRepeat((value) => !value);
  }, []);

  return {
    containerRef,
    current,
    isPlaying,
    volume,
    repeat,
    hasNext,
    hasPrevious,
    play,
    next,
    previous,
    togglePlay,
    toggleRepeat,
    changeVolume,
  } as const;
}
