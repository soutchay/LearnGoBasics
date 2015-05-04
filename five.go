/*Euler 5
2520 is the smallest number that can be divided by each 
of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by 
all of the numbers from 1 to 20?
*/

/*Notes: 
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
  2,3,4,5,6,7,8 : 2*2*3*5*7*2


*/
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Euler 5")
}
