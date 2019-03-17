package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/parallelcointeam/enso9/mod"
	"github.com/parallelcointeam/enso9/tools"
)

func BitNodes(w http.ResponseWriter, r *http.Request) {
	jDB, err := tools.JDB.ReadAll("nodes")

	if err != nil {
		fmt.Println("Error", err)
	}
	var bitnodes []mod.Node
	for _, f := range jDB {
		var nodesFound mod.Node
		if err := json.Unmarshal([]byte(f), &nodesFound); err != nil {
			fmt.Println("Error", err)
		}
		if nodesFound.BitNode {

			var gCoin mod.Coin
			if err := tools.JDB.Read("coins", nodesFound.Coin, &gCoin); err != nil {
				fmt.Println("Error", err)
			}

			nodesFound.CoinData = gCoin
			bitnodes = append(bitnodes, nodesFound)
		}
	}

	tmpl, _ := template.New("").ParseFiles("./tpl/amp/bitnodes.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/header.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "bitnodes", bitnodes)

}
