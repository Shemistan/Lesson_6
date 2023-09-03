package storage

import "errors"

type conn struct {
	val bool
}

func NewConn() IConn {
	return &conn{
		val: false,
	}
}

func (c *conn) Open() error {
	if c.val {
		return errors.New("failed to open conn")
	}
	c.val = true
	return nil
}

func (c *conn) Close() error {
	if !c.val {
		return errors.New("failed to close")
	}
	c.val = false
	return nil
}

func (c *conn) IsClose() bool {
	return !c.val
}
