package rfid

type ReaderChan struct {
	reader RfidReader
	ch     chan string
}

func NewReaderChan(reader RfidReader) (*ReaderChan, error) {
	reducReader, err := NewReducer(reader)
	if err != nil {
		return nil, err
	}
	c := &ReaderChan{
		reader: reducReader,
		ch:     make(chan string, 1),
	}
	go c.loop()
	return c, nil
}

func (c *ReaderChan) loop() {
	for {
		id, err := c.reader.ReadId()
		if err != nil {
			continue
		}
		c.ch <- id
	}
}

func (c *ReaderChan) ReadId() (string, error) {
	return <-c.ch, nil
}

func (c *ReaderChan) GetChan() chan string {
	return c.ch
}
