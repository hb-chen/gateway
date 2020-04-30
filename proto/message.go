package proto

import (
	"io"
	"io/ioutil"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/grpclog"
)

type Message struct {
	data []byte
}

// proto.Message
// "github.com/golang/protobuf/proto"
func (m *Message) ProtoMessage() {}

func (m *Message) Reset() {
	*m = Message{}
}

func (m *Message) String() string {
	return string(m.data)
}

// json.Marshaler
// "encoding/json"
func (m *Message) MarshalJSON() ([]byte, error) {
	return m.data, nil
}

func (m *Message) UnmarshalJSON(data []byte) error {
	m.data = data
	return nil
}

// gRPC gateway marshaler
// "github.com/grpc-ecosystem/grpc-gateway/runtime"
func (m *Message) Marshal(v interface{}) ([]byte, error) {
	grpclog.Infof("message marshal data: %v", m.data)
	return m.data, nil
}

func (m *Message) Unmarshal(data []byte, v interface{}) error {
	grpclog.Infof("message unmarshal data: %v", data)
	m.data = data
	return nil
}

func (m *Message) NewDecoder(r io.Reader) runtime.Decoder {
	return runtime.DecoderFunc(func(value interface{}) error {
		buffer, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		grpclog.Infof("message decoder data: %v", buffer)
		return m.Unmarshal(buffer, value)
	})
}

func (m *Message) NewEncoder(w io.Writer) runtime.Encoder {
	return runtime.EncoderFunc(func(value interface{}) error {
		buffer, err := m.Marshal(value)
		if err != nil {
			return err
		}
		grpclog.Infof("message encoder data: %v", buffer)
		_, err = w.Write(buffer)
		if err != nil {
			return err
		}

		return nil
	})
}

func (m *Message) ContentType() string {
	return "application/json"
}

func NewMessage(data []byte) *Message {
	return &Message{data}
}
