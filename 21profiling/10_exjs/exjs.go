package exjs

import "encoding/json"

var (
	JsonExample = []byte("{\"I\": 10}")
)

//easyjson:json
type A struct {
	I int
	B bool
}

func StandardMarshal() {
	a := &A{I: 15}
	json.Marshal(a)
}

func StandardUnmarshal() {
	a := &A{}
	json.Unmarshal(JsonExample, &a)
}

func EasyMarshal() {
	a := &A{I: 15}
	a.MarshalJSON()
}

func EasyUnmarshal() {
	a := &A{}
	a.UnmarshalJSON(JsonExample)
}
