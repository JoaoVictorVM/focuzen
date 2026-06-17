// Loads the YouTube IFrame Player API script once and resolves when the global
// YT object is ready. The script and the player iframe come from the YouTube
// origins allowed by the backend CSP.
let loadPromise: Promise<typeof YT> | null = null;

export function loadYouTubeIframeAPI(): Promise<typeof YT> {
  if (window.YT?.Player) {
    return Promise.resolve(window.YT);
  }
  if (loadPromise) {
    return loadPromise;
  }

  loadPromise = new Promise((resolve) => {
    const previousCallback = window.onYouTubeIframeAPIReady;
    window.onYouTubeIframeAPIReady = () => {
      previousCallback?.();
      resolve(window.YT as typeof YT);
    };

    const script = document.createElement('script');
    script.src = 'https://www.youtube.com/iframe_api';
    document.head.appendChild(script);
  });

  return loadPromise;
}
