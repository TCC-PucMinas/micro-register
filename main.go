package main

import (
	"fmt"

	"github.com/TCC-PucMinas/micro-register/helpers"
)

func main() {

	value, err := helpers.GenerateJwt("testando")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("access_token", value.AccessToken)
	fmt.Println("refresh_token", value.RefreshToken)
}
