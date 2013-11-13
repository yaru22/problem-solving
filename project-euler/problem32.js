// We shall say that an n-digit number is pandigital if it makes use of all the
// digits 1 to n exactly once; for example, the 5-digit number, 15234, is
// 1 through 5 pandigital.
//
// The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
// multiplicand, multiplier, and product is 1 through 9 pandigital.
//
// Find the sum of all products whose multiplicand/multiplier/product identity
// can be written as a 1 through 9 pandigital.
//
// HINT: Some products can be obtained in more than one way so be sure to only
// include it once in your sum.

(function () {
  'use strict';

  var solProducts = {};
  for (var i = 1; i <= 98; i++) {
    loop:
    for (var j = i; j <= 9876 / i; j++) {
      var product = i * j;
      var str = '' + i + j + product;
      if (str.length !== 9) {
        continue;
      }
      for (var k = 1; k <= 9; k++) {
        if (str.indexOf('' + k) < 0) {
          continue loop;
        }
      }
      console.log(i, j, product);
      solProducts[product] = true;
    }
  }

  var answer = 0;
  for (var sol in solProducts) {
    answer += parseInt(sol, 10);
  }
  console.log(answer);
})();
