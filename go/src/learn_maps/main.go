package main

import "fmt"

func main() {
  ages := map[string]int{
    "bob" : 20,
    "john" : 30,
    "brad" : 92,
  }
  
  zip := map[string]int{}

  zip["bob"] = 28203
  zip["greg"] = 28209
  zip["don"] = 28233


  fmt.Println(zip)
  fmt.Println(ages)
}
