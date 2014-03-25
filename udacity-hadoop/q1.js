'use strict';

var fs    = require('fs');
var split = require('split');

var obj = {};
fs.createReadStream('./purchases.txt')
  .pipe(split())
  .on('data', function (line) {
    var cols = line.split('\t');
    obj[cols[3]] = obj[cols[3]] ? obj[cols[3]] : 0;
    obj[cols[3]] += parseFloat(cols[4]);
  })
  .on('end', function () {
    console.log(
        'Toys', obj['Toys'], '\n',
        'Consumer Electronics', obj['Consumer Electronics'], '\n');
  });
