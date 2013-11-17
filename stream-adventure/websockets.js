var ws = require('websocket-stream');
var stream = ws('ws://localhost:8000');
stream.write('hello\n');
stream.end();
