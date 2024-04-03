package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/asty-org/asty/asty"
)

type input_file_data struct {
	Goslang_code string
}

func main() {

	// set a HTTP request handle function for path /greeting and registrate it
	http.HandleFunc("/goslang-ast_json", func(rw http.ResponseWriter,
		req *http.Request) {
		var input_file input_file_data
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&input_file)
		if err != nil {
			panic(err)
		}
		fo, err := os.Create("input_goslang_code.go")
		if err != nil {
			panic(err)
		}
		fo.Write([]byte(input_file.Goslang_code))
		if err := fo.Close(); err != nil {
			panic(err)
		}
		options := asty.Options{
			WithImports:    false,
			WithComments:   false,
			WithPositions:  false,
			WithReferences: false,
		}
		asty.SourceToJSON("input_goslang_code.go", "output_goslang_ast_json.json", "	", options)
		fileBytes, _ := os.ReadFile("output_goslang_ast_json.json")
		var output_file_ast_json map[string]interface{}
		err_json := json.Unmarshal(fileBytes, &output_file_ast_json)
		if err_json != nil {
			panic(err_json)
		}
		rw.Write(fileBytes)
	})

	// print out the server is going to start and show the time
	log.Println("Starting server....")

	// create server at localhost:8080 and using tcp as the network
	listener, err := net.Listen("tcp", ":8080")

	// if recieve error, record it and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// setup HTTP connection for the listener of the server
	http.Serve(listener, nil)

}
