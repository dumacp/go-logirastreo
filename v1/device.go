package v1

import (
	"fmt"
	"time"

	"github.com/tarm/serial"
)

type device struct {
	conf *serial.Config
	port *serial.Port
}

type Device interface {
	Open() error
	Close() error
	Conf() *serial.Config
	Read() (Event, error)
}

func NewDevice(port string, baud int, timeout time.Duration) Device {

	dev := &device{}

	conf := &serial.Config{
		Name:        port,
		Baud:        baud,
		ReadTimeout: timeout,
	}

	dev.conf = conf

	return dev
}

func (d *device) Open() error {
	if d.conf == nil {
		return fmt.Errorf("device conf is empty")
	}

	p, err := serial.OpenPort(d.conf)
	if err != nil {
		return err
	}
	d.port = p
	return nil
}

func (d *device) Close() error {

	if d.port == nil {
		return fmt.Errorf("device por is empty")
	}

	return d.port.Close()
}

func (d *device) Conf() *serial.Config {
	return d.conf
}
