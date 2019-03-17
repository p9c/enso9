package rts

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/enso9/mod"
	"github.com/parallelcointeam/enso9/tools"
)

func CMCHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	var gCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &gCoin); err != nil {
		fmt.Println("Error", err)
	}
	sym := gCoin.Symbol
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/transaction.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "html", sym)
}

func Frames(w http.ResponseWriter, r *http.Request) {
	// name := mux.Vars(r)["frame"] + "S" + mux.Vars(r)["file"]
	vars := mux.Vars(r)
	coin := vars["coin"]
	frame := vars["frame"]
	file := vars["file"]
	var gCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &gCoin); err != nil {
		fmt.Println("Error", err)
	}
	data := map[string]interface{}{
		"coin": gCoin,
	}
	tmpl, _ := template.New("").ParseFiles("./tpl/frames/" + frame + "/" + file + ".gohtml")
	tmpl.ExecuteTemplate(w, file, data)
}
