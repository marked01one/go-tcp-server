package server

import (
	"fmt"
	"os"
)

/*
 * HELPER CONSTANTS -------------------------------------------------------------
 */
const HOST string = "localhost:8080"
const COLOR_RED string = "\033[31m"
const COLOR_GREEN string = "\033[32m"
const COLOR_NONE string = "\033[0m"

var Colors map[string]string = map[string]string{
	"200 OK":        COLOR_GREEN,
	"404 Not Found": COLOR_RED,
}

/*
 * HELPER TYPES -----------------------------------------------------------------
 */
type Response struct {
	server         string
	date           string
	content_length int
	content_type   string
	content        []byte
}

/*
 * HELPER FUNCTIONS -------------------------------------------------------------
 */
func logRes(method string, path string, status_code string) {
	fmt.Print(method, " ", path, "\t---\t", Colors[status_code], status_code, COLOR_NONE, "\n\n")
}

func panicErr(err error, msg string) {
	if err != nil {
		if msg == "" {
			fmt.Println(msg)
		}
		fmt.Println("ERROR TRACE:\n", err.Error())
		os.Exit(1)
	}
}
