package main 

import "fmt"

func main() {

  arr := []string{}

  arr = append(arr, "Bod", "did","get","sow","kid")

  for _, name := range arr { 
    if name == "did" {
      fmt.Println(name)
     } else {
       fmt.Println("Not it")
   }
  }

}
