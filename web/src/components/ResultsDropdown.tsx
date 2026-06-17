import { useTranslation } from 'react-i18next';

import type { SearchStatus } from '../hooks/useSearch';
import type { Video } from '../types/youtube';

type ResultsDropdownProps = {
  status: SearchStatus;
  results: Video[];
  onSelect: (video: Video) => void;
};

export function ResultsDropdown({ status, results, onSelect }: ResultsDropdownProps) {
  const { t } = useTranslation();

  const message =
    status === 'loading'
      ? t('search.loading')
      : status === 'error'
        ? t('search.error')
        : status === 'success' && results.length === 0
          ? t('search.empty')
          : null;

  return (
    <div className="absolute left-0 right-0 top-full z-10 mt-2 overflow-hidden rounded-lg border border-muted/20 bg-background shadow-lg">
      {message !== null && <p className="px-4 py-3 text-sm text-muted">{message}</p>}

      {results.length > 0 && (
        <ul role="listbox" className="max-h-80 overflow-y-auto py-1">
          {results.map((video) => (
            <li key={video.id} role="option" aria-selected={false}>
              <button
                type="button"
                onClick={() => onSelect(video)}
                className="flex w-full items-center gap-3 px-3 py-2 text-left transition-colors hover:bg-muted/10"
              >
                <img
                  src={video.thumbnailUrl}
                  alt=""
                  loading="lazy"
                  className="h-9 w-16 flex-none rounded object-cover"
                />
                <span className="min-w-0">
                  <span className="block truncate text-sm text-foreground">{video.title}</span>
                  <span className="block truncate text-xs text-muted">{video.channelTitle}</span>
                </span>
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
