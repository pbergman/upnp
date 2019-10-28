package upnp

import (
	"io"
	"net"
)

type writer struct {
	conn  *net.UDPConn
	addr  *net.UDPAddr
	debug io.Writer
}

func (w *writer) Write(p []byte) (n int, err error) {
	if w.debug != nil {
		debugWriter(p, ">> ", w.debug)
	}
	return w.conn.WriteTo(p, w.addr)
}
