import { useRef } from 'react';
import { useTranslation } from 'react-i18next';

const BAR_COUNT = 7;
const KEYBOARD_STEP = 5;

type VolumeControlProps = {
  volume: number;
  onChange: (value: number) => void;
};

// Staircase volume control: bars grow left→right; bars up to the current volume
// are highlighted. Drag across it with the pointer, or use the arrow keys.
export function VolumeControl({ volume, onChange }: VolumeControlProps) {
  const { t } = useTranslation();
  const trackRef = useRef<HTMLDivElement>(null);
  const dragging = useRef(false);

  const setFromClientX = (clientX: number) => {
    const track = trackRef.current;
    if (!track) {
      return;
    }
    const { left, width } = track.getBoundingClientRect();
    const ratio = (clientX - left) / width;
    onChange(Math.round(Math.min(1, Math.max(0, ratio)) * 100));
  };

  const activeBars = Math.round((volume / 100) * BAR_COUNT);

  return (
    <div
      ref={trackRef}
      role="slider"
      tabIndex={0}
      aria-label={t('player.volume')}
      aria-valuemin={0}
      aria-valuemax={100}
      aria-valuenow={volume}
      onPointerDown={(event) => {
        dragging.current = true;
        event.currentTarget.setPointerCapture(event.pointerId);
        setFromClientX(event.clientX);
      }}
      onPointerMove={(event) => {
        if (dragging.current) {
          setFromClientX(event.clientX);
        }
      }}
      onPointerUp={(event) => {
        dragging.current = false;
        event.currentTarget.releasePointerCapture(event.pointerId);
      }}
      onKeyDown={(event) => {
        if (event.key === 'ArrowRight' || event.key === 'ArrowUp') {
          event.preventDefault();
          onChange(Math.min(100, volume + KEYBOARD_STEP));
        } else if (event.key === 'ArrowLeft' || event.key === 'ArrowDown') {
          event.preventDefault();
          onChange(Math.max(0, volume - KEYBOARD_STEP));
        }
      }}
      className="flex h-5 cursor-pointer touch-none select-none items-end gap-[3px] rounded focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-primary"
    >
      {Array.from({ length: BAR_COUNT }, (_, index) => {
        const heightPercent = 30 + (index / (BAR_COUNT - 1)) * 70;
        const active = index < activeBars;
        return (
          <span
            key={index}
            style={{ height: `${heightPercent}%` }}
            className={`w-[5px] rounded-sm border border-foreground transition-opacity ${
              active ? 'bg-foreground' : 'bg-transparent opacity-40'
            }`}
          />
        );
      })}
    </div>
  );
}
