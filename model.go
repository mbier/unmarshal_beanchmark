package unmarshal_benchmark

type Model struct {
	String  string  `json:"string"`
	Float   float64 `json:"float"`
	Integer int64   `json:"integer"`
	//Date    time.Time `json:"date"`
	Boolean bool `json:"boolean"`
}

func NewModel() Model {
	return Model{
		String:  "string",
		Float:   3.33,
		Integer: 3,
		//Date:    time.Now(),
		Boolean: true,
	}
}
