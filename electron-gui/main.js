const { app, BrowserWindow, ipcMain, Menu } = require('electron');
const path = require('path');
const { execFile } = require('child_process');
const axios = require('axios');

let mainWindow;
let goProcess;

function createWindow() {
    mainWindow = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            preload: path.join(__dirname, 'preload.js'),
            contextIsolation: true,
        },
    });
    mainWindow.loadFile('index.html');

    const menuTemplate = [
        {
          label: 'Test',
          submenu: [
            {
              label: 'Undo',
              role: 'undo',
            },
            {
              label: 'Redo',
              role: 'redo',
            },
            { type: 'separator' },
            {
              label: 'Cut',
              role: 'cut',
            },
            {
              label: 'Copy',
              role: 'copy',
            },
            {
              label: 'Paste',
              role: 'paste',
            },
          ],
        },
      ];

      const menu = Menu.buildFromTemplate(menuTemplate);

      Menu.setApplicationMenu(menu);


    setInterval(async () => {
        try {
            const response = await axios.get('http://localhost:8080/download-status');
            if (response.status === 200) {
                mainWindow.webContents.send('download-status-update', response.data);
            }
        } catch (error) {
            console.error('Failed to fetch download status:', error.message);
        }
    }, 2000);

    // اجرای برنامه Go به عنوان child process
    let goExecutable = path.join(__dirname, 'binaries', 'sfd-linux'); // به صورت پیش‌فرض
    if (process.platform === 'win32') {
        goExecutable = path.join(__dirname, 'binaries', 'sfd-windows.exe');
    } else if (process.platform === 'darwin') {
        goExecutable = path.join(__dirname, 'binaries', 'sfd-macos');
    }

    goProcess = execFile(goExecutable, (error, stdout, stderr) => {
        if (error) {
            console.error(`Error starting Go process: ${error}`);
        }
        console.log(`Go process stdout: ${stdout}`);
        console.error(`Go process stderr: ${stderr}`);
    });

    goProcess.on('close', (code) => {
        console.log(`Go process exited with code ${code}`);
    });
}

app.on('ready', createWindow);

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        if (goProcess) {
            goProcess.kill();
        }
        app.quit();
    }
});

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
    }
});

ipcMain.on('start-download', async (event, url) => {
    try {
        console.log('Starting download for:', url); // لاگ برای دیباگ
        const response = await axios.post('http://localhost:8080/start-download', { url });
        if (response.status === 200) {
            event.reply('download-status', 'Download started successfully.');
        } else {
            event.reply('download-status', 'Failed to start download: Unexpected response from server.');
        }
    } catch (error) {
        console.error('Error starting download:', error.message); // لاگ برای دیباگ
        event.reply('download-status', 'Failed to start download: ' + error.message);
    }
});

// مدیریت حذف دانلودهای لغو شده
ipcMain.on('clear-cancelled-downloads', async (event) => {
    try {
        console.log('Clearing cancelled downloads'); // لاگ برای دیباگ
        await axios.post('http://localhost:8080/clear-cancelled');
        event.reply('download-status', 'Cancelled downloads cleared successfully.');
    } catch (error) {
        console.error('Error clearing cancelled downloads:', error.message); // لاگ برای دیباگ
        event.reply('download-status', 'Failed to clear cancelled downloads: ' + error.message);
    }
});


ipcMain.on('pause-download', async (event, id) => {
    try {
        console.log('Pausing download with ID:', id); // لاگ برای دیباگ
        await axios.post(`http://localhost:8080/pause-download`, { id });
        event.reply('download-status', `Download ${id} paused successfully.`);
    } catch (error) {
        console.error('Error pausing download:', error.message); // لاگ برای دیباگ
        event.reply('download-status', `Failed to pause download ${id}: ${error.message}`);
    }
});

ipcMain.on('resume-download', async (event, id) => {
    try {
        console.log('Resuming download with ID:', id); // لاگ برای دیباگ
        await axios.post(`http://localhost:8080/resume-download`, { id });
        event.reply('download-status', `Download ${id} resumed successfully.`);
    } catch (error) {
        console.error('Error resuming download:', error.message); // لاگ برای دیباگ
        event.reply('download-status', `Failed to resume download ${id}: ${error.message}`);
    }
});

ipcMain.on('cancel-download', async (event, id) => {
    try {
        console.log('Cancelling download with ID:', id); // لاگ برای دیباگ
        await axios.post(`http://localhost:8080/cancel-download`, { id });
        event.reply('download-status', `Download ${id} cancelled successfully.`);
    } catch (error) {
        console.error('Error cancelling download:', error.message); // لاگ برای دیباگ
        event.reply('download-status', `Failed to cancel download ${id}: ${error.message}`);
    }
});