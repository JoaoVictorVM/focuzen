import { useTranslation } from 'react-i18next';

import { ClockIcon } from './icons';
import { LanguageToggle } from './LanguageToggle';
import { SearchBar } from './SearchBar';
import { ThemeToggle } from './ThemeToggle';

export function Header() {
  const { t } = useTranslation();

  return (
    <header className="flex items-center gap-4 px-4 py-3 sm:px-6">
      <span className="flex items-center gap-2 text-lg font-medium tracking-wide text-foreground">
        <ClockIcon size={20} strokeWidth={1.5} className="text-primary" />
        {t('app.title')}
      </span>

      <div className="flex flex-1 justify-center">
        <SearchBar />
      </div>

      <div className="flex items-center gap-2">
        <LanguageToggle />
        <ThemeToggle />
      </div>
    </header>
  );
}
