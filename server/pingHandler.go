package main

import (
	"net/http"
	"fmt"
)

type pingHandler struct{}

func (p *pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
