package unmarshal_benchmark

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Benchmark_unmarshalAvro(b *testing.B) {
	codec, err := createCodec()
	assert.NoError(b, err)

	m, err := marshalAvro(codec)
	assert.NoError(b, err)

	for i := 0; i < b.N; i++ {
		json, err := unmarshalAvro(codec, m)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalAvroGenerated(b *testing.B) {

	m, err := marshalAvroGenerated()
	assert.NoError(b, err)

	for i := 0; i < b.N; i++ {
		b1 := bytes.NewBuffer(m.Bytes())
		json, err := unmarshalAvroGenerated(b1)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}
