module Problem9 where

isPythagorean :: Int -> Int -> Int -> Bool
isPythagorean a b c = a^2 + b^2 == c^2

solve :: Int
solve = maximum [a*b*(1000-a-b) | a<-[1..998], b<-[1..998], a+b<1000, isPythagorean a b (1000-a-b)]

main :: IO ()
main = putStrLn . show $ solve

