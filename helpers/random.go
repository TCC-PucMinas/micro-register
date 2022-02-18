package helpers

import (
	"fmt"
	"math/rand"
)

func RandomCodeClient() (string, error) {
	generateCode := fmt.Sprintf("%v%v", rand.Intn(600)-100, rand.Intn(600)-100)
	return GenerateHashSalt(generateCode)
}
