package rts

import (
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
)

func Market(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["slug"]
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/explorer.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/inc/header.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/base.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "base", coin)
}
