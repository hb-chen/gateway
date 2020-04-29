package codec

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/grpclog"
)

var useNumber bool

type jsonCodec struct {
}

// UseNumber fix unmarshal Number(8234567890123456789) to interface(8.234567890123457e+18)
func UseNumber() {
	useNumber = true
}

var jsonpbMarshaler = &jsonpb.Marshaler{}
var jsonpbUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	grpclog.Infof("codec marshal: %+v", v)
	if m, ok := v.(json.Marshaler); ok {
		return m.MarshalJSON()
	}

	if pb, ok := v.(proto.Message); ok {
		grpclog.Infof("codec marshal proto message")
		s, err := jsonpbMarshaler.MarshalToString(pb)

		return []byte(s), err
	}

	grpclog.Infof("codec marshal json")
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	grpclog.Infof("codec unmarshal data: %v", string(data))
	if len(data) == 0 {
		return nil
	}
	if m, ok := v.(json.Unmarshaler); ok {
		return m.UnmarshalJSON(data)
	}

	if pb, ok := v.(proto.Message); ok {
		grpclog.Infof("codec unmarshal proto message")
		err := jsonpbUnmarshaler.Unmarshal(bytes.NewReader(data), pb)
		if err != nil {
			grpclog.Error(err)
		}
		return err
	}

	grpclog.Infof("codec unmarshal json")
	dec := json.NewDecoder(bytes.NewReader(data))
	if useNumber {
		dec.UseNumber()
	}
	return dec.Decode(v)
}

func (jsonCodec) Name() string {
	return CODEC_JSON
}
