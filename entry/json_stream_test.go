package entry_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/bcho/whitetree/entry"
)

func TestJsonStream(t *testing.T) {
	const jsonStream = `
                {"A": "b"}
        `
	type TestMsg struct {
		A string
	}

	msgChan, errChan, eofChan := entry.NewJsonStream(strings.NewReader(jsonStream))

	for {
		select {
		case <-eofChan:
			return
		case err := <-errChan:
			t.Errorf("json stream failed: %q", err)
			return
		case m := <-msgChan:
			msg := new(TestMsg)
			if err := json.Unmarshal(*m, msg); err != nil {
				t.Errorf("json stream failed: %q", err)
				return
			}
			if msg.A != "b" {
				t.Errorf("json stream failed: %s", msg.A)
				return
			}
		}
	}
}

func TestJsonStreamReturnsError(t *testing.T) {
	const jsonStream = `
                {"a": "b"
        `
	msgChan, errChan, eofChan := entry.NewJsonStream(strings.NewReader(jsonStream))

	for {
		select {
		case <-eofChan:
		case <-msgChan:
			t.Errorf("json stream should failed")
			return
		case <-errChan:
			return
		}
	}
}
