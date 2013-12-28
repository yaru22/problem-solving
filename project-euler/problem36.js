// The decimal number, 585 = 1001001001 (binary), is palindromic in both bases.
// Find the sum of all numbers, less than one million, which are palindromic in
// base 10 and base 2.

(function () {
  'use strict';

  function isPalindrome(str) {
    for (var i = 0; i < ~~(str.length / 2); i++) {
      if (str[i] !== str[str.length - 1 - i]) {
        return false;
      }
    }
    return true;
  }

  function toBinary(n) {
    return n.toString(2);
  }

  var sum = 0;
  for (var i = 1; i < 1000000; i++) {
    if (isPalindrome(i + '') && isPalindrome(toBinary(i))) {
      sum = sum + i;
    }
  }
  console.log(sum);
})();
