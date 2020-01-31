package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	tmpl := template.Must(template.New("index").ParseGlob("delivery/http/handler/*.html"))
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
	    fmt.Println(err)
	}
	return
}
