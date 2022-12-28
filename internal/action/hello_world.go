package action

import (
	"errors"
	"fin02/internal/model"
)

func GetHelloWorld() (model.HelloWorld, error) {
	return model.HelloWorld{HelloWorld: true}, errors.New("some error happens sometimes")
}
