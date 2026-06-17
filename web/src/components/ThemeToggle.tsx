import { useTranslation } from 'react-i18next';

import { useTheme } from '../hooks/useTheme';
import { MoonIcon, SunIcon } from './icons';

export function ThemeToggle() {
  const { theme, toggleTheme } = useTheme();
  const { t } = useTranslation();
  const isDark = theme === 'dark';

  return (
    <button
      type="button"
      onClick={toggleTheme}
      aria-label={t('theme.toggle')}
      aria-pressed={isDark}
      className="rounded p-1 text-muted transition-colors hover:text-foreground"
    >
      {isDark ? (
        <SunIcon size={18} strokeWidth={1.5} />
      ) : (
        <MoonIcon size={18} strokeWidth={1.5} />
      )}
    </button>
  );
}
