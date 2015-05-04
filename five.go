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
7) 
*/
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Euler 5")
}
