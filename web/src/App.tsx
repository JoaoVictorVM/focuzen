import { Clock } from './components/Clock';
import { DateDisplay } from './components/DateDisplay';
import { LanguageToggle } from './components/LanguageToggle';
import { ThemeToggle } from './components/ThemeToggle';
import { useNow } from './hooks/useNow';

export default function App() {
  const now = useNow();

  return (
    <main className="relative flex min-h-screen items-center justify-center bg-background text-foreground">
      <div className="absolute right-4 top-4 flex items-center gap-2">
        <LanguageToggle />
        <ThemeToggle />
      </div>
      <div className="text-center">
        <Clock now={now} />
        <DateDisplay now={now} />
      </div>
    </main>
  );
}
