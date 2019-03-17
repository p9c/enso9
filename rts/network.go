package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/enso9/cfg"
	"github.com/parallelcointeam/enso9/mod"
	"github.com/parallelcointeam/enso9/tools"
)

func Network(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]

	var gCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &gCoin); err != nil {
		fmt.Println("Error", err)
	}
	url := "http://" + cfg.Jorm + "/e/" + coin + "/b"

	b, _ := tools.GetData(url)
	var m map[string]int
	json.Unmarshal(b, &m)
	data := map[string]interface{}{
		"coin": gCoin,
		"last": m,
	}

	tmpl, _ := template.New("").ParseFiles("./tpl/amp/network.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/inc/header.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/base.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "base", data)
}
