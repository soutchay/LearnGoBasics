//Euler 2

package main

import (
  "fmt"
)

func main() {
  fmt.Printf("Euler #2\n")
  fibonacci(100)
  evenSum(fibonacci(50))
}

func fibonacci(n int) (fib map[int]int){
  fibMap := map[int]int{}
  for i:=0; i<n+1; i++ {
    if i == 0 {
      fibMap[0] = 1
    } else if i == 1{
      fibMap[1] = 2
    }
    if fibMap[i] == 0 { //if doesn't exist in fibMap
      fibMap[i] = fibMap[i-1] + fibMap[i-2]
    }
  }
  fmt.Println("Fibonacci number", n, ":", fibMap[n])
  return fibMap
}

func evenSum(a map[int]int) {
  fibSlice := []int{} //slice of the numbers that are less than 4million and even
  fibSum := 0
  for k, _ := range a {
    if a[k] < 4000000 && a[k]%2 == 0{
      fibSlice = append(fibSlice, a[k])
      fibSum += a[k]
    }
  }
  fmt.Println("The sum of even and less than 4million: ", fibSum)

}