package v1

func (d *device) Read() (Event, error) {
	buff := make([]byte, 128)
	n, err := d.port.Read(buff)
	if err != nil {
		return nil, err
	}
	cpbuf := make([]byte, n)
	copy(cpbuf, buff[:n])

	evt, err := ParseEvent(cpbuf)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
