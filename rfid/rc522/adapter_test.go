package rc522

import "testing"

func TestCreateAdapter(t *testing.T) {
	instance, err := NewRfidReader()
	if err != nil {
		t.Error(err)
	}
	if instance == nil {
		t.Errorf("instance can not be null")
	}
}
