//Euler 3
/* Notes
  1) We only need to check the square root of the target number for primes if it is even
  2) 2 is the only even prime factor a target number can have
  3) Even numbers cannot be prime factors due to being divisible by 2
*/
package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("Euler", math.Sqrt(9))
  // primeFactors(600851475143)
  // primeNumber(147)
  primeFactors(147)
  primeFactors(97898)
}

//need to use int64 for large integers -9223372036854775808 to 9223372036854775807 
func primeFactors(n int64) {
  primeSlice := []int64{}
  fmt.Println("Finding prime factors of:", n)
  checkNumber := int64(math.Floor(math.Sqrt(float64(n)))) //prime factor not greater than checkNumber
  var i int64 = 3
  if n%2 == 0 {
    fmt.Println("Check up to", checkNumber)
    primeSlice = append(primeSlice, 2)
  }
  for ; i<=n; i++{
    if n%i == 0 && isPrime(i){
      fmt.Println("divisible by", i)
      primeSlice = append(primeSlice, i)
    }
  }
  fmt.Println(primeSlice)
}

func isPrime(n int64) bool{
  var i int64 = 2 //start at 2 since prime number is a factor of itself and 1
  for ; i<n; i++ {
    if n%i == 0 { //check if n is divisible by any number other than itself
      return false //not a prime number if it is divisible by a number other than itself and 1
    }
  }
  return true
}