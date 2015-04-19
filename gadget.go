package main

import (
  "fmt"
  //"io/ioutil" // Implements some I/O utility functions.
  m "math"    // Math library with local alias m
  //"net/http"  // web server
  //"strconv"   // string conversions
)

func main() {
  fmt.Printf("hello, world\n")
  beyondHello()
  learnTypes()
  learnArrays()
  learnMaps()
  learnFlowControl()
}

func beyondHello() {
  var x int
  x = 3 //variable assignment
  y := 4 // short declaration to infer type, declare, and assign
  summ, prodd := learnMultiple(x, y) //function returns 2 values

  fmt.Println("sum:", summ, "prod:", prodd)
}

func learnMultiple(x, y int) (sum, prod int){
  return x + y, x * y  //return two values as sum and prod respectively as integers
}

func learnTypes() {
  str := "give me a string" // string
  s2 := `a "raw" string literal 
  with line breaks` //using ` has line breaks
  g := 'Σ' // int32
  f := 3.14195 //float64, 64bit floating point integer
  c := 3 + 4i //complex 128

  var p float64

  //underscore allows you to use the variable and discard its variable
  //unused variables will produce an error otherwise
  _ = p
  fmt.Println("str:", str, "\ns2:", s2, "\ng:", g,"\nf:", f, "\nc:", c)
}

//possible to have returned named values
//z is implicit because we named it earlier
func learnNamedReturns(x,y int) (z int){
  z = x * y
  return
}

func learnArrays() {
  //arrays have fixed size
  var a4 [4]int //array of 4 initialized with 0
  a5 := [...]int{1,2,3,4,5} //array with fixed size 5

  //Slices have dynamic size
  a6 := []int{1,2,3}
  a7 := make([]int, 4) //Slice of 4 ints at 0

  //can append to slice since they are dynamic
  s := []int{3,2,1}
  s2 := append(s, 0, -1, -2)

  fmt.Println(a4, a5, a6, a7, s, s2)

  s = append(s, []int{8,9,10}...) //Slice literal argument
  fmt.Println(s)
}

func learnMaps() {
  //maps are similar to objects in javascript and dictionaries in python; associative array
  m := map[string]int{"three": 3, "four": 4}
  m["one"] = 1

  fmt.Println(m)
}

func expensiveComputation() float64 {
    return m.Exp(10)
}

func learnFlowControl() {
  if 1+1==2 {
    fmt.Println("true")
  }
  if 1+3!=0 {
    fmt.Println("false")
  } else {
    fmt.Println("true")
  }
  //switch cases for multiple if statements
  x := 42.0
  switch x {
  case 0:
    fmt.Println("0")
  case 1:
    fmt.Println("1")
  case 42:
    fmt.Println("42") //case does not fall through
  case 43:
    fmt.Println("43")
  }

  //basic for loop
  for x:= 0; x<5; x++ {
    fmt.Println("iterate over x", x+1)
  }
  
  a :=[]int{1,2,3,4,5}
  //iterate over an array or slice
  for index, value := range a {
    fmt.Println("iterating over slice...\n", index, value)
  }

  //iterate over map
  for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
      // for each pair in the map, print key and value
      fmt.Printf("key=%s, value=%d\n", key, value)
  }
  // assign y then compare it to x, then make x equal to y
  // x= 42
  if y := expensiveComputation(); y>x {
    x = y
  }
  // x = e^10

  //function literals are closures
  xBig := func() bool {
    return x > 10000
  }
  fmt.Println("xBig?:", xBig())
  x = 1.3e3
  fmt.Println("xBig?:", xBig())

  //func literals can be used as arguments for functions
  fmt.Println("Add + double numbers: ",
    func(a,b int) int{
      return (a+b) * 2
    }(10,2))
  goto stuff

  fmt.Println("will not print because of goto stuff")

stuff:
  learnFunctionFactory()
  learnDefer()
  learnInterfaces()
  moreInterfaces()

}

//Function returning functions (sentenceFactory)
func learnFunctionFactory() {
  fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

  //more practical way to write the above:
  d := sentenceFactory("summer")
  fmt.Println(d("A beautiful", "day!"))
  fmt.Println(d("A lazy", "afternoon!"))  
}
func sentenceFactory(mystring string) func(before, after string) string {
  return func(before, after string) string {
    //Sprintf formats according to a format specifier and returns the resulting string
    return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
  }
}

func learnDefer() {
  // Deferred statements are executed just before the function returns.
  defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
  defer fmt.Println("\nThis line is being printed first because")
  // Defer is commonly used to close a file, so the function closing the
  // file stays close to the function opening the file.
  return
}

//Interfaces!
//interface type with method String
type Stringer interface{
  String() string
}
//struct type with two fields x and y as integers
type pair struct {
  x, y int
}
//define method on type pair which implements Stringer
func (p pair) String() string { // p is the receiver
  return fmt.Sprintf("(x: %d, y: %d)", p.x, p.y)
}


func learnInterfaces() {
  //initiliaze p to pair struct with x as 3 and y as 4
  p := pair{3, 4}
  fmt.Println(p.String()) // turning x and y from pair p to strings

  var i Stringer //declaring i as interface type Stringer
  i = p //works because pair implements Stringer from above
  fmt.Println("same as above", i.String())

  fmt.Println(p) //don't need to call String() method since Println does
  fmt.Println(i)
}

//More Interface
type geometry interface {
  area() float64
  perim() float64
}

type square struct {
  width, height float64
}
type circle struct {
  radius float64
}
type triangle struct {
  base, height float64
}

//Functions for squares
func (s square) area() float64 {
  return s.width * s.height
}
func (s square) perim() float64 {
  return 2*(s.width + s.height)
}

//Functions for circles
func (c circle) area() float64 {
  return c.radius * c.radius * m.Pi
}
func (c circle) perim() float64 {
  return c.radius * 2 * m.Pi
}

//Functions for right triangles
func (t triangle) area() float64 {
  return 0.5 * t.base * t.height
}
func (t triangle) perim() float64 {
  return t.base + t.height + m.Sqrt(t.base*t.base + t.height*t.height)
}

func measure(g geometry) {
  fmt.Println(g)
  fmt.Println("Area:", g.area())
  fmt.Println("Perimeter:", g.perim())
}

//putting it all together for squares, circles, triangles
func moreInterfaces() {
  s := square{width: 3, height: 4} //make square s with width and height
  c := circle{radius: 5} //make circle c with radius
  t := triangle{base: 5, height: 10}

  measure(s)
  measure(c)
  measure(t)
}
















