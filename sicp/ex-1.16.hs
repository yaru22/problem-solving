{- Design a procedure that evolves an iterative exponentiation process that uses
   successive squaring and uses a logarithmic number of steps, as does fast-expt. 
 -}

fastexp b n = fastexp' 1 b n
fastexp' a b n | n == 0 = a
               | even n = fastexp' a (b*b) (n `div` 2)
               | otherwise = fastexp' (a*b) b (n-1)

main = print $ fastexp 2 87
