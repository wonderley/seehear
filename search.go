// 0f9f7f388dcc4e069f817ca751907715
// podcast title and episode title search parameters
// strategy: first find the podcast, then search for the episode
// another strategy: search for all episodes with the text,
// then narrow as needed
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type podcast struct {
	ID string
}

type podcastResponse struct {
	Results []podcast
}

type episode struct {
	Audio          string
	AudioLengthSec int64
	ID             string
	TitleOriginal  string
}

type episodeResponse struct {
	Results []episode
}

func getBody(params map[string]string) []byte {
	url := "https://listen-api.listennotes.com/api/v2/search"
	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-ListenAPI-Key", "0f9f7f388dcc4e069f817ca751907715")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

// Returns the podcast
func searchForPodcast() podcast {
	params := make(map[string]string)
	params["q"] = "the indie hackers podcast"
	params["type"] = "podcast"
	body := getBody(params)
	var podcastResponse podcastResponse
	json.Unmarshal(body, &podcastResponse)
	if len(podcastResponse.Results) == 0 {
		panic("result count 0")
	}
	podcast := podcastResponse.Results[0]
	fmt.Print("podcast ID: ", podcast.ID, "\n")
	return podcast
}

func searchForEpisode(podcast podcast) episode {
	params := make(map[string]string)
	params["q"] = "becoming indistractable"
	params["type"] = "episode"
	params["ocid"] = podcast.ID
	body := getBody(params)
	var episodeResponse episodeResponse
	json.Unmarshal(body, &episodeResponse)
	if len(episodeResponse.Results) == 0 {
		panic("result count 0")
	}
	ep := episodeResponse.Results[0]
	fmt.Print("episode ID: ", ep.ID, "\n")
	return ep
}

func downloadEpisode(episode episode) {
	resp, err := http.Get(episode.Audio)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	out, err := os.Create("download.mp3")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(out, resp.Body)
}

func main() {
	podcast := searchForPodcast()
	episode := searchForEpisode(podcast)
	downloadEpisode(episode)
}
