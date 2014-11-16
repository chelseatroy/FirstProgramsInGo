package main

import "fmt"

func main(){
  fmt.Println("sum(4,6) returns: ", sum(4, 6))
  fmt.Println("half_and_even(4) returns: ")
  fmt.Println("maximum(5, 6, 8, 2, 1) returns: ", maximum(5, 6, 8, 2, 1))
  fmt.Println("fib(8)) returns: ", fib(8))
}

func sum(x int, y int) int{
  return x + y
}

func half_and_even(x int) (int, bool){
  return x/2, (x%2 == 0)
}

func maximum(args ...int) int{
  max := args[0]
  for i := 1; i < len(args); i++{
    if args[i] > max{
      max = args[i]
    }
  }
  return max
}

func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func fib(x int) int{
  if x <= 0{
    return 0
  } else {
    return fib(x-1) + fib(x-2)
  }
}

func swap(x *int, y *int) (*int, *int){
  temp = x
  x = y
  y = temp
  return x, y
}
