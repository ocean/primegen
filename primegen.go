package main

import (
  "fmt"
  "math"
)

// Start at 3 as we know 1 and 2 are prime.
var n float64 = 3

// Initial halving divisor
var div float64 = 2

// Prime checking divisor
var j float64 = 2

func main() {
  for i := 0; i < 20; i++ {
    var half float64 = math.Floor(n / div)
    for j < half {
      if math.Mod(n, j) == 0 {
        //fmt.Printf("Number is %v, %v, %v\n", n, h, j)
        j++
        break
        // continue
      } else {
        // continue
        fmt.Printf("Number is %v, %v, %v\n", n, half, j)
        j++
      }
      // for math.Mod(n, j) > 0 {
      //   j++
      //   // break
      // }
      // fmt.Printf("Number is %v, %v, %v\n", n, half, j)
      // continue
      // break
      // j++
    }
    // if j == half {
    //   fmt.Printf("Number is %v, %v, %v\n", n, half, j)
    //   continue
    // }
    // fmt.Printf("Number is %v, %v, %v\n", n, half, j)
    n += 2
  }
}

