package unmarshal_benchmark

import (
	"bytes"
	"github.com/hamba/avro"
	"github.com/linkedin/goavro/v2"
	stub "github.com/mbier/unmarshal_benchmark/avro"
	"io"
	"io/ioutil"
	"time"
)

func createCodec() (*goavro.Codec, error) {
	file, err := ioutil.ReadFile("model.avsc")
	if err != nil {
		return nil, err
	}

	return goavro.NewCodec(string(file))
}

func createSchema() (avro.Schema, error) {
	file, err := ioutil.ReadFile("model.avsc")
	if err != nil {
		return nil, err
	}

	return avro.Parse(string(file))
}

func marshalAvro(codec *goavro.Codec) ([]byte, error) {
	m := map[string]interface{}{
		"string":  "string",
		"float":   3.33,
		"integer": 3,
		"date":    time.Now().String(),
		"boolean": true,
	}

	return codec.SingleFromNative(nil, m)
}

func unmarshalAvroLinkedin(codec *goavro.Codec, model []byte) (*Model, error) {
	m := &Model{}

	native, _, err := codec.NativeFromSingle(model)
	if err != nil {
		return nil, err
	}

	d := native.(map[string]interface{})

	m.String = d["string"].(string)
	m.Float = float64(d["float"].(float32))
	m.Integer = int64(d["integer"].(int32))
	//m.Date = d["string"].(string)
	m.Boolean = d["boolean"].(bool)

	return m, nil
}

func unmarshalAvroGenerated(model io.Reader) (*Model, error) {
	m := &Model{}

	deserializeModel, err := stub.DeserializeModel(model)
	if err != nil {
		return nil, err
	}

	m.String = deserializeModel.String
	m.Float = float64(deserializeModel.Float)
	m.Integer = int64(deserializeModel.Integer)
	//m.Date = deserializeModel.Date
	m.Boolean = deserializeModel.Boolean

	return m, nil
}

func marshalAvroGenerated() (*bytes.Buffer, error) {
	var buf bytes.Buffer

	model := stub.NewModel()
	model.String = "string"
	model.Float = 3.33
	model.Integer = 3
	model.Boolean = true

	err := model.Serialize(&buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func unmarshalAvroHamba(schema avro.Schema, model []byte) (*Model, error) {
	m := &Model{}

	err := avro.Unmarshal(schema, model, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
