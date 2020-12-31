import(
  "flag"
  "fmt"
)

func mainTest() {
  var err error = nil
  fmt.Println("Hello World!!")
  
  var flagVal string
  flag.StringVar(&flagVal, "args")
  
  flag.Parse()
  
  str, err := testCLI(flagVal)
  
  if err != nil {
    log.Panic(err)
  }
}

func testCLI(arg string) (string, error) {
  if len(arg) < 1 {
    err := errors.New("The argument length is less than 1")
    return "", err
  }
  
  return arg, nil
}
