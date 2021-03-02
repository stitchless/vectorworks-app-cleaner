const { shell, app } = require('electron')
const path = require('path');
const appPath = path.resolve(__dirname, "../../app")

// app.whenReady().then(() => {
    shell.openPath(path.join(appPath, "VectorworksUtility.exe")).then(r => {
        console.log(r)
    })

app.quit()


// })
// const ChildProcess = require('child_process');


// ChildProcess.spawn(`${path.resolve(__dirname, "../../app", "VectorworksUtility.exe")}`).unref();

// const { execFile } = require('child_process');
//
// const p = execFile((`${path.join(appPath, "VectorworksUtility.exe")}`), {
//     cwd: appPath,
// })
// p.stdout.on('data', (data) => {
//     console.log('stdout: ' + data)
// })
//
// p.stderr.on('data', (data) => {
//     console.log('stderr: ' + data)
// });
//
// p.on('close', (code) => {
//     console.log('child process exited with code ' + code)
// })