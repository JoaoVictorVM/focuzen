import { Clock } from './components/Clock';
import { DateDisplay } from './components/DateDisplay';
import { Footer } from './components/Footer';
import { Header } from './components/Header';
import { useNow } from './hooks/useNow';

export default function App() {
  const now = useNow();

  return (
    <div className="flex min-h-screen flex-col bg-background text-foreground">
      <Header />
      <main className="flex flex-1 items-center justify-center">
        <div className="text-center">
          <Clock now={now} />
          <DateDisplay now={now} />
        </div>
      </main>
      <Footer />
    </div>
  );
}
