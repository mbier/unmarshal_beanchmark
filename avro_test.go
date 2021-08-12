package unmarshal_benchmark

import (
	"bytes"
	"github.com/hamba/avro"
	"github.com/linkedin/goavro/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var codec *goavro.Codec
var schema avro.Schema
var avroByte []byte
var avroByteGen *bytes.Buffer

func init() {
	codec, _ = createCodec()
	schema, _ = createSchema()
	avroByte, _ = marshalAvro(codec)
	avroByteGen, _ = marshalAvroGenerated()
}

func Benchmark_unmarshalAvroLinkedin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json, err := unmarshalAvroLinkedin(codec, avroByte)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalAvroGenerated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 := bytes.NewBuffer(avroByteGen.Bytes())
		json, err := unmarshalAvroGenerated(b1)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalAvroHamba(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json, err := unmarshalAvroHamba(schema, avroByte)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}
