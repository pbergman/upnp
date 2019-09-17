package upnp

import (
	"bufio"
	"bytes"
	"sync"
)

type Reader struct {
	reader *bufio.Reader
	inner  *bytes.Buffer
	pool   *sync.Pool
}

func (r *Reader) Close() error {
	r.pool.Put(r)
	return nil
}

func (r *Reader) Reset(data []byte) error {
	r.inner.Reset()
	if _, err := r.inner.Write(data); err != nil {
		return err
	}
	_, _ = r.reader.Discard(r.reader.Buffered())
	return nil
}

func newReaderPool() *sync.Pool {
	readers := &sync.Pool{}
	readers.New = func() interface{} {
		reader := new(Reader)
		reader.inner = new(bytes.Buffer)
		reader.reader = bufio.NewReader(reader.inner)
		reader.pool = readers
		return reader
	}
	return readers
}
