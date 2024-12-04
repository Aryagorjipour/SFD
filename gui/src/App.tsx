import { useState, useCallback } from 'react';
import { Download } from './types/download';
import { DownloadForm } from './components/DownloadForm';
import { DownloadItem } from './components/DownloadItem';
import { Download as DownloadIcon } from 'lucide-react';
import { ThemeToggle } from './components/ThemeToggle';

function App() {
  const [downloads, setDownloads] = useState<Download[]>([]);

  const handleNewDownload = useCallback((url: string) => {
    const fileName = url.split('/').pop() || 'unknown';
    const newDownload: Download = {
      id: Math.random().toString(36).substr(2, 9),
      fileName,
      url,
      progress: 0,
      speed: 0,
      size: 0,
      status: 'downloading',
      timeRemaining: 0,
      createdAt: new Date(),
    };

    setDownloads(prev => [newDownload, ...prev]);

    // Simulate download progress (remove this in production and integrate with actual backend)
    let progress = 0;
    const interval = setInterval(() => {
      progress += Math.random() * 10;
      if (progress >= 100) {
        progress = 100;
        clearInterval(interval);
        setDownloads(prev =>
            prev.map(d =>
                d.id === newDownload.id
                    ? { ...d, progress: 100, status: 'completed' }
                    : d
            )
        );
      } else {
        setDownloads(prev =>
            prev.map(d =>
                d.id === newDownload.id
                    ? {
                      ...d,
                      progress,
                      speed: 1024 * 1024 * (Math.random() * 2 + 1),
                      size: 1024 * 1024 * 100,
                      timeRemaining: (100 - progress) * 2,
                    }
                    : d
            )
        );
      }
    }, 1000);
  }, []);

  const handlePauseResume = useCallback((id: string) => {
    setDownloads(prev =>
        prev.map(d =>
            d.id === id
                ? { ...d, status: d.status === 'paused' ? 'downloading' : 'paused' }
                : d
        )
    );
  }, []);

  const handleCancel = useCallback((id: string) => {
    setDownloads(prev => prev.filter(d => d.id !== id));
  }, []);

  return (
      <div className="min-h-screen bg-gray-50 dark:bg-gray-900 transition-colors">
        <div className="max-w-4xl mx-auto p-6">
          <div className="flex items-center justify-between mb-8">
            <div className="flex items-center space-x-3">
              <DownloadIcon className="w-8 h-8 text-blue-500" />
              <h1 className="text-2xl font-semibold text-gray-900 dark:text-white">Download Manager</h1>
            </div>
            <ThemeToggle />
          </div>

          <DownloadForm onSubmit={handleNewDownload} />

          <div className="space-y-4">
            {downloads.map(download => (
                <DownloadItem
                    key={download.id}
                    download={download}
                    onPauseResume={handlePauseResume}
                    onCancel={handleCancel}
                />
            ))}

            {downloads.length === 0 && (
                <div className="text-center py-12">
                  <DownloadIcon className="w-12 h-12 text-gray-400 dark:text-gray-600 mx-auto mb-4" />
                  <p className="text-gray-500 dark:text-gray-400">No downloads yet. Add a URL to start downloading.</p>
                </div>
            )}
          </div>
        </div>
      </div>
  );
}

export default App;