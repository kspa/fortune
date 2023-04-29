package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("fortune")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Fortune is:", stdout.String())
		fmt.Fprintf(w, "<h1>Fortune:</h1><p>%s</p>", stdout.String())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
