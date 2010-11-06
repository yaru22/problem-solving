sol :: Integer -> Integer
sol n = (sum [3,6..n]) + (sum [5,10..n]) - (sum [15,30..n])


