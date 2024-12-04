import React, { useState } from 'react';
import { Plus } from 'lucide-react';

interface DownloadFormProps {
  onSubmit: (url: string) => void;
}

export const DownloadForm: React.FC<DownloadFormProps> = ({ onSubmit }) => {
  const [url, setUrl] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (url.trim()) {
      onSubmit(url.trim());
      setUrl('');
    }
  };

  return (
      <form onSubmit={handleSubmit} className="mb-6">
        <div className="flex gap-2">
          <input
              type="url"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="Enter download URL"
              className="flex-1 px-4 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              required
          />
          <button
              type="submit"
              className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors dark:focus:ring-offset-gray-900"
          >
            <Plus className="w-5 h-5" />
          </button>
        </div>
      </form>
  );
};