{- How many different ways can we make change of $ 1.00, given half-dollars,
   quarters, dimes, nickels, and pennies? More generally, can we write a
   procedure to compute the number of ways to change any given amount of 
   money? -}

-- Solution 1: Using recursive procedure.
denominations = [50, 25, 10, 5, 1]
countChange1 amnt denoms 
    | denoms == [] = 0
    | amnt < 0     = 0
    | amnt == 0    = 1
    | otherwise    = (countChange1 (amnt - (head denoms)) denoms) +
                     (countChange1 amnt (tail denoms))
main = print $ countChange1 100 denominations

-- Solution 2: Using dynamic Programming.
countChange2 = undefined
--main = print $ countChange2
