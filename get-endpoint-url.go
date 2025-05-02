package main

import (
  "fmt"
  "log"
)

func doError() {
  log.Println("ERROR MSG 1")
  log.Fatal("ERROR MSG 2")
  log.Fatal("ERROR MSG %s", "3")
}

func main() {
  url := "final-url"
  doError()
  fmt.Printf(url)
}
