/*Euler 5
2520 is the smallest number that can be divided by each 
of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by 
all of the numbers from 1 to 20?
*/

/*Some thoughts: 
1) every number is divisible by 1
2) number must be even to be divisible by 2
3) adding all individual digits must be divisible by 3
4) divisible by 4: last two digits must be divisible by 4
5) divisible by 5: last digit must be 5 or 0; thus last digit must be 0
6) divisible by 6: #2 and #3 must be true; both divisible by 2 and 3
7) if a number is divisible by 12 then it is also divisble by 2,3,6,4
8) using the rule for factors we can see that we need each factor for the numbers 1-20
9) we can check if a number is a prime
10) let's look for a pattern in terms of primes, excluding 1
  2 - prime factors: 2
  3 - prime factors: 3
  4 - prime factors: 4=2*2
  5 - prime factors: 5

  We know that for a number to be able to divisible by both 2 and 3 we need both prime factors
  divisible, therefore a number needs to be divisible by all prime factors
  For smallest value to be divisible by all numbers if first column:
  2 : 2 = 2
  2,3 : 2*3 = 6
  //since we already have 2 from before, we need only to include another 2 to make up 4
  2,3,4 : 2*2*3 = 12
  2,3,4,5 : 2*2*3*5 = 60
  //we already have the prime factors to make up 6 included from before thus we do not need to add another factor
  2,3,4,5,6 : 2*2*3*5 = 60
  //7 is a prime factor
  2,3,4,5,6,7 : 2*2*3*5*7 = 420
  // 8 is 2*2*2 thus we only need to multiply by another 2
  2,3,4,5,6,7,8 : 2*2*3*5*7*2 = 840
  ...
  2-20 : 2^4*3^2*5*7*11*13*17*19 = 232792560

  -Need to keep count how many of each prime factor is necessary from 1 to 20
  -iterate from 2 to 20, divide by every known prime less than that number
    most number of 2s needed for values 1 to 10? 2^3=8, 2^4=12 thus 3
    most number of 2s needed for values 1 to 20? 2^4=16, 2^5=32 thus 4
    most number of 3s needed for values 1 to 20? 3^2= 9, 3^3= 27 thus 2
  -thus we need a double for loop that iterates 2-20
  arrayNum := map[int]int{1:1}
  for i:=2; i<=number; i++ {
    for j:=1; j<=number; j++ {
      if i^j > number {
        arrayNum[i] = j-1 // arrayNum = map[int]int{1:1, 2:3, 3:4}
      }
    }
  }
*/
package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("Euler 5")
  start := time.Now()
  fmt.Println("Smallest number is:", bruteForce())
  fmt.Println("It took brute force:", time.Since(start))
  start = time.Now()
  fmt.Println("Smallest number is:", lowestNumberNeeded(20))
  fmt.Println("It took prime factor:", time.Since(start))
}

func bruteForce() int{
  target := 0
  for i:= 3; 
    i %  2 != 0 || i % 3  != 0 || i %  4 != 0 || i % 5  != 0 ||
    i %  6 != 0 || i %  7 != 0 || i %  8 != 0 || i %  9 != 0 ||
    i % 10 != 0 || i % 11 != 0 || i % 12 != 0 || i % 13 != 0 ||
    i % 14 != 0 || i % 15 != 0 || i % 16 != 0 || i % 17 != 0 ||
    i % 18 != 0 || i % 19 != 0 || i % 20 != 0; i++ {
    target = i+1 //add 1 because loop stops before i++
  }
  return target
}

func lowestNumberNeeded(number int) int{
  //arrayNum will keep track of what values and what exponent we will need for prime factors
  arrayNum := map[int]int{1:1}
  for i:=2 ; i<=number; i = i+1 {
    for j:=1; j<=20; j = j+1 {
      if power(i,j) > number && isPrime(i){
        arrayNum[i] = j-1 // arrayNum = map[int]int{1:1, 2:3, 3:4}
        break
      }
    }
  }
  //iterate over the map and multiply to get the lowest number that can be evenly divided by all numbers
  total := 1
  for key, value := range arrayNum {
    total *= power(key, value)
  }
  return total //return lowest number
}

func power(a int, b int) int{
  //a is the base, b is the exponent
  value := 1
  for i:=0; i<b; i++ {
    value *= a
  }
  return value
}

func isPrime(n int) bool{
  var i int = 2 //start at 2 since prime number is a factor of itself and 1
  for ; i<n; i++ {
    if n%i == 0 { //check if n is divisible by any number other than itself
      return false //not a prime number if it is divisible by a number other than itself and 1
    }
  }
  return true
}





