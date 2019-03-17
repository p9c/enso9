package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/parallelcointeam/enso9/mod"
	"github.com/parallelcointeam/enso9/tools"
)

func Home(w http.ResponseWriter, r *http.Request) {
	jDB, err := tools.JDB.ReadAll("coins")

	if err != nil {
		fmt.Println("Error", err)
	}
	var bitnodes []mod.Coin
	for _, f := range jDB {
		var coinFound mod.Coin
		if err := json.Unmarshal([]byte(f), &coinFound); err != nil {
			fmt.Println("Error", err)
		}
		if coinFound.BitNode {
			bitnodes = append(bitnodes, coinFound)
		}
	}

	home := map[string]interface{}{
		"bitnodes": bitnodes,
		// "coins": m,
	}

	tmpl, _ := template.New("").ParseFiles("./tpl/amp/home.gohtml", "./tpl/amp/amp.gohtml", "./tpl/amp/inc/header.gohtml", "./tpl/amp/inc/footer.gohtml", "./tpl/amp/svgs.gohtml", "./tpl/amp/style.css")
	tmpl.ExecuteTemplate(w, "home", home)

}
