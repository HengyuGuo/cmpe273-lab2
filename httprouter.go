package main

import (
    "encoding/json"
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

type(
    User struct{
        Name string `json:"name"`
    }
)

type(
    Res struct{
        Greet string `json:"greeting"`
    }
)

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func hellojson(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    u := User{}
    json.NewDecoder(req.Body).Decode(&u)
    r := Res{}
    r.Greet = "Hello, " + u.Name + "!"
    uj, _ := json.Marshal(r)
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(rw, "%s", uj)
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello/", hellojson)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
