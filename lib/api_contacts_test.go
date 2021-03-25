package lib

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/antihax/optional"
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
	email := "example@example.com" //string | Email (urlencoded) OR ID of the contact
	attr := map[string]interface{}{
		"EMAIL":     "example2@example2.com",
		"FIRSTNAME": "John Doe",
	}
	var params = UpdateContact{
		ListIds:    []int64{10},
		Attributes: attr,
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

	var params = &ContactsApiGetContactsOpts{
		Sort: optional.NewString("asc"),
	}

	response, res, err := sib.ContactsApi.GetContacts(ctx, params)
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

func TestGetContactInfo(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)
	identifier := "example@example.com" //string|email identifier
	params := &ContactsApiGetContactStatsOpts{
		StartDate: optional.NewString("2013-10-20"),
		EndDate:   optional.NewString("2020-09-20"),
	}

	model, resp, err := sib.ContactsApi.GetContactStats(ctx, identifier, params)
	if err != nil {
		fmt.Println("===in Get Contact Detail error===")
		t.Fatal(err)
	}
	fmt.Println("====Contact Get Model====   Response:", resp, "   ====error====   ", err, model)
}

func TestCreateAttribute(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)
	body := CreateDoiContact{
		Email:          "example@example.com",
		IncludeListIds: []int64{2},
		TemplateId:     int64(1),
		RedirectionUrl: "https://redirectionurl@yourdomain.com",
	}

	resp, err := sib.ContactsApi.CreateDoiContact(ctx, body)
	if err != nil {
		fmt.Println("===in Get Contact Detail error===")
		t.Fatal(err)
	}
	fmt.Println("====Contact Get Model====   Response:", resp, "   ====error====   ", err)
}
