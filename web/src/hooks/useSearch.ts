import { useEffect, useState } from 'react';

import { searchVideos } from '../lib/api';
import type { Video } from '../types/youtube';

const DEBOUNCE_MS = 350;
export const MIN_QUERY_LENGTH = 2;

export type SearchStatus = 'idle' | 'loading' | 'success' | 'error';

// useSearch debounces the query, cancels in-flight requests on change, and
// exposes the current results and status. Queries shorter than the minimum stay
// idle so we never spend YouTube quota on a single keystroke.
export function useSearch() {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState<Video[]>([]);
  const [status, setStatus] = useState<SearchStatus>('idle');

  useEffect(() => {
    const trimmed = query.trim();
    if (trimmed.length < MIN_QUERY_LENGTH) {
      setResults([]);
      setStatus('idle');
      return;
    }

    const controller = new AbortController();
    const timer = setTimeout(async () => {
      setStatus('loading');
      try {
        const videos = await searchVideos(trimmed, controller.signal);
        setResults(videos);
        setStatus('success');
      } catch {
        if (controller.signal.aborted) {
          return;
        }
        setResults([]);
        setStatus('error');
      }
    }, DEBOUNCE_MS);

    return () => {
      clearTimeout(timer);
      controller.abort();
    };
  }, [query]);

  return { query, setQuery, results, status } as const;
}
