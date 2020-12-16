package test

import (
	"fmt"
	"testing"
)

func TestGetAccount(t *testing.T) {
	sib := getAPIClient(t)
	res, err := sib.Account.GetAccount(nil, nil)
	if err != nil {
		fmt.Println("*** Error ***")
		fmt.Println(err)
		t.Fatal(err)
	}
	t.Log(spewConfig.Sdump(res.Payload))
}
