{- What is the sum of the digits of the number 2^1000? -}

import Data.Char

sol = sum . map digitToInt . show $ 2^1000