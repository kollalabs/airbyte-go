package airbyte

import (
	"bytes"
	"strings"
	"testing"
)

func TestStreamState(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	sw := newStateWriter(buff)

	sd := StreamDescriptor{
		Name:      "test",
		Namespace: "testnamespace",
	}

	streamData := struct {
		FieldA string
		FieldB int
	}{
		FieldA: "test",
		FieldB: 1,
	}

	err := sw(sd, streamData)
	if err != nil {
		t.Error(err)
	}

	expected := `{"type":"STATE","state":{"state_type":"STREAM","stream":{"stream_descriptor":{"name":"test","namespace":"testnamespace"},"stream_state":{"FieldA":"test","FieldB":1}}}}`
	o := buff.String()
	if strings.TrimSpace(o) != strings.TrimSpace(expected) {
		t.Errorf("Expected\n\t%s \ngot\n\t%s", expected, o)
	}

}
