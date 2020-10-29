package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go supports encoding and decoding JSON, including to and from custom
// and built-in data types

type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded and decoded in json
// Fields must start in caps to be exported
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(12)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(23.43)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["apple","peach"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println("num:", num)
	strs := dat["strs"].([]interface{})
	fmt.Println("strs:", strs)
	fmt.Println("strs[0]:", strs[0].(string))
	fmt.Println("strs[1]:", strs[1].(string))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res2 := response2{}
	json.Unmarshal([]byte(str), &res2)
	fmt.Println(res2, res2.Fruits[1])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
