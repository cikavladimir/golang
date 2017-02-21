package main

import "fmt"

func f() (int, int) {
  return 5, 6
}


func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}

func main() {
  x, y := f()
  fmt.Println(x)
  fmt.Println(y)


  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))
}