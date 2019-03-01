package main

import ("fmt"
"net/http")


func main() {
	fmt.Printf("ciao\n")

	handler := new http.Handler()
	

	http.ListenAndServe("localhost:8080",handler)
    
    
}
