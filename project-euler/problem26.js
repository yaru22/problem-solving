// Find the value of d < 1000 for which 1/d contains the longest recurring cycle
// in its decimal fraction part.

(function () {
  'use strict';

  function getCycleLength(n) {
    var value = 1;
    var remainders = {};
    for (var i = 1;; i++) {
      value = value * 10;
      value = value % n;
      if (value === 0) {
        return 0;
      } else if (remainders[value]) {
        return i - remainders[value];
      }
      remainders[value] = i;
    }
  }

  var d = 0;
  var longestLength = -1;
  for (var i = 999; i >= 1; i--) {
    var cycleLength = getCycleLength(i);
    if (cycleLength > longestLength) {
      d = i;
      longestLength = cycleLength;
    }
    if (i < longestLength) {
      break;
    }
  }
  console.log('The length of the longest recurring cycle is %d for d = %d', longestLength, d);
})();
