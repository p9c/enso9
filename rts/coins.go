package rts

import (
	"encoding/base64"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/enso9/mod"
	"github.com/parallelcointeam/enso9/tools"
)

func Coins(w http.ResponseWriter, r *http.Request) {
	jDB, err := tools.JDB.ReadAll("coins")
	if err != nil {
		fmt.Println("Error", err)
	}
	var scoins []mod.Coin
	for _, f := range jDB {
		var coinFound mod.Coin
		if err := json.Unmarshal([]byte(f), &coinFound); err != nil {
			fmt.Println("Error", err)
		}
		scoins = append(scoins, coinFound)
	}

	tmpl, _ := template.New("").ParseFiles("./tpl/amp/coins.gohtml", "./tpl/amp/amp.gohtml")
	tmpl.ExecuteTemplate(w, "coins", scoins)

}

func CoinPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	var gCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &gCoin); err != nil {
		fmt.Println("Error", err)
	}
	data := map[string]interface{}{
		"coin": gCoin,
	}
	tmpl, _ := template.New("").ParseFiles("./tpl/amp/coin.gohtml", "./tpl/amp/pnl/infotitle.gohtml", "./tpl/amp/pnl/specs.gohtml", "./tpl/amp/pnl/trends.gohtml", "./tpl/amp/pnl/infoprice.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/inc/header.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/base.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "base", data)

}
func ImgHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	size := vars["size"]
	var vCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &vCoin); err != nil {
		fmt.Println("Error", err)
	}
	var img string
	switch size {
	case "16":
		img = vCoin.Imgs.Img16
	case "32":
		img = vCoin.Imgs.Img32
	case "64":
		img = vCoin.Imgs.Img64
	case "128":
		img = vCoin.Imgs.Img128
	case "256":
		img = vCoin.Imgs.Img256
	}
	encoded, _ := base64.StdEncoding.DecodeString(img)
	w.Write(encoded)
}

func IcoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	var vCoin mod.Coin
	if err := tools.JDB.Read("coins", coin, &vCoin); err != nil {
		fmt.Println("Error", err)
	}
	img := vCoin.Imgs.Img16
	encoded, _ := base64.StdEncoding.DecodeString(img)
	w.Write(encoded)
}
