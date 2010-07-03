module Problem4 where

isPalindrome :: String -> Bool
isPalindrome s = s == (reverse s)

isNumPalindrome :: Integer -> Bool
isNumPalindrome = isPalindrome . show 

solve :: Integer
solve = maximum [x*y | x<-[100..999], y<-[100..999], isNumPalindrome (x*y)]

main :: IO ()
main = putStrLn . show $ solve

