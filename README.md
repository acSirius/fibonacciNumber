# fibonacciNumber

A simple app made with go that generates the first 80 fibonacci numbers.

The first two numbers are hard coded, that is because at least 2 numbers are needed to generate the next (fibonacciNumber[i] can't be generated if fibonacciNumber[i - 1] and fibonacciNumber[i - 2] don't exist)

# How the app works

An array is made with a length of 80, then the first two numbers are added as the index 0 and 1 of the array.
A for loop will be executed 80 times, the first two times it just prints the first two fibonacci numbers, no need to generate them as they are hard coded.

After that, the loop will generate fibonacciNumber[i], it does this by adding the last 2 numbers in the array, and then the newly generated number is printed in the console.
