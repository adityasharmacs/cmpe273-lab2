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
type NameResponse struct {
    Greeting string `json:"greeting"`
}
func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "name: , %s!\n", p.ByName("name"))
}

func hi(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
    var name NameRequest
    var resp NameResponse
    decoder := json.NewDecoder(req.Body)
    fmt.Println(req.Body)

    err1 := decoder.Decode(&name)
    if err1 != nil {
    panic(err1)
    }

    resp.Greeting="Hello,"+name.Name

    output,_ := json.Marshal(resp)
    fmt.Fprintf(rw,string(output))
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", hi)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
