// Considering quadratics of the form:
// n² + an + b, where |a| < 1000 and |b| < 1000
// where |n| is the modulus/absolute value of n
// e.g. |11| = 11 and |−4| = 4
// Find the product of the coefficients, a and b, for the quadratic expression
// that produces the maximum number of primes for consecutive values of n,
// starting with n = 0.

(function () {
  'use static';

  function isPrime(n) {
    n = (n < 0) ? -n : n;
    for (var i = 2; i * i <= n; i++) {
      if (n % i === 0) {
        return false;
      }
    }
    return true;
  }

  var maxSeq = -1;
  var product = -1;
  for (var a = -1000; a < 1000; a++) {
    for (var b = -1000; b < 1000; b++) {
      for (var n = 0; isPrime(n*n + a*n + b); n++) {
        if (n > maxSeq) {
          maxSeq = n;
          product = a * b;
        }
      }
    }
  }

  console.log(maxSeq, product);
})();
