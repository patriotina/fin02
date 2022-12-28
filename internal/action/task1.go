package action

import (
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
	fmt.Println(body)
	var ttt = model.ResponseArray{1, 2, 3}
	return moveArray(ttt, 3), nil
}

func moveArray(arrin []int, n int) (response model.ResponseArray) {
	response = make([]int, len(arrin))
	fmt.Println(arrin)
	for i := 0; i < len(arrin); i++ {
		if i+n >= len(arrin) {
			response[i+n-len(arrin)] = arrin[i]
		} else {
			response[i+n] = arrin[i]
		}
	}
	return response
}
