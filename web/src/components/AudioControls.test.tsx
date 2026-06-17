import { fireEvent, render, screen } from '@testing-library/react';
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
    hasNext: true,
    onTogglePlay: vi.fn(),
    onNext: vi.fn(),
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
        hasNext={false}
        onTogglePlay={vi.fn()}
        onNext={vi.fn()}
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

  it('reports volume changes', () => {
    const props = renderControls();
    fireEvent.change(screen.getByRole('slider', { name: 'Volume' }), {
      target: { value: '40' },
    });
    expect(props.onVolumeChange).toHaveBeenCalledWith(40);
  });
});
