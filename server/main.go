package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
