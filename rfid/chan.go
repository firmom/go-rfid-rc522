package rfid

type ReaderChan struct {
  reader RfidReader
  ch chan string
}

func NewReaderChan(reader RfidReader) (ReaderChan, error) {
  reducReader, err := NewReducer(reader)
  if err != nil {
    return nil, err
  }
  c := &ReaderChan{
    reader: reducReader,
  }
  go c.loop()
  return c, nil
}

func (c RfidChan) loop() {
  for {
    id, err := c.reader.ReadId()
    if err != nil {
      continue
    }
    c.ch <- id
  }
}

func (c RfidChan) ReadId() (string, error) {
  return <- c.ch
}

func (c RfidChan) getChan() (chan string) {
  return c.ch
}
