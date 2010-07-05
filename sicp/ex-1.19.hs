{- A clever algorithm for computing the Fibonacci numbers in a logarithmic number
   of steps. -}

-- How I got p' and q': applied Tpq twice and grouped the terms to be in the
-- form of bQ + aQ + aP.
fib n = fibIter 1 0 0 1 n
fibIter a b p q count 
    | count == 0 = b
    | even count = fibIter a b (p^2 + q^2) (2*p*q + q^2) (count `div` 2)
    | otherwise  = fibIter (b*q + a*q + a*p) (b*p + a*q) p q (count-1)
                   
