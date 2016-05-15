package rfid

import (
	"fmt"
)

type Reducer struct {
	reader   RfidReader
	oldvalue string
}

func NewReducer(reader RfidReader) (RfidReader, error) {
	return RfidReader(&Reducer{
		reader:   reader,
		oldvalue: "",
	}), nil
}

func (r *Reducer) ReadId() (string, error) {
	id, err := rfid.ReadId()
	if err != nil {
		return id, err
	}
	if id == r.oldvalue {
		return id, fmt.Errorf("RFID: duplicate id echo")
	}
	r.oldvalue = id
	return id, err
}
