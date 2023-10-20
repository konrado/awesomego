package data

type View struct {
	Data  interface{} `json:"data"`
	Extra interface{} `json:"extra,omitempty"`
}

type ListView struct {
	Data  []interface{} `json:"data"`
	Extra interface{}   `json:"extra,omitempty"`
}
