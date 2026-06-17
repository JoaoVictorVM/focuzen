import { http, HttpResponse } from 'msw';

import type { Video } from '../../types/youtube';

export const sampleResults: Video[] = [
  {
    id: 'v1',
    title: 'Lofi Beats',
    channelTitle: 'Chillhop',
    thumbnailUrl: 'https://i.ytimg.com/vi/v1/mqdefault.jpg',
  },
  {
    id: 'v2',
    title: 'Jazz Cafe',
    channelTitle: 'Cafe Music',
    thumbnailUrl: 'https://i.ytimg.com/vi/v2/mqdefault.jpg',
  },
];

export const handlers = [
  http.get('*/api/v1/search', ({ request }) => {
    const query = new URL(request.url).searchParams.get('q') ?? '';
    if (query.includes('fail')) {
      return new HttpResponse(null, { status: 502 });
    }
    return HttpResponse.json({ results: sampleResults });
  }),
];
