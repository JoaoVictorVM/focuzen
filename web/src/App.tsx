import { useTranslation } from 'react-i18next';

import { LanguageToggle } from './components/LanguageToggle';

export default function App() {
  const { t } = useTranslation();

  return (
    <main className="relative flex min-h-screen items-center justify-center bg-background text-foreground">
      <div className="absolute right-4 top-4">
        <LanguageToggle />
      </div>
      <div className="text-center">
        <h1 className="text-3xl font-light tracking-wide text-primary">{t('app.title')}</h1>
        <p className="mt-2 text-muted">{t('app.tagline')}</p>
      </div>
    </main>
  );
}
