package storage

import "errors"

type Conn struct {
	val bool
}

func (c *Conn) Open() error {
	if c.val {
		return errors.New("failed to open Conn")
	}
	c.val = true
	return nil
}

func (c *Conn) Close() error {
	if !c.val {
		return errors.New("failed to close")
	}
	c.val = false
	return nil
}

func (c *Conn) IsClose() bool {
	return !c.val
}
