import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// In dev the SPA and the Go backend run on different ports; proxying /api keeps
// the frontend on the same-origin contract it has in production (embed.FS).
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
});
