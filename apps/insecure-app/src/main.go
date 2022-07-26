package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "The Secret is: " + os.Getenv("THE_SECRET") )
    resp, err := http.Get(os.Getenv("THE_TARGET_URL"))
    if err != nil {
        fmt.Fprintf(w, "Error: "+ err.Error() )
    } else {
    fmt.Fprintf(w, "The Secret is: " + os.Getenv("THE_SECRET") + "<br/>Status Code is: " + fmt.Sprintf("%v",resp.StatusCode)  )
    }
}

