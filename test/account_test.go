package test

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	sib := getAPIClient(t)
	res, err := sib.Account.GetAccount(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spewConfig.Sdump(res.Payload))
}
