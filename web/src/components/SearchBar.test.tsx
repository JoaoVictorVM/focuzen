import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';

import i18n from '../i18n';
import { sampleResults } from '../test/msw/handlers';
import { SearchBar } from './SearchBar';

describe('SearchBar', () => {
  beforeEach(async () => {
    await i18n.changeLanguage('en');
  });

  it('searches and reports the selected video with the results queue', async () => {
    const onSelect = vi.fn();
    render(<SearchBar onSelect={onSelect} />);

    await userEvent.type(screen.getByRole('searchbox'), 'lofi');

    const firstResult = await screen.findByText('Lofi Beats');
    await userEvent.click(firstResult);

    expect(onSelect).toHaveBeenCalledWith(sampleResults[0], sampleResults);
  });

  it('shows the error message when the search fails', async () => {
    render(<SearchBar onSelect={vi.fn()} />);

    await userEvent.type(screen.getByRole('searchbox'), 'fail');

    expect(await screen.findByText('Search failed. Try again.')).toBeInTheDocument();
  });
});
