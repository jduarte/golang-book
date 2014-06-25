package main

import "fmt"

func sum(args... int) int {
  total := 0
  for _, v := range args {
    total = total + v
  }

  return total
}

func half(i int) (bool, int) {
  return i % 2 == 0, (i / 2)
}

func greatest(args... int) int {
  g := 0
  for _, v := range args {
    if v > g {
      g = v
    }
  }

  return g
}

func main() {
  fmt.Println(sum(5,6,7,8,9))

  fmt.Println(half(1))
  fmt.Println(half(2))

  fmt.Println(greatest(1,2,3,4,5,6))
}
