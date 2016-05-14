package main

import (
  "fmt"
  rc522 "github.com/firmom/go-rfid-rc522/rfid/rc522"
)

func main()  {
  rfid, err := rc522.NewRfidReader()
  if err != nil {
    fmt.Errorf(err)
    return
  }
  for {
    id := ReadId()
    fmt.Println(id)
  }
}
