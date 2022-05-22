package myfunction

import (
	"fmt"
	"log"
	"os"
	"runtime"
)


func printStack() {
  var buf [4096]byte
  n := runtime.Stack(buf[:], false)
  os.Stdout.Write(buf[:n])
}

func f(x int) {
  fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
  defer fmt.Printf("defer %d\n", x)
  f(x - 1)
}


func Defer() {
  //defer printStack()
  //f(3)
  panicFunc(nil)
}

func panicFunc(param []string) {
  defer func() {
    if r := recover(); r != nil {
      panic(r)
    }
  }()
  
  if (param == nil) {
    panic(2)
  }
  log.Panicln("not panicking")
}