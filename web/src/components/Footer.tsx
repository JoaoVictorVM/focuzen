import { useTranslation } from 'react-i18next';

// Points at GitHub Releases, where the CLI binaries are published (Phase 6/7).
const cliReleasesUrl = 'https://github.com/JoaoVictorVM/focuzen/releases';

export function Footer() {
  const { t } = useTranslation();

  return (
    <footer className="px-4 py-3 text-center text-sm sm:px-6">
      <a
        href={cliReleasesUrl}
        target="_blank"
        rel="noreferrer"
        className="text-muted underline-offset-4 transition-colors hover:text-foreground hover:underline"
      >
        {t('footer.download')}
      </a>
    </footer>
  );
}
