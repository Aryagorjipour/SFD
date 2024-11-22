const { app, BrowserWindow } = require('electron');
const path = require('path');
const { execFile } = require('child_process');

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
