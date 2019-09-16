// podcast title and episode title search parameters
// strategy: first find the podcast, then search for the episode
// another strategy: search for all episodes with the text,
// then narrow as needed
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
