import { useTranslation } from 'react-i18next';

import { supportedLanguages, type SupportedLanguage } from '../i18n';

const shortLabels: Record<SupportedLanguage, string> = {
  'pt-BR': 'PT',
  en: 'EN',
};

export function LanguageToggle() {
  const { i18n, t } = useTranslation();

  return (
    <div role="group" aria-label={t('language.label')} className="flex gap-1 text-sm">
      {supportedLanguages.map((language) => {
        const active = i18n.resolvedLanguage === language;
        return (
          <button
            key={language}
            type="button"
            onClick={() => void i18n.changeLanguage(language)}
            aria-pressed={active}
            className={`rounded px-2 py-1 transition-colors ${
              active ? 'font-medium text-foreground' : 'text-muted hover:text-foreground'
            }`}
          >
            {shortLabels[language]}
          </button>
        );
      })}
    </div>
  );
}
