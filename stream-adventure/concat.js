var concat = require('concat-stream');

process.stdin
    .pipe(concat(function (buf) {
      process.stdout.write(buf.toString().split('').reverse().join('') + '\n');
    }));
