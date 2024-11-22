const { contextBridge, ipcRenderer } = require('electron');

contextBridge.exposeInMainWorld('electron', {
    startDownload: (url) => ipcRenderer.send('start-download', url),
    pauseDownload: (id) => ipcRenderer.send('pause-download', id),
    resumeDownload: (id) => ipcRenderer.send('resume-download', id),
    cancelDownload: (id) => ipcRenderer.send('cancel-download', id),
    clearCancelledDownloads: () => ipcRenderer.send('clear-cancelled-downloads'),
    onDownloadStatus: (callback) => ipcRenderer.on('download-status', (event, status) => {
        console.log('Received download status:', status); // لاگ برای دیباگ
        callback(status);
    }),
    onDownloadStatusUpdate: (callback) => ipcRenderer.on('download-status-update', (event, statuses) => {
        console.log('Received download statuses update:', statuses); // لاگ برای دیباگ
        callback(statuses);
    }),
});
