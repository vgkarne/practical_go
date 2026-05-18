package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Given a github user login, retuen name and number of public repos

func main() {
	fmt.Println(UserInfo("ardanlabs"))
}
func demo() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad statys - %s\n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")

	fmt.Println("conrent-type", ctype)

	var reply struct {
		Name         string
		Public_Repos int
	}

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println(reply.Name, reply.Public_Repos)
}

func UserInfo(login string) (string, int, error) {
	url := "https://api.github.com/users" + login
	resp, err := http.Get("https://api.github.com/users/ardanlabs")
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad statys - %s\n", resp.Status)
		return "", 0, fmt.Errorf("%q -bad status:%s", url, resp.Status)
	}

	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name         string
		Public_Repos int
	}

	dec := json.NewDecoder(r)

	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR: ", err)
		return "", 0, err
	}
	return reply.Name, reply.Public_Repos, nil
}
