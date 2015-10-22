package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type NameRequest struct {
    Name string `json:"name"`
  }
type Response struct {
    Greeting string `json:"greeting"`
}
func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "name: , %s!\n", p.ByName("name"))
}

func hi(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
    var a NameRequest
    var b Response
    decoder := json.NewDecoder(req.Body)
    fmt.Println(req.Body)

    err := decoder.Decode(&a)
    if err != nil {
    panic(err)
    }

    b.Greeting="Hello,"+a.Name
    show,_ := json.Marshal(r)
    fmt.Fprintf(rw,string(show))
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hi", hi)
    server := http.Server{
            Addr:        "localhost:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}