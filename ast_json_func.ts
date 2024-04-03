const get_ast_json_from_goslang =  (goslang_code) => {
    const ast_json_from_goslang = (async (goslang_code) => {
    const res = await fetch("http://localhost:8080/goslang-ast_json", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify({"goslang_code": goslang_code})
    })
    const res_json = await res.json();
    return res_json;
    })(goslang_code);
    return ast_json_from_goslang
}

/*
let gslang_code = `package main

import (
    "fmt"
)

func GetFoo() {
    hello := make(chan int, 10)
    test := foo()
    fmt.Println(test)
}

func foo() int {
    return 0
}`;

get_ast_json_from_goslang(gslang_code).then((res) => {console.log(res); console.log(Object.keys(res));

console.log(res.Decls[1].Body.List); console.log(Object.keys(res.Decls[1].Body.List));})
*/