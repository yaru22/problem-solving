{- Write a procedure that computes elements of Pascal's triangle by means of a
   recursive process. 
   1
   1 1
   1 2 1
   1 3 3 1
   1 4 6 4 1
   ...
-}

-- Solution using recursive process; it seems very inefficient.
pascalTri1 r | r <= 0 = undefined
pascalTri1 1 = [1]
pascalTri1 r = let prevRow = pascalTri1 (r-1)
               in zipWith (+) (0:prevRow) (prevRow ++ [0])
--main = print $ [pascalTri1 r | r<-[1..5]]

-- Solution exploiting laziness.
pascalTri2 = [1] : (zipWith (\x y -> zipWith (+) (0:x) (y ++ [0])) pascalTri2 pascalTri2)
main = print . take 10 $ pascalTri2
