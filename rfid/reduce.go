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
	id := r.reader.ReadId()
}
