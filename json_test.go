package unmarshal_benchmark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Benchmark_unmarshalJson(b *testing.B) {
	m, _ := marshalJson()

	for i := 0; i < b.N; i++ {
		json, err := unmarshalJson(m)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalJsonDynamic(b *testing.B) {
	m, _ := marshalJson()

	for i := 0; i < b.N; i++ {
		json, err := unmarshalJsonDynamic(m)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}
