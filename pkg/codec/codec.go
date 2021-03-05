package codec

import (
	"google.golang.org/grpc/encoding"
)

const CODEC_JSON = "grpc-gateway/json"

func init() {
	encoding.RegisterCodec(jsonCodec{})
}
