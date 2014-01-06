// Take the number 192 and multiply it by each of 1, 2, and 3:
// 192 × 1 = 192
// 192 × 2 = 384
// 192 × 3 = 576
//
// We will call 192384576 the concatenated product of 192 and (1,2,3).
//
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as
// the concatenated product of an integer with (1,2, ... , n) where n > 1?

(function () {
  'use strict';

  function isPanDigital(n) {
    var nStr = n + '';
    if (nStr.length != 9) {
      return false;
    }
    for (var i = 1; i <= 9; i++) {
      if (nStr.indexOf(i + '') == -1) {
        return false;
      }
    }
    return true;
  }

  function product(num, n) {
    var str = '';
    for (var i = 1; i <= n; i++) {
      str = str + (num * i);
    }
    return parseInt(str, 10);
  }

  var max = -1;
  for (var i = 9999; i >= 1; i--) {
    for (var j = 2; j <=9; j++) {
      var n = product(i, j);
      if (isPanDigital(n) && n > max) {
        max = n;
      }
    }
  }
  console.log(max);
})();
