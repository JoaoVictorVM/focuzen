import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it } from 'vitest';

import i18n from '../i18n';
import { LanguageToggle } from './LanguageToggle';

describe('LanguageToggle', () => {
  beforeEach(async () => {
    await i18n.changeLanguage('pt-BR');
  });

  it('marks the active language and switches on click', async () => {
    render(<LanguageToggle />);

    const pt = screen.getByRole('button', { name: 'PT' });
    const en = screen.getByRole('button', { name: 'EN' });

    expect(pt).toHaveAttribute('aria-pressed', 'true');
    expect(en).toHaveAttribute('aria-pressed', 'false');

    await userEvent.click(en);

    expect(i18n.resolvedLanguage).toBe('en');
    expect(en).toHaveAttribute('aria-pressed', 'true');
  });
});
