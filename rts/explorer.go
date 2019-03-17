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

func Explorer(w http.ResponseWriter, r *http.Request) {
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
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/explorer.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/inc/header.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/base.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "base", data)
}
func Block(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	blockid := vars["id"]
	var gCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &gCoin); err != nil {
		fmt.Println("Error", err)
	}

	data := map[string]interface{}{
		"coin":  gCoin,
		"block": blockid,
	}
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/block.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "base", data)

}
func Hash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	hash := vars["id"]

	var gCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &coin); err != nil {
		fmt.Println("Error", err)
	}

	data := map[string]interface{}{
		"coin": gCoin,
		"hash": hash,
	}
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/block.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "base", data)

}
func Tx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	txid := vars["id"]
	var gCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &gCoin); err != nil {
		fmt.Println("Error", err)
	}

	data := map[string]interface{}{
		"coin": gCoin,
		"txid": txid,
	}
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/tx.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/inc/header.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/base.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "base", data)

}

func DoSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	r.ParseForm()
	search := r.FormValue("chsrc")

	tps := []string{"block", "hash", "tx", "addr"}
	var tpt = "noresults"
	for _, tp := range tps {
		url := "http://" + cfg.Jorm + "/e/" + coin + "/" + tp + "/" + search
		fmt.Println("urlurlurlurlurlurlurlurlurlurlurl", url)
		resp, err := tools.GetData(url)
		var search map[string]interface{}
		json.Unmarshal(resp, &search)
		if err != nil {
			fmt.Println("Read error", err)
		}
		if search["d"] != nil {
			tpt = tp

		}

	}

	http.Redirect(w, r, fmt.Sprintf("/explorer/"+tpt+"/"+search), http.StatusPermanentRedirect)
}
