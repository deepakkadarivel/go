package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int 	`json:"page"`
	Fruits []string	`json:"fruits"`
}

func main() {

	// Encoding basic data types to JSON string

	bloB, _ := json.Marshal(true)
	fmt.Println(string(bloB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.24)
	fmt.Println(string(floatB))

	stringB, _ := json.Marshal("Hello")
	fmt.Println(string(stringB))

	slcB, _ := json.Marshal([]string{"apple", "orange", "watermelon", "grapes"})
	fmt.Println(string(slcB))

	mapB,_ := json.Marshal(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	fmt.Println(string(mapB))

	// Encoding struct datatypes to json

	res1D := Response1{1,[]string{"apple", "orange", "grapes", "mango"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := Response2{2, []string{"cat", "dog", "fish", "parrot"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//Decoding JSON Data to GO Values

	// Generic data structure
	byt := []byte(`{"num":1, "str":["a", "b"]}`)

	var data map[string]interface{}

	if error := json.Unmarshal(byt, &data); error != nil {
		panic(error)
	}
	fmt.Println(data)

	num := data["num"].(float64)
	fmt.Println("num : ", num)

	strs := data["str"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("str 1 : ", str1)

	// Unmarshal ex 2
	str := `{"page": 3, "fruits": ["apple", "orange"]}`

	res2E := Response2{}

	if error2 := json.Unmarshal([]byte(str), &res2E); error2 != nil {
		panic(error2)
	}
	fmt.Println("Res 2 : ", res2E)

	fruit1 := res2E.Fruits[0]
	fmt.Println("Fruit 1 : ", fruit1)

	// Stream JSON encodings directly to os.Stdout
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 1, "orange": 2}
	enc.Encode(d)

}
