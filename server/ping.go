package main

import (
	"net/http"
	"fmt"
)

type ping struct{}

func (p *ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
