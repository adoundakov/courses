package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Reply struct {
	Name         string
	Public_Repos int
	Followers    int
	Github_Id    int `json:"id"`
	// What if we wanted to rewrite a key to be something else?
	// Add a field tag, there are lots that are supported, but for JSON
}

// Given a login, find the name and number of repos
func githubInfo(usename string) (string, int, error) {
	resp, err := http.Get("https://api.github.com/users/" + url.PathEscape(usename))
	if err != nil {
		return "", 0, fmt.Errorf("error: couldn't fetch - %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("error: not ok - %s", resp.Status)
	}

	var r Reply
	dec := json.NewDecoder(resp.Body)

	// Need to pass pointer to r, not r itself so you can actually write
	// data to it, otherwise Decode gets a copy of r
	if err := dec.Decode(&r); err != nil {
		return "", 0, fmt.Errorf("error: couldn't decode - %s", err)
	}

	return r.Name, r.Public_Repos, nil
}

func main() {
	// resp, err := http.Get("https://api.github.com/users/adoundakov")
	// // NOTE: There's a shortcut for this using `iferr`
	// if err != nil {
	// 	log.Fatalf("error: %s", err)
	// }

	// if resp.StatusCode != http.StatusOK {
	// 	log.Fatalf("error: %s", resp.Status)
	// }

	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// This is a useful pattern, called conditional w/ initialization
	// when we don't care about the return value from io.Copy, but do
	// want to know if there was an error.
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: can't copy - %s", err)
	// }

	// Reading the body is fine and all but we can't do anything yet
	// Need to marshall / unmarshall into Go (serialize / deserialize)
	// However, need to tell Go a bit about how to unmarshall the JSON
	// eg: number -> int8, int16, etc..
	// eg: array -> []T (a slice of type T, or any)
	// eg: object -> map[string]any, struct

	// How do we get from JSON to and from Go?
	// JSON -> io.Reader -> Go: use json.Decoder
	// JSON -> []byte (in mem) -> Go: use json.Unmarshal
	// Go -> io.Writer -> JSON: use json.Encoder
	// Go -> []byte -> JSON: use json.Marshal

	// So let's define a type (above) to process the response
	// var r Reply
	// dec := json.NewDecoder(resp.Body)

	// Need to pass pointer to r, not r itself so you can actually write
	// data to it, otherwise Decode gets a copy of r
	// if err := dec.Decode(&r); err != nil {
	// 	log.Fatalf("error: can't decode - %s", err)
	// }

	// Can use this to log
	// fmt.Println(r)
	// But probably better to use this so you get type info
	// fmt.Printf("%#v\n", r)

	name, repos, err := githubInfo("adoundakov")
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("Name: %s | Repos: %v", name, repos)
}
