package upnp

import (
	"io"
	"net/http"
	"sync"
)

func parseResponse(wg *sync.WaitGroup, data []byte, pool *sync.Pool, queue chan<- *http.Response, debug io.Writer) {
	defer wg.Done()
	reader := pool.Get().(*Reader)
	defer reader.Close()
	if debug != nil {
		debugWriter(data, ">> ", debug)
	}
	_ = reader.Reset(data)
	request := &http.Request{}
	if response, err := http.ReadResponse(reader.reader, request); err == nil {
		queue <- response
	}
}
