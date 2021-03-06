// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     message.avsc
 */

package avro

import (
	"github.com/actgardner/gogen-avro/container"
	"io"
)

type Message struct {
	Time    int64
	Channel string
	Nick    string
	Message string
}

func DeserializeMessage(r io.Reader) (*Message, error) {
	return readMessage(r)
}

func NewMessageWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := &Message{}
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

func NewMessage() *Message {
	v := &Message{}

	return v
}

func (r *Message) Schema() string {
	return "{\"fields\":[{\"name\":\"Time\",\"type\":\"long\"},{\"name\":\"Channel\",\"type\":\"string\"},{\"name\":\"Nick\",\"type\":\"string\"},{\"name\":\"Message\",\"type\":\"string\"}],\"name\":\"message\",\"type\":\"record\"}"
}

func (r *Message) Serialize(w io.Writer) error {
	return writeMessage(r, w)
}
