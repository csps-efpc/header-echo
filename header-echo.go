package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"regexp"
)

var matcher regexp.Regexp

func respond(w http.ResponseWriter, r *http.Request) {
	respondable := make(map[string]string)
	for key, v := range r.Header {
		if matcher.MatchString(key) {
			respondable[key] = v[0]
		}
	}
	var jsonResponse, err = json.MarshalIndent(respondable, "", " ")
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		w.Write(jsonResponse)
	}
}

func handleRequests(port int, pattern string) {

	matcherPtr := regexp.MustCompile(pattern)
	matcher = *matcherPtr
	http.HandleFunc("/echo", respond)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	portPtr := flag.Int("port", 10001, "Port on which to listen")
	patternPtr := flag.String("pattern", "^.*$", "Regex pattern to filter headers")
	flag.Parse()
	port := *portPtr
	pattern := *patternPtr
	fmt.Println("Listening on port:", port)
	handleRequests(port, pattern)

}
