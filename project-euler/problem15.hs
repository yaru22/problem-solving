{- Starting in the top left corner of a 2x2 grid, there are 6 routes (without 
   backtracking) to the bottom right corner.
    _ _
   |_|_|
   |_|_|
   
   How many routes are there through a 20x20 grid?
 -}

-- Solution: This doesn't even need programming. Can be solved with mathematics.
-- sol = 40 C 20 (i.e. 40 choose 20)
sol = (product [21..40]) `div` (product [1..20])

