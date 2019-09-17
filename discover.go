package upnp

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

// Discover does an UPnP request, see: http://www.upnp-hacks.org/upnp.html
//
// The ifi parameter can be null although this is not recommended because the
// assignment depends on platforms and sometimes it might require routing
// configuration (see net.ListenMulticastUDP).
func Discover(queue chan<- *http.Response, timeout time.Duration, ifi *net.Interface, headers map[string]string, debug io.Writer) error {

	readers := newReaderPool()
	addr := &net.UDPAddr{IP: []byte{239, 255, 255, 250}, Port: 1900}
	request := newRequest(timeout, addr, headers)
	socket, err := net.ListenMulticastUDP("udp4", ifi, addr)

	if nil != err {
		return fmt.Errorf("failed to initialize UPnP conn: %s", err)
	}

	defer socket.Close()

	if err := socket.SetDeadline(time.Now().Add(timeout)); err != nil {
		return fmt.Errorf("failed to set deadline on UPnP conn: %s", err)
	}

	if err := request.Write(&writer{conn: socket, addr: addr, debug: debug}); err != nil {
		if e, ok := err.(net.Error); !ok || !e.Timeout() {
			return fmt.Errorf("failed to write request to UPnP conn: %s", err)
		}
		return nil
	}

	var wg sync.WaitGroup

	for {
		resp := make([]byte, 4096)
		n, _, err := socket.ReadFrom(resp)
		if err != nil {
			if e, ok := err.(net.Error); !ok || !e.Timeout() {
				return fmt.Errorf("UPnP read error: %s", err)
			}
			break
		}
		wg.Add(1)
		go parseResponse(&wg, resp[:n], readers, queue, debug)
	}

	wg.Wait()
	return nil
}
