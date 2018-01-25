import Control.Monad
import System.Exit

main = do
    putStr "Please enter a character to be guessed\n"
    c <- getLine
    let oc = read c :: Int 
    moinloop oc

moinloop oc = do
    putStr "Please enter a character to guess\n"
    guesStr <- getLine
    let guess = read guesStr :: Int
    let mif oc guess = do
        if oc == guess
        then putStrLn("You win")
        else if oc > guess
            then do putStrLn("Too low loser"); moinloop oc
            else do putStrLn("Too high sucker"); moinloop oc
    mif oc guess
