// Euler Four
// Palindromic numbers read the same both ways ie. 9009, 8008, 1111
// Largest from product of 2-digit numbers is 9009 = 91 * 99
package main

import (
  "fmt"
  "strconv"
  "time"
)

func main() {
  fmt.Println("Euler 4")
  fmt.Println(isPalindromeInt(906609)) //if input is an integer
  fmt.Println(isPalindrome("2112")) //if input were to be a string
  start := time.Now()
  fmt.Println("Largest palindrome is:", largestPalindrome(1000))
  fmt.Println("Time taken:", time.Since(start))
}

//If input is a string type number
func isPalindrome(s string) bool{
  //reverse the string number
  r := make([]byte, len(s))
  for i := 0; i < len(s); i++ {
      r[i] = s[len(s)-1-i]
  }
  revString := string(r) //reverse is intially a string
  revInt, _ := strconv.Atoi(revString) //convert string to int
  origInt, _ := strconv.Atoi(s) //convert string to integer

  fmt.Println(origInt, revInt)
  return origInt == revInt //check if the reverse integer is equal to original
}

//if input is an integer type number
func reverseInt(a int) int{
  revNum := 0
  for ; a > 0 ; a = a/10 {
    revNum = 10 * revNum + a % 10
  }
  return revNum  
}

func isPalindromeInt(a int) bool{
  return a==reverseInt(a)
}

//find largest product 
func largestPalindrome(numDigits int) int{
  number := 0
  rangeNum := (numDigits-1)/10 //only need to scan for values from this to numDigits
  for i:=numDigits-1; i>=rangeNum; i-- {
    for j:=i; j>=rangeNum; j-- {
      if isPalindromeInt(j*i) && j*i>number{
        number = j*i
      }
    }
  }
  return number
}