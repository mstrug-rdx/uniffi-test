package main

import (
  "github.com/mstrug-rdx/uniffi-test/test_package"
)

func main() {
  var tmp test_package.Byte = 1
  println("Some C declared type:", tmp)
}
