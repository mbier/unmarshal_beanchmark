package unmarshal_benchmark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonByte []byte

func init() {
	jsonByte, _ = marshalJson()
}

func Benchmark_unmarshalJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		model, err := unmarshalJson(jsonByte)
		assert.NoError(b, err)
		assert.NotNil(b, model)
	}
}

func Benchmark_unmarshalJsonDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		model, err := unmarshalJsonDynamic(jsonByte)
		assert.NoError(b, err)
		assert.NotNil(b, model)
	}
}
