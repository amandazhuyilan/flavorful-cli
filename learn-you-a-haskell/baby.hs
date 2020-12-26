doubleMe x = x + x
doubleUs x y = x*2 + y*2
doubleSmallNumber x = if x > 100
                         then x
                         else x * 2
-- we use ' to either denote a strict version/slightly modified function or variable.
doubleSmallNumber' x = (if x > 100 then x else x * 2) + 1

-- write our own 'length'
length' xs = sum [1 | _ <- xs]

removeLowerCase st = [c | c <- st, c `elem` ['A' .. 'Z]]
removeUpperCase st = [c | c < -st, c `elem` ['a'..'z']]