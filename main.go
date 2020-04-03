package main

import (
    "fmt"
	"net/http"
	"os"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	var color string

	switch environ := os.Getenv("CANARY_ENVIRONMENT"); environ {
	case "release":
		color = "green"
	case "candidate":
		color = "red"
	case "baseline":
		color = "teal"
	default:
		color = "blue"
	}

	fmt.Fprintf(w, "<html><head></head><body><font color='%s'>%s</font></body></html>",color,os.Getenv("CANARY_ENVIRONMENT"))
}
