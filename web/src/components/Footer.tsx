import { useTranslation } from 'react-i18next';

export function Footer() {
  const { t } = useTranslation();

  return (
    <footer className="px-4 py-3 text-center text-sm sm:px-6">
      {/* Same-origin route; the backend redirects to GitHub Releases. */}
      <a
        href="/download"
        target="_blank"
        rel="noreferrer"
        className="text-muted underline-offset-4 transition-colors hover:text-foreground hover:underline"
      >
        {t('footer.download')}
      </a>
    </footer>
  );
}
