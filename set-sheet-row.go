package main

import (
  "os"
  "fmt"
  "github.com/xuri/excelize/v2"
)

func main() {

  args := os.Args[1:]
  values := os.Args[4:]

  xslx, err := excelize.OpenFile(args[0])
  if err != nil {
    fmt.Println(err)
    return
  }

  err = xslx.SetSheetRow(args[1], args[2], &values)
  if err != nil {
    fmt.Println(err)
    return
  }

  err = xslx.Save()
  if err != nil {
    fmt.Println(err)
    return
  }

}