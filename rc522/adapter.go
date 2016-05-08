package rc522

// #cgo LDFLAGS: -L${SRCDIR}/libs -lbcm2835
// #include <rc522_public.h>
import (
  "C"
  "fmt"
  rfid "github.com/firmom/go-rfid-rc522"
)

var (
  instance = nil //it is singleton, it is not possible to connect multiple rc522 readers to one rpi
)

type RfidReader struct {}

func NewRfidReader() (rfid.RfidReader, error) {
  if instance == nil {
    errCode, err := C.InitRC522RfidReader(void)
    if err != nil {
      return err
    }
    if errCode != 0 {
      return nil, fmt.Errorf("Init rc522 fail with code: ", errCode)
    }
    instance = rfid.RfidReader(RfidReader{})
  }
  return instance, nil
}

func (r RfidReader) ReadId() (string, error) {
  response, err := C.ReadIdByRC522()
  if err != nil {
    return err
  }
  if response.ErrorCode != C.NO_ERROR {
    return "", fmt.Errorf("Read rc522 fail with code: ", errCode)
  }
  if response.Status != C.NO_TAG_STATUS {
    return "", nil
  }
  return C.GoString(response.id), nil
}
