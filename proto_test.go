package unmarshal_benchmark

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

var protoByte []byte
var md *desc.MessageDescriptor
var mdOfficial protoreflect.MessageDescriptor

func init() {
	protoByte, _ = marshalProto()
	md, _ = loadDescriptor("m.Model", []string{"./proto/"})
	mdOfficial, _ = loadDescriptorOfficial()
}

func Benchmark_unmarshalProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json, err := unmarshalProto(protoByte)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalProto_dynamicJhump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json, err := unmarshalProtoDynamicJhump(md, protoByte)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}

func Benchmark_unmarshalProto_dynamicOfficial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json, err := unmarshalProtoDynamicOfficial(mdOfficial, protoByte)
		assert.NoError(b, err)
		assert.NotNil(b, json)
	}
}
