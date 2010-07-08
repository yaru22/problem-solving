{- n -> n/2 (n is even)
   n -> 3n + 1 (n is odd)
   Which starting number, under one million, produces the longest chain?
-}

import Control.Arrow
import Data.List
import Data.Array

collatz = assocs colarray
    where
      colarray = array (1,1000000)
                 $ (1,1) : [(n, coll n) | n<-[2..1000000]]
      coll n | odd n = 1 + check (3*n+1)
             | otherwise = 1 + check (n `div` 2)
      check n | n > 1000000 = coll n
              | otherwise = colarray!n

sol = foldl1' (\a b -> if (snd a > snd b) then a else b) collatz