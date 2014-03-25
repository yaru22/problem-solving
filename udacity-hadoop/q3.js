'use strict';

var fs    = require('fs');
var split = require('split');

var obj = {
  numSales: 0,
  totalSales: 0
};

fs.createReadStream('./purchases.txt')
  .pipe(split())
  .on('data', function (line) {
    var cols = line.split('\t');
    if (typeof cols[4] === 'undefined') return;
    obj.numSales++;
    obj.totalSales += parseFloat(cols[4]);
  })
  .on('end', function () {
    console.log(
        'Total # sales:', obj.numSales, '\n',
        'Total Value of Sales:', obj.totalSales);
  });
