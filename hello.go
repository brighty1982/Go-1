package main

import (
    "fmt"
    "net/http"
)

func fishingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, this is the fishing page! %s", r.URL.Path[1:])  //write I love /pagename
}

func footballHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, this is the football page! %s", r.URL.Path[1:])  //write I love /pagename
}

func main() {
    http.HandleFunc("/fishingHandler", fishingHandler)   //fishing page handler
    http.HandleFunc("/footballHandler", footballHandler)
    
    http.ListenAndServe(":8080", nil)   //listen on port 8080
}