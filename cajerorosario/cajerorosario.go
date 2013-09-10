package cajerorosario

import (
	"appengine"
    "appengine/user"
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", cajerorosario)
}

func cajerorosario(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }
    fmt.Fprintf(w, "Bienvenido, %v!", u)
}
