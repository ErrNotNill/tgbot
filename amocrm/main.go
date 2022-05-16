package main

import (
	"fmt"
	"github.com/alexeykhan/amocrm"
)

func main() {
	amoCRM := amocrm.New("clientID", "clientSecret", "redirectURL")

	state := amocrm.RandomState()
	mode := amocrm.PostMessageMode

	authURL, err := amoCRM.AuthorizeURL(state, mode)
	if err != nil {
		fmt.Println("failed to get auth url:", err)
		return
	}

	fmt.Println("Redirect user to this URL:")
	fmt.Println(authURL)
}

