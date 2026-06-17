import { expect, test } from '@playwright/test';

const searchResults = {
  results: [
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
  ],
};

test.beforeEach(async ({ page }) => {
  // Deterministic backend: mock the search proxy.
  await page.route('**/api/v1/search*', (route) => route.fulfill({ json: searchResults }));
  // Keep the test offline: stub the YouTube IFrame API script.
  await page.route('**/iframe_api', (route) =>
    route.fulfill({ contentType: 'text/javascript', body: '' }),
  );
});

test('search, see results and play a track', async ({ page }) => {
  await page.goto('/');

  await page.getByRole('searchbox').fill('lofi');

  // Results dropdown appears.
  await expect(page.getByText('Lofi Beats')).toBeVisible();
  await expect(page.getByText('Jazz Cafe')).toBeVisible();

  // Selecting a result starts playback: the now-playing title and the
  // play/pause control appear.
  await page.getByText('Lofi Beats').click();

  await expect(page.getByText('Lofi Beats')).toBeVisible();
  await expect(page.getByRole('button', { name: /play|pause|tocar|pausar/i })).toBeVisible();
});
