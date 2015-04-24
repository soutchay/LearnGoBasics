package main

import (
  "fmt"
  "time"
)

const target int = 200

func main() {
  coinSlice := []int{1,2,5,10,20,50,100,200}
  start := time.Now()
  fmt.Println("There are", findWays(coinSlice), "ways to make", target)
  fmt.Println("Time for execution:", time.Since(start))
}

func findWays(coinSlice []int) int{
  //total - coin = newTotal
  numWays := [201]int{} 
  //store the number of ways to reach the total, each index will be initialized as 0
  numWays[0] = 1 //there is only one way to make any value of target with 1p
  // for i := 1; i<=target; i++ {
  //   for _, coin := range coinSlice {
  //     numWays[target] = numWays[current] + numWays[target-coinSlice[i]]
  //   }
  // }
  for _, coin := range coinSlice {
    for i := coin; i <=target; i++ {
      //numWays[value] = numWays[total-coinSlice[]] + numWays[]
      //look up previously calculated value
      numWays[i] = numWays[i-coin] + numWays[i]
      /*ie.
      when coin is 1
      numWays[1] = numWays[0] + numWays[1]
      numWays[2] = numWays[1] + numWays[2]
      when coin is 2
      numWays[2] = numWays[0] + numWays[2]
      numWays[3] = numWays[1] + numWays[3] 1+1 = 2
      numWays[4] = numWays[2] + numWays[4] 2+1 = 3
      */
    }
  }
  // fmt.Println(numWays[target])
  return numWays[target]
}