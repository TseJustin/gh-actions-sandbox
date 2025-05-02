package main

import (
  "encoding/json"
  "fmt"
  "log"
  "os"
  "os/exec"
)

func doError() {
  err := "ERROR MESSAGE"
  log.Println(err)
  log.Fatal(err)
}

func main() {
  url := "final-url"
  fmt.Printf(url)
}
