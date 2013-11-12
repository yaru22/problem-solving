// 1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
// How many different ways can £2 be made using any number of coins?

(function () {
  'use strict';

  var coins = [0, 1, 2, 5, 10, 20, 50, 100, 200];
  var ways = {};
  ways[0] = {};
  ways[0][0] = 1;

  for (var i = 1; i <= 200; i++) {
    ways[i] = {};
    ways[i][0] = 0;
    for (var j = 1; j < coins.length; j++) {
      var coin = coins[j];
      if (coin > i) {
        break;
      }

      for (var k = j; k >= 1; k--) {
        if (coins[k] <= i - coin) {
          break;
        }
      }
      ways[i][j] = ways[i][j - 1] + ways[i - coin][k];
    }
  }

  console.log(ways[200][8]);
})();

// coins = [1, 2, 5, 10, 20, 50, 100, 200] (assume 1-index based);
// a[i][j] = number of ways to make up 'i' pences using coins[1..j]
//
// Initial data
// a[0][0] = 1
//
// Then, a[i][j] = a[i][j - 1] + a[i - coins[j]][k]
// s.t. coins[k] <= i - coins[j] and k <= j
//
// a[1][1] = a[1][0] + a[0][0] = 0 + 1 = 1; 1
// a[2][1] = a[2][0] + a[1][1] = 0 + 1 = 1; 1 1
// a[2][2] = a[2][1] + a[0][0] = 1 + 1 = 2; 1 1, 2
// a[3][1] = a[2][0] = 1; 1 1 1
// a[3][2] = a[3][0] + a[1][1] = 2; 1 1 1, 1 2
// a[4][1] = a[4][0] + a[3][1] = 1; 1 1 1 1
// a[4][2] = a[4][1] + a[2][2] = 3; 1 1 1 1, 1 1 2, 2 2
// a[5][1] = a[5][0] + a[4][1] = 1; 1 1 1 1 1
// a[5][2] = a[5][1] + a[3][2] = 3; 1 1 1 1 1, 1 1 1 2, 1 2 2
// a[5][3] = a[5][2] + a[0][0] = 4; 1 1 1 1 1, 1 1 1 2, 1 2 2, 5
// ...
