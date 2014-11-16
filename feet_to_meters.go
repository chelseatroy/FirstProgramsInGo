package main

import "fmt"

func main(){
  fmt.Print("Enter a number of feet: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 0.3048
  fmt.Println(output) 
}