import { useTranslation } from 'react-i18next';

type DateDisplayProps = {
  now: Date;
};

export function DateDisplay({ now }: DateDisplayProps) {
  const { i18n } = useTranslation();
  const date = new Intl.DateTimeFormat(i18n.resolvedLanguage, {
    weekday: 'long',
    month: 'long',
    day: 'numeric',
  }).format(now);

  return <p className="mt-3 text-lg text-muted first-letter:uppercase">{date}</p>;
}
