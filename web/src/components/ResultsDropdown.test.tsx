import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';

import i18n from '../i18n';
import { sampleResults } from '../test/msw/handlers';
import { ResultsDropdown } from './ResultsDropdown';

describe('ResultsDropdown', () => {
  beforeEach(async () => {
    await i18n.changeLanguage('en');
  });

  it('renders results and reports the selected one', async () => {
    const onSelect = vi.fn();
    render(<ResultsDropdown status="success" results={sampleResults} onSelect={onSelect} />);

    expect(screen.getByText('Lofi Beats')).toBeInTheDocument();
    expect(screen.getAllByRole('option')).toHaveLength(2);

    await userEvent.click(screen.getByText('Jazz Cafe'));
    expect(onSelect).toHaveBeenCalledWith(sampleResults[1]);
  });

  it('shows the empty message when a successful search has no results', () => {
    render(<ResultsDropdown status="success" results={[]} onSelect={vi.fn()} />);

    expect(screen.getByText('No results.')).toBeInTheDocument();
    expect(screen.queryByRole('option')).toBeNull();
  });

  it('shows the error message on error', () => {
    render(<ResultsDropdown status="error" results={[]} onSelect={vi.fn()} />);

    expect(screen.getByText('Search failed. Try again.')).toBeInTheDocument();
  });
});
