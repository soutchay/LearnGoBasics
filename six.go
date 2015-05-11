/* Euler 6
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and
the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers
and the square of the sum.
*/

package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("Euler 6")
  start := time.Now()
  fmt.Println(differenceBrute(sumSquares(100), squareSum(100)))
  fmt.Println("It took brute:", time.Since(start))
  start = time.Now()
  fmt.Println(difference(100))
  fmt.Println("It took math:", time.Since(start))
}

func sumSquares(a int) int{
  sum := 0
  for i:=1;i<=a; i++{
    sum += power(i, 2)
  }
  return sum
}

func squareSum(a int) int{
  sum := 0
  for i:=1; i<=a; i++ {
    sum += i
  }
  return power(sum, 2)
}

func differenceBrute(a int, b int) int {
  //a is sum of the squares, b is square of the sum
  return b-a
}

func power(a int, b int) int{
  //a is the base, b is the exponent
  value := 1
  for i:=0; i<b; i++ {
    value *= a
  }
  return value
}

func difference(a int) int{
  sum := a*(a+1)/2 //sum of natural numbers
  square := (a*(a+1)*(2*a+1))/6 //sum of squares of natural numbers
  return sum*sum-square
}
