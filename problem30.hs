-- Find the sum of all the numbers that can be written as the sum of fifth
-- powers of their digits.

-- NOTE: First, we need to find an upper bound.
-- Since x * 9^5 has to be x digits, if we find such x, it is 6 and 6 * 9^5 = 354294.

power10 :: [Integer]
power10 = 1 : map (*10) power10

-- [flip div 10, flip div 100, flip div 1000, ..]
divs :: [Integer -> Integer]
divs = map (flip div) power10

digits :: Integer -> [Integer]
digits n = map (`mod` 10) . takeWhile (/= 0) . map ($ n) $ divs

powerDigits :: Integer -> Integer -> [Integer]
powerDigits p n = map (^ p) $ digits n

sumPowerDigits :: Integer -> Integer
sumPowerDigits n = sum $ powerDigits 5 n

solve :: [Integer]
solve = [n | n <- [2..354295], sumPowerDigits n == n]

main :: IO ()
main = print $ sum solve
