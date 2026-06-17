// Mirrors the backend's Video JSON (server/internal/youtube).
export type Video = {
  id: string;
  title: string;
  channelTitle: string;
  thumbnailUrl: string;
};
