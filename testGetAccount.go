package main

import (
	"testing"
	"fmt"
)

func TestGetAccount(t *testing.T) {
	sib := getAPIClient(t)
	res, err := sib.Account.GetAccount(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

