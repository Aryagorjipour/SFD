import React from 'react';
import { Play, Pause, X, AlertCircle, CheckCircle, Download } from 'lucide-react';
import { Download as DownloadType } from '../types/download';
import { formatBytes, formatSpeed, formatTimeRemaining } from '../utils/format';

interface DownloadItemProps {
  download: DownloadType;
  onPauseResume: (id: string) => void;
  onCancel: (id: string) => void;
}

export const DownloadItem: React.FC<DownloadItemProps> = ({
                                                            download,
                                                            onPauseResume,
                                                            onCancel,
                                                          }) => {
  const getStatusIcon = () => {
    switch (download.status) {
      case 'downloading':
        return <Download className="w-5 h-5 text-blue-500" />;
      case 'completed':
        return <CheckCircle className="w-5 h-5 text-green-500" />;
      case 'error':
        return <AlertCircle className="w-5 h-5 text-red-500" />;
      default:
        return null;
    }
  };

  return (
      <div className="bg-white dark:bg-gray-800 rounded-lg shadow-sm p-4 mb-3">
        <div className="flex items-center justify-between mb-2">
          <div className="flex items-center space-x-2">
            {getStatusIcon()}
            <h3 className="font-medium text-gray-900 dark:text-white">{download.fileName}</h3>
          </div>
          <div className="flex items-center space-x-2">
            {download.status !== 'completed' && (
                <>
                  <button
                      onClick={() => onPauseResume(download.id)}
                      className="p-1 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-full transition-colors"
                  >
                    {download.status === 'paused' ? (
                        <Play className="w-5 h-5 text-gray-600 dark:text-gray-400" />
                    ) : (
                        <Pause className="w-5 h-5 text-gray-600 dark:text-gray-400" />
                    )}
                  </button>
                  <button
                      onClick={() => onCancel(download.id)}
                      className="p-1 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-full transition-colors"
                  >
                    <X className="w-5 h-5 text-gray-600 dark:text-gray-400" />
                  </button>
                </>
            )}
          </div>
        </div>

        <div className="space-y-2">
          <div className="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2">
            <div
                className="bg-blue-500 h-2 rounded-full transition-all duration-300"
                style={{ width: `${download.progress}%` }}
            />
          </div>

          <div className="flex justify-between text-sm text-gray-500 dark:text-gray-400">
            <div className="space-x-4">
              <span>{formatBytes(download.size)}</span>
              {download.status === 'downloading' && (
                  <>
                    <span>{formatSpeed(download.speed)}</span>
                    <span>{formatTimeRemaining(download.timeRemaining)}</span>
                  </>
              )}
            </div>
            <span>{download.progress.toFixed(1)}%</span>
          </div>

          {download.error && (
              <p className="text-sm text-red-500 dark:text-red-400">{download.error}</p>
          )}
        </div>
      </div>
  );
};