{- Tower of Hanoi -}

hanoi n = moveTower n 1 3 2

moveDisc n from to = print $ "Move disc " ++ (show n) ++ 
                     " from pole " ++ (show from) ++ " to pole " ++ (show to)
moveTower n from to through 
    | n == 0 = return ()
    | otherwise = do
  moveTower (n-1) from through to
  moveDisc n from to
  moveTower (n-1) through to from
