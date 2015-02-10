package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeDisplay(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Local()
	layout := "Jan 2, 2006 at 3:04pm (MST)"
	fmt.Fprintf(w, "The local time here is: %s \n", t.Format(layout))
}

func main() {
	http.HandleFunc("/", timeDisplay)
	http.ListenAndServe(":8080", nil)
}
