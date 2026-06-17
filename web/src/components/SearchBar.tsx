import { useState } from 'react';
import { useTranslation } from 'react-i18next';

import { SearchIcon } from './icons';

// Visual shell only: the input holds local state but does not search yet. The
// search flow (debounce → backend → results dropdown) is wired in a later step.
export function SearchBar() {
  const { t } = useTranslation();
  const [query, setQuery] = useState('');

  return (
    <form
      role="search"
      onSubmit={(event) => event.preventDefault()}
      className="relative w-full max-w-md"
    >
      <SearchIcon
        size={16}
        strokeWidth={1.5}
        className="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-muted"
      />
      <input
        type="search"
        value={query}
        onChange={(event) => setQuery(event.target.value)}
        placeholder={t('search.placeholder')}
        aria-label={t('search.label')}
        className="w-full rounded-full border border-muted/30 bg-transparent py-2 pl-9 pr-4 text-sm text-foreground placeholder:text-muted focus:border-primary focus:outline-none"
      />
    </form>
  );
}
