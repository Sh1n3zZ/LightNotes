package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func Validate(token string) bool {
	res, err := Post("https://api.deeptrain.net/app/validate", map[string]string{
		"Content-Type": "application/json",
	}, map[string]interface{}{
		"password": viper.GetString("auth.access"),
		"token":    token,
		"hash":     Sha2Encrypt(token + viper.GetString("auth.salt")),
	})

	if err != nil || res == nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(res)
	return true
}
