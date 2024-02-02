package main

import "fmt"

func main() {
 // names := [3]string{"Bob", "John", "brad"}
 // var names [3]string 

// names[0] = "Bob"
// names[1] = "John"
// names[2] = "brad"

  names := []string{}

  names = append(names, "bob" , "John", "brad")
  fmt.Println(names)
}
