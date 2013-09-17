package cajerorosario

import (
    "appengine"
    "appengine/datastore"
    "appengine/user"
    "html/template"
    "net/http"
    "time"
)

type Cajero struct {
    Author  string
    Content string
    Date    time.Time
}

var guestbookTemplate = template.Must(template.ParseFiles("templates/base.html", 
	"templates/list.html"))

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/sign", sign)
}
 
func root(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    q := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
    greetings := make([]Cajero, 0, 10)
    if _, err := q.GetAll(c, &greetings); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
	if err := guestbookTemplate.Execute(w, greetings); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func sign(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    g := Cajero{
		Content: r.FormValue("content"),
		Date:    time.Now(),
    }
    if u := user.Current(c); u != nil {
        g.Author = u.String()
    }
    _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Greeting", nil), &g)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/", http.StatusFound)
}
