// The number, 197, is called a circular prime because all rotations of the
// digits: 197, 971, and 719, are themselves prime.
// How many circular primes are there below one million?

(function () {
  'use strict';

  var memoize = {};
  function isPrime(n) {
    if (memoize[n]) {
      return memoize[n];
    }
    n = (n < 0) ? -n : n;
    for (var i = 2; i * i <= n; i++) {
      if (n % i === 0) {
        memoize[n] = false;
        return false;
      }
    }
    memoize[n] = true;
    return true;
  }

  function isCircularPrime(n) {
    var circularPrime = true;
    var nStr = n + '';
    for (var i = 0; i < nStr.length; i++) {
      circularPrime = circularPrime && isPrime(n);
      if (!circularPrime) {
        break;
      }
      nStr = nStr.substring(1) + nStr[0];
      n = parseInt(nStr, 10);
    }
    return circularPrime;
  }

  var count = 0;
  for (var i = 2; i <= 1000000; i++) {
    if (isCircularPrime(i)) {
      count++;
    }
  }
  console.log(count);
})();
