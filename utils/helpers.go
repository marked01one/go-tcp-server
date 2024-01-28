package utils

import (
	"fmt"
	"os"
)

func handleError(err error, msg string) {
	if err != nil {
		if msg == "" {
			fmt.Println(msg)
		}
		fmt.Println("ERROR TRACE:\n", err.Error())
		os.Exit(1)
	}
}
