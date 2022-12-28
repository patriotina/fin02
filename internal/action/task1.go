package action

import (
	"encoding/json"
	"fin02/internal/model"
	"fmt"
	"io"
	"net/http"
)

func GetTask1() (model.ResponseArray, error) {
	resp, err := http.Get("https://kuvaev-ituniversity.vps.elewise.com/tasks/Циклическая ротация")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	//byt := []byte(body)
	//var dat map[string]interface{}

	type bodyst []interface{}
	type bodysts []bodyst
	var bodys bodysts
	json.Unmarshal(body, &bodys)

	fmt.Println(bodys[0])

	var arrin model.ResponseArray
	var N int

	for i := 0; i < len(bodys); i++ {
		arrin := bodys[0]
		N := bodys[1]
		fmt.Println(moveArray(arrin, N))
	}

	var ttt = model.ResponseArray{1, 2, 3}
	return moveArray(ttt, 3), nil
}

func moveArray(arrin []int, n int) (response model.ResponseArray) {
	response = make([]int, len(arrin))
	fmt.Println(arrin)
	if n > len(arrin) {
		n = n % len(arrin)
	}
	for i := 0; i < len(arrin); i++ {
		if i+n >= len(arrin) {
			response[i+n-len(arrin)] = arrin[i]
		} else {
			response[i+n] = arrin[i]
		}
	}
	return response
}
