// JSON stream reader.
package entry

import (
	"encoding/json"
	"io"
)

func NewJsonStream(reader io.Reader) (chan *json.RawMessage, chan error, chan struct{}) {
	dec := json.NewDecoder(reader)
	msgChan := make(chan *json.RawMessage)
	errChan := make(chan error)
	eofChan := make(chan struct{})

	go func() {
		for {
			m := new(json.RawMessage)
			if err := dec.Decode(m); err == io.EOF {
				eofChan <- struct{}{}
				break
			} else if err != nil {
				errChan <- err
			}

			msgChan <- m
		}
	}()

	return msgChan, errChan, eofChan
}
