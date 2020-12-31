package main

import(
  "flag"
  "fmt"
  "errors"
  "log"
)

func main() {
  var err error = nil
  fmt.Println("Hello World!!")
  
  var flagVal string
  flag.StringVar(&flagVal, "args", "", "a string variable")
  
  flag.Parse()
  
  str, err := testCLI(flagVal)
  
  if err != nil {
    log.Panic(err)
  }

  fmt.Println("Argument:", str)
}

func testCLI(arg string) (string, error) {
  if len(arg) < 1 {
    err := errors.New("The argument length is less than 1")
    return "", err
  }
  
  return arg, nil
}
