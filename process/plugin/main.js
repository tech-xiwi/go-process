const fs = require('fs');
const { spawn } = require('child_process');
const out = fs.openSync('./out.log', 'a');
const err = fs.openSync('./out.log', 'a');
var server = spawn(
    "D:\\workspace\\golang\\src\\studio.xiwi\\process\\process.exe",[]
);
// var server = spawn(
//     "D:\\workspace\\golang\\src\\studio.xiwi\\process\\process.exe",[], {
//         detached: false,
//         stdio: [ 'inherit', out, err ]
//     }
// );
console.log("start timeout");
setTimeout(function () {
    console.log("kill child process");
    server.kill();
},30*1000);