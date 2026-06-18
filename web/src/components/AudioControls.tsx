import { useTranslation } from 'react-i18next';

import type { Video } from '../types/youtube';
import { NextIcon, PauseIcon, PlayIcon, PreviousIcon, RepeatIcon, VolumeIcon } from './icons';
import { Tooltip } from './Tooltip';

type AudioControlsProps = {
  current: Video | null;
  isPlaying: boolean;
  volume: number;
  repeat: boolean;
  hasNext: boolean;
  hasPrevious: boolean;
  onTogglePlay: () => void;
  onNext: () => void;
  onPrevious: () => void;
  onToggleRepeat: () => void;
  onVolumeChange: (value: number) => void;
};

export function AudioControls({
  current,
  isPlaying,
  volume,
  repeat,
  hasNext,
  hasPrevious,
  onTogglePlay,
  onNext,
  onPrevious,
  onToggleRepeat,
  onVolumeChange,
}: AudioControlsProps) {
  const { t } = useTranslation();

  if (!current) {
    return null;
  }

  return (
    <div className="mt-10 flex flex-col items-center gap-4">
      <span className="max-w-[20rem] truncate text-sm text-muted" title={current.title}>
        {current.title}
      </span>

      <div className="flex items-center gap-4">
        <button
          type="button"
          onClick={onPrevious}
          disabled={!hasPrevious}
          aria-label={t('player.previous')}
          className="rounded-full p-2 text-foreground transition-colors hover:text-primary disabled:cursor-not-allowed disabled:opacity-40"
        >
          <PreviousIcon size={18} strokeWidth={1.5} />
        </button>

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

        <button
          type="button"
          onClick={onNext}
          disabled={!hasNext}
          aria-label={t('player.next')}
          className="rounded-full p-2 text-foreground transition-colors hover:text-primary disabled:cursor-not-allowed disabled:opacity-40"
        >
          <NextIcon size={18} strokeWidth={1.5} />
        </button>

        <Tooltip label={repeat ? t('player.repeatOn') : t('player.repeatOff')}>
          <button
            type="button"
            onClick={onToggleRepeat}
            aria-label={t('player.repeat')}
            aria-pressed={repeat}
            className={`rounded-full p-2 transition-colors ${
              repeat ? 'text-primary' : 'text-foreground hover:text-primary'
            }`}
          >
            <RepeatIcon size={18} strokeWidth={1.5} />
          </button>
        </Tooltip>

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
    </div>
  );
}
