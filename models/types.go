package models

type TimePicker struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (t *TimePicker) As(time string) string {
	return time
	//t.Value = time
	//jsonString, _ := json.Marshal(t)
	//return jsonString
}
