package main

import (
"net/http"
"encoding/json"
"fmt"
)    

func main() {
    http.HandleFunc("/iot", httpHandler)            
    http.ListenAndServe(":8080", nil)
}

func httpHandler(w http.ResponseWriter, req *http.Request) { 
    var err error
    resp := map[string]interface{}{}
    if req.Method == "POST" {
        params := map[string]interface{}{}
        err = json.NewDecoder(req.Body).Decode(&params)
        if err != nil {
            fmt.Fprintf(w, err.Error()) 
        }
        resp, err = createPoint(params)
    } 
    if req.Method == "GET" {
        resp, err = getPoints()
    }
    enc := json.NewEncoder(w)
    enc.SetIndent("", "  ") 
    if err != nil {
        fmt.Println(err.Error())
    } else {
       if err := enc.Encode(resp); err != nil {
         fmt.Println(err.Error())
       } 
    }
}
