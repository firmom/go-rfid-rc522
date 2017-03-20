package rfid

type RfidReader interface {
	ReadId() (string, error)
}
