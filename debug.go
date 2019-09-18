package upnp

import (
	"bytes"
	"io"
)

func debugWriter(data []byte, prefix string, writer io.Writer) {
	lines := bytes.Split(data, []byte{'\r', '\n'})
	out := "\n"
	for i, c := 0, len(lines); i < c; i++ {
		out += prefix + string(lines[i]) + "\n"
	}
	out = out[:len(out)-1]
	writer.Write([]byte(out))
}
