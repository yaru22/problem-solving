// Find the sum of all numbers which are equal to the sum of the factorial of
// their digits.
// Note: as 1! = 1 and 2! = 2 are not sums they are not included.

(function () {
  'use strict';

  var memoize = {};
  memoize[0] = 1;
  function fact(n) {
    if (memoize[n]) {
      return memoize[n];
    }
    var answer = n * fact(n-1);
    memoize[n] = answer;
    return answer;
  }

  var answer = 0;

  // Since 9! * 7 = 2540160 (7 digits), that's the upper bound.
  for (var n = 3; n < 2540160; n++) {
    var sum = 0;
    var num = n;
    while (num !== 0) {
      var digit = num % 10;
      sum += fact(digit);
      num = Math.floor(num/10);
    }
    if (sum === n) {
      answer += sum;
    }
  }

  console.log(answer);
})();
