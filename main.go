package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "The port to listen on.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		req := bytes.Buffer{}
		r.Write(&req)

		fmt.Printf("\n=== %s ===\n%s\n",
			time.Now().Format("15:04:05"),
			req.String(),
		)
	})

	fmt.Println("Listening for requests at 0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}
