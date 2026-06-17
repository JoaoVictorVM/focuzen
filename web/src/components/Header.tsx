import { useTranslation } from 'react-i18next';

import { LanguageToggle } from './LanguageToggle';
import { SearchBar } from './SearchBar';
import { ThemeToggle } from './ThemeToggle';

export function Header() {
  const { t } = useTranslation();

  return (
    <header className="flex items-center gap-4 px-4 py-3 sm:px-6">
      <span aria-label={t('app.title')} className="text-2xl tracking-wide text-foreground">
        <span aria-hidden="true">
          focu<span className="font-script text-accent">zen</span>
        </span>
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
