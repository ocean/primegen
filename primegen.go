package main

import (
  "fmt"
  "math"
)

var n float64 = 0

func main() {
  for i := 0; i < 20; i++ {
    var div float64 = 2
    var h float64 = math.Floor(n / div) + 1
    var j float64 = 2
    for j = 2; j < h; j++ {
      if math.Mod(n, j) != 0 {
        //fmt.Printf("Number is %v, %v, %v\n", n, h, j)
        //break
        continue
      } else {
        fmt.Printf("Number is %v, %v, %v\n", n, h, j)
      }
    }
    n++
  }
}
  