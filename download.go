package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ep := "http://traffic.libsyn.com/startup/TheStartupChatep446.mp3"
	resp, err := http.Get(ep)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}
