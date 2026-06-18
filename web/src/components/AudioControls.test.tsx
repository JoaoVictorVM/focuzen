import { cleanup, fireEvent, render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';

import i18n from '../i18n';
import type { Video } from '../types/youtube';
import { AudioControls } from './AudioControls';

const video: Video = {
  id: 'v1',
  title: 'Lofi Beats',
  channelTitle: 'Chillhop',
  thumbnailUrl: 'https://i.ytimg.com/vi/v1/mqdefault.jpg',
};

function renderControls(overrides: Partial<Parameters<typeof AudioControls>[0]> = {}) {
  const props = {
    current: video,
    isPlaying: false,
    volume: 70,
    repeat: false,
    hasNext: true,
    hasPrevious: true,
    isFullscreen: false,
    onTogglePlay: vi.fn(),
    onNext: vi.fn(),
    onPrevious: vi.fn(),
    onToggleRepeat: vi.fn(),
    onToggleFullscreen: vi.fn(),
    onVolumeChange: vi.fn(),
    ...overrides,
  };
  render(<AudioControls {...props} />);
  return props;
}

describe('AudioControls', () => {
  beforeEach(async () => {
    await i18n.changeLanguage('en');
  });

  it('renders nothing when there is no current track', () => {
    const { container } = render(
      <AudioControls
        current={null}
        isPlaying={false}
        volume={70}
        repeat={false}
        hasNext={false}
        hasPrevious={false}
        isFullscreen={false}
        onTogglePlay={vi.fn()}
        onNext={vi.fn()}
        onPrevious={vi.fn()}
        onToggleRepeat={vi.fn()}
        onToggleFullscreen={vi.fn()}
        onVolumeChange={vi.fn()}
      />,
    );
    expect(container).toBeEmptyDOMElement();
  });

  it('shows the title and toggles play', async () => {
    const props = renderControls({ isPlaying: false });

    expect(screen.getByText('Lofi Beats')).toBeInTheDocument();
    await userEvent.click(screen.getByRole('button', { name: 'Play' }));
    expect(props.onTogglePlay).toHaveBeenCalledOnce();
  });

  it('labels the button as Pause while playing', () => {
    renderControls({ isPlaying: true });
    expect(screen.getByRole('button', { name: 'Pause' })).toBeInTheDocument();
  });

  it('disables Next at the end of the queue', () => {
    renderControls({ hasNext: false });
    expect(screen.getByRole('button', { name: 'Next' })).toBeDisabled();
  });

  it('goes to the previous track and disables Previous at the start', async () => {
    const props = renderControls({ hasPrevious: true });
    await userEvent.click(screen.getByRole('button', { name: 'Previous' }));
    expect(props.onPrevious).toHaveBeenCalledOnce();
  });

  it('disables Previous at the start of the queue', () => {
    renderControls({ hasPrevious: false });
    expect(screen.getByRole('button', { name: 'Previous' })).toBeDisabled();
  });

  it('toggles repeat, reflects its pressed state and shows a state tooltip', async () => {
    const props = renderControls({ repeat: true });

    const repeatButton = screen.getByRole('button', { name: 'Repeat' });
    expect(repeatButton).toHaveAttribute('aria-pressed', 'true');
    expect(screen.getByRole('tooltip')).toHaveTextContent('Repeat on');

    await userEvent.click(repeatButton);
    expect(props.onToggleRepeat).toHaveBeenCalledOnce();
  });

  it('shows the off tooltip when repeat is inactive', () => {
    renderControls({ repeat: false });
    expect(screen.getByRole('tooltip')).toHaveTextContent('Repeat off');
  });

  it('toggles fullscreen and reflects its state', async () => {
    const enter = renderControls({ isFullscreen: false });
    await userEvent.click(screen.getByRole('button', { name: 'Fullscreen' }));
    expect(enter.onToggleFullscreen).toHaveBeenCalledOnce();

    cleanup();

    renderControls({ isFullscreen: true });
    expect(screen.getByRole('button', { name: 'Exit fullscreen' })).toBeInTheDocument();
  });

  it('reports volume changes', () => {
    const props = renderControls();
    fireEvent.change(screen.getByRole('slider', { name: 'Volume' }), {
      target: { value: '40' },
    });
    expect(props.onVolumeChange).toHaveBeenCalledWith(40);
  });
});
