package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "The port to listen on.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		headers := bytes.Buffer{}
		r.Header.Write(&headers)

		body, _ := ioutil.ReadAll(r.Body)

		fmt.Printf(`===
%s %s %s

Headers:
%s
Body: %s
`, time.Now().Format("15:04:05"), r.Method, r.URL.Path, headers.String(), string(body))
	})

	fmt.Println("Listening for requests at 0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}
