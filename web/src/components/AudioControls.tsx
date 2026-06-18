import { useTranslation } from 'react-i18next';

import type { Video } from '../types/youtube';
import {
  MaximizeIcon,
  MinimizeIcon,
  MuteIcon,
  NextIcon,
  PauseIcon,
  PlayIcon,
  PreviousIcon,
  RepeatIcon,
  VolumeIcon,
} from './icons';
import { Tooltip } from './Tooltip';
import { VolumeControl } from './VolumeControl';

type AudioControlsProps = {
  current: Video | null;
  isPlaying: boolean;
  volume: number;
  repeat: boolean;
  muted: boolean;
  hasNext: boolean;
  hasPrevious: boolean;
  isFullscreen: boolean;
  onTogglePlay: () => void;
  onNext: () => void;
  onPrevious: () => void;
  onToggleRepeat: () => void;
  onToggleMute: () => void;
  onToggleFullscreen: () => void;
  onVolumeChange: (value: number) => void;
};

export function AudioControls({
  current,
  isPlaying,
  volume,
  repeat,
  muted,
  hasNext,
  hasPrevious,
  isFullscreen,
  onTogglePlay,
  onNext,
  onPrevious,
  onToggleRepeat,
  onToggleMute,
  onToggleFullscreen,
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
          onClick={onToggleFullscreen}
          aria-label={isFullscreen ? t('player.exitFullscreen') : t('player.enterFullscreen')}
          aria-pressed={isFullscreen}
          className="rounded-full p-2 text-foreground transition-colors hover:text-primary"
        >
          {isFullscreen ? (
            <MinimizeIcon size={18} strokeWidth={1.5} />
          ) : (
            <MaximizeIcon size={18} strokeWidth={1.5} />
          )}
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

        <button
          type="button"
          onClick={onToggleMute}
          aria-label={muted ? t('player.unmute') : t('player.mute')}
          aria-pressed={muted}
          className="rounded-full p-2 text-foreground transition-colors hover:text-primary"
        >
          {muted ? (
            <MuteIcon size={18} strokeWidth={1.5} />
          ) : (
            <VolumeIcon size={18} strokeWidth={1.5} />
          )}
        </button>

        <VolumeControl volume={volume} onChange={onVolumeChange} />
      </div>
    </div>
  );
}
