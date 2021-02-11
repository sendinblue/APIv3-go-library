package lib

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
)

func TestCreateContact(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)
	email := getTestEmailUnique()

	var params = CreateContact{
		Email: email,
	}

	model, response, err := sib.ContactsApi.CreateContact(ctx, params)
	if err != nil {
		fmt.Println("===in get createContact error===")
		t.Fatal(err)
	}
	fmt.Println("====Contact Model:", model, "    ====   Response:", response, "   ====error====   ", err)
}

func TestUpdateContact(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)
	email := "test+4600995338790023574@sendinblue.com"

	var params = UpdateContact{
		ListIds: []int64{10},
	}

	response, err := sib.ContactsApi.UpdateContact(ctx, params, email)
	if err != nil {
		fmt.Println("===in get Update Contact error===")
		t.Fatal(err)
	}
	fmt.Println("====Contact Update Model====   Response:", response, "   ====error====   ", err)
}

func TestGetContact(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)

	//var params = *ContactsApiGetContactsOpts{}

	response, res, err := sib.ContactsApi.GetContacts(ctx, nil)
	if err != nil {
		fmt.Println("===in get Get Contact error===")
		t.Fatal(err)
	}
	fmt.Println("====Contact Get Model====   Response:", response, "   ====error====   ", err, res)
}

const (
	testEmail       = "test@sendinblue.com"
	testEmailUnique = "test+%d@sendinblue.com"
)

func getTestEmailUnique() string {
	return fmt.Sprintf(testEmailUnique, rand.Int63())
}
