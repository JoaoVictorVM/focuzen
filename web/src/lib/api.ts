import type { Video } from '../types/youtube';

type SearchResponse = {
  results: Video[];
};

// searchVideos calls the backend search proxy. The backend hides the YouTube
// API key and shapes the response; here we just unwrap the results.
export async function searchVideos(query: string, signal?: AbortSignal): Promise<Video[]> {
  const url = new URL('/api/v1/search', window.location.origin);
  url.searchParams.set('q', query);

  const response = await fetch(url, { signal });
  if (!response.ok) {
    throw new Error(`search failed with status ${response.status}`);
  }
  const data = (await response.json()) as SearchResponse;
  return data.results;
}
