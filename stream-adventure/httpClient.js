var request = require('request');
var req = request.post('http://localhost:8000');

process.stdin
    .pipe(req)
    .pipe(process.stdout);
