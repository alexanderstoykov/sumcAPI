package helpers

import "encoding/json"

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", " ")
	println(string(b))
}
