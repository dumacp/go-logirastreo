package v1

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
)

type event struct {
	header   string
	inputs1  int
	outputs1 int
	inputs2  int
	outputs2 int
	locks1   int
	locks2   int
	raw      []byte
}

type Event interface {
	Header() string
	Inputs1() int
	Outputs1() int
	Inputs2() int
	Outputs2() int
	Locks1() int
	Locks2() int
	RawResponse() []byte
}

func (e *event) RawResponse() []byte {
	return e.raw
}

func (e *event) Header() string {
	return e.header
}

func (e *event) Inputs1() int {
	return e.inputs1
}

func (e *event) Outputs1() int {
	return e.outputs1
}

func (e *event) Inputs2() int {
	return e.inputs2
}

func (e *event) Outputs2() int {
	return e.outputs2
}

func (e *event) Locks1() int {
	return e.locks1
}

func (e *event) Locks2() int {
	return e.locks2
}

func ParseEvent(data []byte) (Event, error) {

	if !bytes.Contains(data, []byte("AT$POST")) {
		return nil, fmt.Errorf("data don't contains header = %q", data)
	}

	evt := &event{}

	buf := bytes.NewBuffer(data)
	r := bufio.NewReader(buf)

	header, err := r.ReadBytes('=')
	if err != nil {
		return nil, fmt.Errorf("evt parse error = %s, data = %q", err, data)
	}
	header = header[:len(header)-1]

	evt.header = string(header)

	for i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		v, err := r.ReadBytes(',')
		if err != nil {
			return nil, fmt.Errorf("evt parse error = %s, data = %q", err, data)
		}
		switch i {
		case 4:
			if len(v) > 1 {
				inputs1, _ := strconv.Atoi(string(v[:len(v)-1]))
				evt.inputs1 = inputs1
			}
		case 5:
			if len(v) > 1 {
				outputs1, _ := strconv.Atoi(string(v[:len(v)-1]))
				evt.inputs1 = outputs1
			}
		case 6:
			if len(v) > 1 {
				inputs2, _ := strconv.Atoi(string(v[:len(v)-1]))
				evt.inputs2 = inputs2
			}
		case 7:
			if len(v) > 1 {
				outputs2, _ := strconv.Atoi(string(v[:len(v)-1]))
				evt.outputs2 = outputs2
			}
		case 8:
			if len(v) > 1 {
				locks1, _ := strconv.Atoi(string(v[:len(v)-1]))
				evt.locks1 = locks1
			}
		case 9:
			if len(v) > 1 {
				locks2, _ := strconv.Atoi(string(v[:len(v)-1]))
				evt.locks2 = locks2
			}
		}

	}

	evt.raw = data
	return evt, nil

}
