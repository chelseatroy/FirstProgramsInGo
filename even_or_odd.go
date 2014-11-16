package main    

import "fmt"

func main(){
  for count := 1; count <= 10; count++ {
    fmt.Print(count)
    if count%2 == 0 {
      fmt.Println(" even")
    } else {
      fmt.Println(" odd")
    }
  }
}