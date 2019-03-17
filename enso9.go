package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/enso9/rts"
)

const (
	// HTTPMethodOverrideHeader is a commonly used
	// http header to override a request method.
	// HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
	// HTTPMethodOverrideFormKey is a commonly used
	// HTML form key to override a request method.
	// HTTPMethodOverrideFormKey = "_method"
	// vhost                     = "com-http.us"

	vhost = "com-http.us"
)

var r = mux.NewRouter()

func main() {
	// port := "443"
	// r.PathPrefix("/s/").Handler(http.StripPrefix("/s/", http.FileServer(http.Dir("./tpl/static/"))))
	// r.PathPrefix("/amp/").Handler(http.StripPrefix("/amp/", http.FileServer(http.Dir("./tpl/amp/"))))

	r.Path("/").HandlerFunc(rts.Home).Name("index")

	r.Path("/coins/").HandlerFunc(rts.Coins).Methods("GET")
	r.Path("/bitnodes/").HandlerFunc(rts.BitNodes).Methods("GET")

	r.Path("/coins/{coin}").HandlerFunc(rts.CoinPage).Methods("GET")
	r.Path("/coins/{coin}/favicon.ico").HandlerFunc(rts.IcoHandler).Name("ico")

	r.Path("/coins/{coin}/network").HandlerFunc(rts.Network).Methods("GET")
	r.Path("/coins/{coin}/community").HandlerFunc(rts.Community).Methods("GET")
	r.Path("/coins/{coin}/art").HandlerFunc(rts.Art).Methods("GET")

	r.Path("/coins/{coin}/explorer").HandlerFunc(rts.Explorer).Methods("GET")
	r.Path("/coins/{coin}/explorer/block/{id}").HandlerFunc(rts.Block).Name("block")
	r.Path("/coins/{coin}/explorer/hash/{id}").HandlerFunc(rts.Hash).Name("hash")
	r.Path("/coins/{coin}/explorer/tx/{id}").HandlerFunc(rts.Tx).Name("tx")
	// r.Path("/{coin}/explorer/address/{id}").HandlerFunc(rts.Address).Name("address")
	r.Path("/coins/{coin}/explorer/search").HandlerFunc(rts.DoSearch).Name("search")

	// r.PathPrefix("/i/").Handler(http.StripPrefix("/i/", http.FileServer(http.Dir("/COMHTTP/imgs"))))

	// r.HandleFunc("/", rts.AmpFrontHandler) // GET
	// r.HandleFunc("/coins/{slug}/f/cmc", rts.CMCHandler).Name("cmc")

	r.Path("/coins/{coin}/img/{size}").HandlerFunc(rts.ImgHandler).Name("img")
	r.Path("/coins/{coin}/frame/{frame}/{file}").HandlerFunc(rts.Frames).Name("frames")
	// r.Path("/lib").PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/marcetin/go/src/github.com/parallelcointeam/enso9/static/libs/"))))
	// r.Path("/libs/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/libs/"))))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// r.Schemes("https")

	// go log.Fatal(http.ListenAndServe(":80", handlers.CORS()(r)))

	// srv := &http.Server{
	// 	Handler: handlers.CORS()(handlers.CompressHandler(r)),
	// 	Addr:    ":443",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Fatal(srv.ListenAndServeTLS("./cfg/comhttp.crt", "./cfg/comhttp.key"))

	fmt.Println("Listening on port:", 9888)
	log.Fatal(http.ListenAndServe(":9888", handlers.CORS()(r)))

}
