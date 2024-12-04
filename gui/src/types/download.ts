export interface Download {
  id: string;
  fileName: string;
  url: string;
  progress: number;
  speed: number;
  size: number;
  status: 'downloading' | 'paused' | 'completed' | 'error';
  timeRemaining: number;
  error?: string;
  createdAt: Date;
}