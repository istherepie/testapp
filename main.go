package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func logger(req *http.Request) {
	var uri string = req.RequestURI
	var method string = req.Method
	log.Printf("%s %s", method, uri)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {

	logger(req)
	var message string = "This is a response from the testapp"
	fmt.Fprintf(w, "%s\n", message)
}

func headers(w http.ResponseWriter, req *http.Request) {
	logger(req)

	fmt.Fprintf(w, "Display headers for this request: \n")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Usage
	flag.Usage = func() {
		fmt.Printf("Usage of `%s`\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Flags
	var serverPort = flag.String("port", "8080", "Specify which port the server runs on")
	flag.Parse()

	// Info
	fmt.Printf("Webserver started and fully operational on port: %s\n", *serverPort)

	// Register routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/headers", headers)

	// Run
	var port = fmt.Sprintf(":%s", *serverPort)
	http.ListenAndServe(port, nil)
}
