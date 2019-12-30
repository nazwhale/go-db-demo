package handler

import (
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Hello. This is our first Go web app on Heroku!")
}
