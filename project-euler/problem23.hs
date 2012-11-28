-- A perfect number is a number for which the sum of its proper divisors is exactly
-- equal to the number. For example, the sum of the proper divisors of 28 would be
-- 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

-- A number n is called deficient if the sum of its proper divisors is less than n
-- and it is called abundant if this sum exceeds n.

-- As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
-- number that can be written as the sum of two abundant numbers is 24. By
-- mathematical analysis, it can be shown that all integers greater than 28123 can
-- be written as the sum of two abundant numbers. However, this upper limit cannot
-- be reduced any further by analysis even though it is known that the greatest
-- number that cannot be expressed as the sum of two abundant numbers is less than
-- this limit.

-- Find the sum of all the positive integers which cannot be written as the sum of
-- two abundant numbers.

import Data.Array

sumOfDivisors :: Integer -> Integer
sumOfDivisors n = 1 + sum [d + (theOther d) | d <- [2..intSqrt], n `mod` d == 0]
  where
    intSqrt = floor . sqrt . fromIntegral $ n
    theOther d | (n `div` d /= d) = n `div` d
               | otherwise        = 0

isAbundant :: Integer -> Bool
isAbundant n = sumOfDivisors n > n

abundsArray :: Array Integer Bool
abundsArray = listArray (1, 28124) $ map isAbundant [1..28124]

abunds :: [Integer]
abunds = filter (abundsArray !) [1..28124]

isSumOfAbundants :: Integer -> Bool
isSumOfAbundants n = any (abundsArray !) . map (n -) . takeWhile (<= n `div` 2) $ abunds

main :: IO ()
main = print . sum . filter (not . isSumOfAbundants) $ [1..28124]
--main = print . sumOfDivisors $ 12
--main = print abunds
