package common

import (
	"fmt"
	"net/http"
)

func InfoLog(str string) {
	fmt.Println(str)
}

func InfoResLog(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, str)
}
