package main

import (
  "fmt"
  rc522 "github.com/firmom/go-rfid-rc522/rfid/rc522"
)

func main()  {
  rfid, err := rc522.NewRfidReader()
  if err != nil {
    fmt.Println(err)
    return
  }
  for {
    id, err := rfid.ReadId()
    if err != nil {
      continue
    }
    fmt.Println(id)
  }
}
