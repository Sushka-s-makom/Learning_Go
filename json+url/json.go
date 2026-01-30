package json_url

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func Jsonex() {
	// ну просто джонс фактически ничего не меняет
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.0)
	fmt.Println(string(fltB))

	res1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1_1, _ := json.Marshal(res1)
	fmt.Println(string(res1_1))

	res2 := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2_2, _ := json.Marshal(res2)
	fmt.Println(string(res2_2))
}
