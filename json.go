package unmarshal_benchmark

import (
	"encoding/json"
	"github.com/Jeffail/gabs/v2"
)

func marshalJson() ([]byte, error) {

	marshal, err := json.Marshal(NewModel())
	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func unmarshalJson(model []byte) (*Model, error) {

	var m Model

	err := json.Unmarshal(model, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func unmarshalJsonDynamic(model []byte) (*Model, error) {

	parseJSON, err := gabs.ParseJSON(model)
	if err != nil {
		return nil, err
	}

	return &Model{
		String:  parseJSON.Path("string").Data().(string),
		Float:   parseJSON.Path("float").Data().(float64),
		Integer: int64(parseJSON.Path("integer").Data().(float64)),
		//Date:    parseJSON.Path("date").Data().(string),
		Boolean: parseJSON.Path("boolean").Data().(bool),
	}, nil
}
