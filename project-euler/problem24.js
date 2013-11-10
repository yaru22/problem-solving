// A permutation is an ordered arrangement of objects. For example, 3124 is one
// possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are
// listed numerically or alphabetically, we call it lexicographic order.
// The lexicographic permutations of 0, 1 and 2 are:
// 012   021   102   120   201   210
// What is the millionth lexicographic permutation of the digits 0~9?

(function () {
  'use strict';

  function fact(n) {
    if (n === 0) {
      return 1;
    }
    return n * fact(n - 1);
  }

  function search(index, permStr, answer) {
    for (var i = 0; i < permStr.length; i++) {
      var ch = permStr[i];
      var rest = permStr.substring(0, i) + permStr.substring(i + 1);
      var numPermutations = fact(rest.length);
      if (index <= numPermutations) {
        return search(index, rest, answer + ch);
      } else if (index > numPermutations) {
        index = index - numPermutations;
      }
    }
    return answer;
  }

  console.log(search(1000000, '0123456789', ''));
})();
