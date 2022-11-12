package main

import (
	"fmt"
	"net/http"
)

type pingHandler struct{}

func (p *pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
