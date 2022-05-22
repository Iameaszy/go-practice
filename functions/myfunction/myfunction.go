package myfunction

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
  //name, age := myfunction.BareReturn()
  //fmt.Printf("{ Name: %s, Age: %d }", name, age);
  //fmt.Printf("result: %d\n", myfunction.VariadicFunction(2, 3, 4));
  //myfunction.BigSlowOperation()
  Defer()
}

func BareReturn() (name string, age int) {
  name = "Yusuf";
  age = 23;

  details := func (name string) string { return name;}
  fmt.Printf(details("Yusuf"))
  return
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
  const timeout = 1 * time.Minute
  deadline := time.Now().Add(timeout)
  for tries := 0; time.Now().Before(deadline); tries++ {
    _, err := http.Head(url)
    if err == nil {
      return nil // success
    }
    log.Printf("server not responding (%s); retrying...", err)
    time.Sleep(time.Second << uint(tries)) // exponential back-off
  }
  return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func VariadicFunction(vals ...int) int {
  total := 0;
  for _, value:= range vals {
    total += value;
  }
  return total;
}

func BigSlowOperation() {
  defer trace("bigSlowOperation")() // don't forget the extra parentheses
  // ...lots of work...
  time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
  start := time.Now()
  log.Printf("enter %s", msg)
  return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
