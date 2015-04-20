//Euler 1
package main

import (
  "fmt"
)

func main() {
  fmt.Printf("Euler # 1\n")
  sum := sumArray(threeOrFive(1000))
  fmt.Println("Sum is: ", sum)
}

func threeOrFive(a int) (numb []int){
  numbers := []int{}
  for x := 0; x<a; x++ {
    if x%3 == 0 {
      numbers = append(numbers, x)
    } else if x%5 == 0 {
      numbers = append(numbers, x)
    }
  }
  return numbers
}

func sumArray(a []int) int{
  sum := 0
  for i:=0; i<len(a); i++ {
    sum += a[i]
  }
  return sum
}