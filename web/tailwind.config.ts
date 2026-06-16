import type { Config } from 'tailwindcss';

// Colors map to CSS variables (see src/styles/tokens.css). The dark theme only
// reassigns those variables under .dark — components never hardcode a color.
export default {
  content: ['./index.html', './src/**/*.{ts,tsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        background: 'var(--color-background)',
        foreground: 'var(--color-foreground)',
        primary: 'var(--color-primary)',
        muted: 'var(--color-muted)',
      },
    },
  },
  plugins: [],
} satisfies Config;
