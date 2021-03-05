package codec

import (
	"bytes"
	"encoding/json"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var useNumber bool

type JsonCodec struct {
	jsonCodec
}

type jsonCodec struct {
}

// UseNumber fix unmarshal Number(8234567890123456789) to interface(8.234567890123457e+18)
func UseNumber() {
	useNumber = true
}

var jsonpbMarshaler = &protojson.MarshalOptions{EmitUnpopulated: true}
var jsonpbUnmarshaler = &protojson.UnmarshalOptions{AllowPartial: true}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	grpclog.Infof("codec marshal: %+v", v)
	if m, ok := v.(json.Marshaler); ok {
		return m.MarshalJSON()
	}

	if pb, ok := v.(proto.Message); ok {
		grpclog.Infof("codec marshal proto message")
		s, err := jsonpbMarshaler.Marshal(pb)

		return s, err
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

		err := jsonpbUnmarshaler.Unmarshal(data, pb)
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
