// The number 3797 has an interesting property. Being prime itself, it is possible
// to continuously remove digits from left to right, and remain prime at each
// stage: 3797, 797, 97, and 7. Similarly we can work from right to
// left: 3797, 379, 37, and 3.

// Find the sum of the only eleven primes that are both truncatable from left to
// right and right to left.

// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

(function () {
  'use strict';

  var memoize = { 1: false };
  function isPrime(n) {
    if (typeof memoize[n] !== 'undefined') {
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

  function isTruncatablePrime(n) {
    var truncatablePrime = true;
    var nL = n;
    var nR = n;
    var nStrL = n + '';
    var nStrR = n + '';
    for (var i = 0, l = nStrL.length; i < l; i++) {
      truncatablePrime = truncatablePrime && isPrime(nL) && isPrime(nR);
      if (!truncatablePrime) {
        break;
      }
      nStrL = nStrL.substring(1);
      nStrR = nStrR.substring(0, nStrR.length - 1);
      nL = parseInt(nStrL, 10);
      nR = parseInt(nStrR, 10);
    }
    return truncatablePrime;
  }

  var count = 0;
  var sum = 0;
  for (var i = 10; count < 11; i++) {
    if (isTruncatablePrime(i)) {
      count++;
      sum = sum + i;
    }
  }
  console.log(sum);
})();
