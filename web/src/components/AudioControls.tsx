import { useTranslation } from 'react-i18next';

import type { Video } from '../types/youtube';
import { PauseIcon, PlayIcon, VolumeIcon } from './icons';

type AudioControlsProps = {
  current: Video | null;
  isPlaying: boolean;
  volume: number;
  onTogglePlay: () => void;
  onVolumeChange: (value: number) => void;
};

export function AudioControls({
  current,
  isPlaying,
  volume,
  onTogglePlay,
  onVolumeChange,
}: AudioControlsProps) {
  const { t } = useTranslation();

  if (!current) {
    return null;
  }

  return (
    <div className="mt-10 flex items-center justify-center gap-4">
      <button
        type="button"
        onClick={onTogglePlay}
        aria-label={isPlaying ? t('player.pause') : t('player.play')}
        aria-pressed={isPlaying}
        className="rounded-full p-2 text-foreground transition-colors hover:text-primary"
      >
        {isPlaying ? (
          <PauseIcon size={20} strokeWidth={1.5} />
        ) : (
          <PlayIcon size={20} strokeWidth={1.5} />
        )}
      </button>

      <span className="max-w-[16rem] truncate text-sm text-muted" title={current.title}>
        {current.title}
      </span>

      <span className="flex items-center gap-2">
        <VolumeIcon size={16} strokeWidth={1.5} className="text-muted" />
        <input
          type="range"
          min={0}
          max={100}
          value={volume}
          onChange={(event) => onVolumeChange(Number(event.target.value))}
          aria-label={t('player.volume')}
          className="h-1 w-24 cursor-pointer accent-primary"
        />
      </span>
    </div>
  );
}
