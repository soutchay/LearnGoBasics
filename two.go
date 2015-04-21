//Euler 2

package main

import (
  "fmt"
)

func main() {
  fmt.Printf("Euler #2\n")
  fibonacci(5)
}

func fibonacci(n int) {
  fibMap := map[int]int{}
  // fmt.Println(fibMap[0])
  for i:=0; i<n+1; i++ {
    if i == 0 {
      fibMap[0] = 1
    } else if i == 1{
      fibMap[1] = 2
    }
    if fibMap[i] == 0 {
      fibMap[i] = fibMap[i-1] + fibMap[i-2]
    }

  }
  fmt.Println(fibMap)
}