package swicth 

import "fmt"

func swicth() {
  ages := map[string]int{}
  ages["Bob"] = 12
  
  switch {
  case ages["Bob"] < 18:
    fmt.Println("Bob cant vote")
  } case ages["Bob"] < 67:
     fmt.Println("Bob is not of retirement age")
    default:
      fmt.Println("Bob is of retirenment")
  }
}
