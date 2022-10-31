package main

import (
	"encoding/json"
)

type OCRResult struct {
	numDown string
	numUp   string
}

func (O *OCRResult) NumDown() string {
	return O.numDown
}

func (O *OCRResult) NumUp() string {
	return O.numUp
}

func (O *OCRResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NumDown string `json:"num_down"`
		NumUp   string `json:"num_up"`
	}{
		NumDown: O.numDown,
		NumUp:   O.numUp,
	})
}

func (O *OCRResult) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		NumDown string `json:"num_down"`
		NumUp   string `json:"num_up"`
	}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	O.numDown = tmp.NumDown
	O.numUp = tmp.NumUp
	return nil
}
