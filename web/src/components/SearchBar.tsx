import { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { MIN_QUERY_LENGTH, useSearch } from '../hooks/useSearch';
import type { Video } from '../types/youtube';
import { SearchIcon } from './icons';
import { ResultsDropdown } from './ResultsDropdown';

type SearchBarProps = {
  onSelect?: (video: Video) => void;
};

export function SearchBar({ onSelect }: SearchBarProps) {
  const { t } = useTranslation();
  const { query, setQuery, results, status } = useSearch();
  const [open, setOpen] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    function handlePointerDown(event: MouseEvent) {
      if (containerRef.current && !containerRef.current.contains(event.target as Node)) {
        setOpen(false);
      }
    }
    document.addEventListener('mousedown', handlePointerDown);
    return () => document.removeEventListener('mousedown', handlePointerDown);
  }, []);

  const showDropdown = open && query.trim().length >= MIN_QUERY_LENGTH;

  return (
    <div ref={containerRef} className="relative w-full max-w-md">
      <form role="search" onSubmit={(event) => event.preventDefault()} className="relative">
        <SearchIcon
          size={16}
          strokeWidth={1.5}
          className="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-muted"
        />
        <input
          type="search"
          value={query}
          onChange={(event) => {
            setQuery(event.target.value);
            setOpen(true);
          }}
          onFocus={() => setOpen(true)}
          onKeyDown={(event) => {
            if (event.key === 'Escape') {
              setOpen(false);
            }
          }}
          placeholder={t('search.placeholder')}
          aria-label={t('search.label')}
          className="w-full rounded-full border border-muted/30 bg-transparent py-2 pl-9 pr-4 text-sm text-foreground placeholder:text-muted focus:border-primary focus:outline-none"
        />
      </form>

      {showDropdown && (
        <ResultsDropdown
          status={status}
          results={results}
          onSelect={(video) => {
            onSelect?.(video);
            setOpen(false);
          }}
        />
      )}
    </div>
  );
}
