{- If all the numbers from 1 to 1000 (one thousand) inclusive were written out 
   in words, how many letters would be used? 
 -}

import Data.Char

num1 = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
        "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
        "seventeen", "eighteen", "nineteen"]
num10 = ["twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", 
         "ninety"]
num100 = "hundred"
num1000 = "thousand"

readNum n | n == 1000 = readNum (n `div` 1000) ++ " " ++ num1000
          | n >= 100 = 
              let m = n `mod` 100
              in if m == 0 then 
                     readNum (n `div` 100) ++ " " ++ num100 
                 else
                     readNum (n `div` 100) ++ " " ++ num100 ++ " and " ++ 
                     readNum (n `mod` 100)
          | n >= 20 = num10 !! ((n `div` 10) - 2) ++ " " ++ readNum (n `mod` 10)
          | n == 0 = ""
          | otherwise = num1 !! (n - 1)

countChar = filter isAlpha

sol = sum $ map (length . countChar . readNum) [1..1000]
                      