package main

import (
	"fmt"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return (err != nil)
}
