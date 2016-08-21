package main

import (
    "fmt"
    "net/http"
)

//fishing handler
func fishingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, this is the fishing page! %s", r.URL.Path[1:])  //write I love /pagename
}

//football handler
func footballHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, this is the football page! %s", r.URL.Path[1:]) 
}

func main() {
    http.HandleFunc("/fishingHandler", fishingHandler)   //register fishing page handler
    http.HandleFunc("/footballHandler", footballHandler) //register football page handler
    
    http.ListenAndServe(":8080", nil)   //listen on port 8080
}