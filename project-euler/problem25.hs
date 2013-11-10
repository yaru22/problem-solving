-- What is the first term in the Fibonacci sequence to contain 1000 digits?

fibs :: [Integer]
fibs = 1 : 1 : zipWith (+) fibs (tail fibs)

solve :: Int
solve = length . takeWhile (\n -> (length . show $ n) < 1000) $ fibs

main :: IO ()
main = print (1 + solve)
