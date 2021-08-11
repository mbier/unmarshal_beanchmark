package unmarshal_benchmark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Benchmark_unmarshalProto(b *testing.B) {
	m, _ := marshalProto()
	for i := 0; i < b.N; i++ {
		json, err := unmarshalProto(m)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalProto_dynamicJhump(b *testing.B) {
	md, _ := loadDescriptor("m.Model", []string{"./proto/"})

	m, _ := marshalProto()
	for i := 0; i < b.N; i++ {
		json, err := unmarshalProtoDynamicJhump(md, m)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalProto_dynamicOfficial(b *testing.B) {
	md, _ := loadDescriptorOfficial()
	m, _ := marshalProto()
	for i := 0; i < b.N; i++ {
		json, err := unmarshalProtoDynamicOfficial(md, m)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}
