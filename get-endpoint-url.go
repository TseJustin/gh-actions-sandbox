package main

import (
  "fmt"
  "log"
)

func doError() {
  err := "ERROR MESSAGE"
  log.Println(err)
  log.Fatal(err)
}

func main() {
  url := "final-url"
  doError()
  fmt.Printf(url)
}
