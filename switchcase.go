package main 

import "fmt"

func main(){
  fmt.Println("Enter a number 1-5 to see it spelled in English: ")
  var input float64
  fmt.Scanf("%f", &input)

  var output string
  switch input {
   case 1: output = "One"
   case 2: output = "Two"
   case 3: output = "Three"
   case 4: output = "Four"
   case 5: output = "Five"
   default: output = "Unknown Number"
  }
  fmt.Println(output)
}