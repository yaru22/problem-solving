'use strict';

var fs    = require('fs');
var split = require('split');

var obj = {};
fs.createReadStream('./purchases.txt')
  .pipe(split())
  .on('data', function (line) {
    var cols = line.split('\t');
    obj[cols[2]] = obj[cols[2]] ? obj[cols[2]] : 0;
    obj[cols[2]] = Math.max(obj[cols[2]], parseFloat(cols[4]));
  })
  .on('end', function () {
    console.log(
      'Reno', obj['Reno'], '\n',
      'Toledo', obj['Toledo'], '\n',
      'Chandler', obj['Chandler']);
  });
