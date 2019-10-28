package upnp

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
)

func isTimeout(err error) bool {

	e, ok := err.(net.Error)

	if !ok {
		return false
	}

	return e.Timeout()
}

func read(socket *net.UDPConn) ([]byte, error) {
	resp := make([]byte, 4096)
	size, _, err := socket.ReadFrom(resp)
	if err != nil {
		if false == isTimeout(err) {
			return nil, fmt.Errorf("UPnP read error: %s", err)
		} else {
			return nil, err
		}
	}
	return resp[:size], nil
}

func parse(wg *sync.WaitGroup, data []byte, pool *sync.Pool, handler func(resp *http.Response), debug io.Writer) {
	defer wg.Done()
	reader := pool.Get().(*Reader)
	defer reader.Close()
	_ = reader.Reset(data)
	request := &http.Request{}
	if response, err := http.ReadResponse(reader.reader, request); err == nil {
		if debug != nil {
			debugWriter(data, "<< ", debug)
		}
		handler(response)
	} else {
		if debug != nil {
			debugWriter(data, "<< [D] ", debug)
		}
	}
}
