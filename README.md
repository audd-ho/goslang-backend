## Go-slang backend
- A locally hosted server that receives Goslang Code and returns the Abstract Syntax Tree in a JSON format of the Go code.
## Prerequisites
- Go installed locally. (can be found at https://go.dev/doc/install)
## Usage (to run the backend server)
```
$ git clone https://github.com/audd-ho/goslang-backend.git
$ cd goslang-backend
$ go run main.go
```
## Notes
- The program will run locally on port 8080
- Request to local server will be via http://localhost:8080/goslang-ast_json
- Receives a string of Go code in an object, in the format ```{"goslang_code": goslang_code}``` where the Go code is the variable goslang_code

## Usage (to send the request to the server)
- ```index.ts``` file can be copied to your typescript project and this module can be used to send Go code to a locally running backend server
## Notes
- ```ast_json_func.ts``` is a prototype for index.ts

## Used depedencies
- https://github.com/asty-org/asty (used for the parsing and conversion to AST-JSON by the backend)
