import { AudioControls } from './components/AudioControls';
import { AudioPlayer } from './components/AudioPlayer';
import { Clock } from './components/Clock';
import { DateDisplay } from './components/DateDisplay';
import { Footer } from './components/Footer';
import { Header } from './components/Header';
import { useNow } from './hooks/useNow';
import { usePlayer } from './hooks/usePlayer';

export default function App() {
  const now = useNow();
  const player = usePlayer();

  return (
    <div className="flex min-h-screen flex-col bg-background text-foreground">
      <Header onSearchSelect={player.play} />
      <main className="flex flex-1 items-center justify-center">
        <div className="text-center">
          <Clock now={now} />
          <DateDisplay now={now} />
          <AudioControls
            current={player.current}
            isPlaying={player.isPlaying}
            volume={player.volume}
            repeat={player.repeat}
            hasNext={player.hasNext}
            hasPrevious={player.hasPrevious}
            onTogglePlay={player.togglePlay}
            onNext={player.next}
            onPrevious={player.previous}
            onToggleRepeat={player.toggleRepeat}
            onVolumeChange={player.changeVolume}
          />
        </div>
      </main>
      <Footer />
      <AudioPlayer containerRef={player.containerRef} />
    </div>
  );
}
