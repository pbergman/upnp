package upnp

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func newRequest(timeout time.Duration, addr net.Addr, headers map[string]string) *http.Request {
	request := new(http.Request)
	request.Method = "M-SEARCH"
	request.URL = &url.URL{Path: "*"}
	request.Proto = "HTTP/1.1"
	request.Host = addr.String()
	request.ProtoMajor = 1
	request.ProtoMinor = 1
	request.Header = http.Header(map[string][]string{
		//"ST":  {"ssdp:all"},
		"MAN": {"ssdp:discover"},
		"MX":  {strconv.Itoa(int(timeout / time.Second))},
	})
	for key, value := range headers {
		request.Header[key] = []string{value}
	}
	return request
}
