package main

import "fmt"

func main() {
  ages := map[string]int{}

  ages["john"] = 33
  ages["chuck"] = 21
  ages["harris"] = 24
  ages["pete"] = 43

  if ages["john"] < 70 {
  
    fmt.Println("John is", ages["john"], "they are not of the retirment age")

 } else if ages["john"] < 18 {
   fmt.Println(" Cant vote")
 } else {
   fmt.Println("John is of age") 
 }
}
