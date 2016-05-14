package rc522

import (
  "fmt"
  "testing"
)

func TestCreateAdapter(t *testing.T) {
  instance, err := NewRfidReader()
  if err != nil {
    t.Error(err)
  }
  if instance == nil {
    fmt.Errorf("instance can not be null")
  }
}
