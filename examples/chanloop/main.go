package main

import (
	"fmt"
	"time"

	rfid "github.com/firmom/go-rfid-rc522/rfid"
	rc522 "github.com/firmom/go-rfid-rc522/rfid/rc522"
)

func main() {
	reader, err := rc522.NewRfidReader()
	if err != nil {
		fmt.Println(err)
		return
	}
	readerChan, err := rfid.NewReaderChan(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	rfidChan := readerChan.GetChan()
	doSomethingCount := 0
	for {
		select {
		case id := <-rfidChan:
			fmt.Println(id)
		default:
			doSomethingCount++
			fmt.Println(" ... ", doSomethingCount)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
