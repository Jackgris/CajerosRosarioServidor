package cajerorosario

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", cajerorosario)
}

func cajerorosario(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Test App Engine</h1>")
}
