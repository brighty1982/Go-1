package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])  //write I love /pagename
}

func main() {
    http.HandleFunc("/", handler)   //handle all requests to webroot with handler function
    http.ListenAndServe(":8080", nil)   //listen on port 8080
}