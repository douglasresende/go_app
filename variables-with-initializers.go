package main

import "fmt"

var i, j int = 1, 2

func variables_with_initializers() {
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}
