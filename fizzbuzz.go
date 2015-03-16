package main

import "fmt"
var p = fmt.Println
func main(){
  for i := 1; i < 100; i++ {
    switch{
      case i % 3 == 0 && i % 5 == 0:
        p("FizzBuzz")
      case i % 3 == 0:
        p("Fizz")
      case i % 5 == 0:
        p("Buzz")
      default:
        p(i)
    }
  }
}
