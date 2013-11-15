// Digit canceling fractions
// 1/2 = 49/98 --> 4/8 = 1/2
//
// There are exactly four non-trivial examples of this type of fraction, less
// than one in value, and containing two digits in the numerator and denominator.
// If the product of these four fractions is given in its lowest common terms,
// find the value of the denominator.

(function () {
  'use strict';

  for (var i = 1; i <= 99; i++) {
    for (var j = i; j <= 99; j++) {
      if (i === j) {
        continue;
      }

      var iStr = '' + i;
      var jStr = '' + j;
      var canceled = false;
      for (var k = 1; k <= 9; k++) {
        var kStr = '' + k;
        if (iStr.search(kStr) >= 0 &&
            jStr.search(kStr) >= 0) {
          canceled = true;
          iStr = iStr.replace(kStr, '');
          jStr = jStr.replace(kStr, '');
        }
      }
      if (canceled && i / j === parseInt(iStr, 10) / parseInt(jStr, 10)) {
        console.log(i, j);
      }
    }
  }
})();
