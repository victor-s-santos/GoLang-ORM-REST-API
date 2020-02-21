package main

import (
	"fmt"
	"net/http"
)

func Users(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Endpoint users")
}
