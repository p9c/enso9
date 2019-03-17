package tools

import (
	"io/ioutil"
	"net/http"

	scribble "github.com/nanobox-io/golang-scribble"
)

var JDB, _ = scribble.New("/web/comhttp/jdb/", nil)

func GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)
	return mapBody, err
}
