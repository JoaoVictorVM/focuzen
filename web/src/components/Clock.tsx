import { useTranslation } from 'react-i18next';

type ClockProps = {
  now: Date;
};

export function Clock({ now }: ClockProps) {
  const { i18n } = useTranslation();
  const time = new Intl.DateTimeFormat(i18n.resolvedLanguage, {
    hour: '2-digit',
    minute: '2-digit',
  }).format(now);

  return (
    <time
      dateTime={now.toISOString()}
      className="block font-extralight tabular-nums tracking-tight text-foreground text-7xl sm:text-8xl"
    >
      {time}
    </time>
  );
}
