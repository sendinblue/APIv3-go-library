package test

import (
	"testing"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/sendinblue/APIv3-go-library/client/contacts"
	"github.com/sendinblue/APIv3-go-library/models"
)

func TestContacts(t *testing.T) {
	sib := getAPIClient(t)
	email := getTestEmailUnique()

	createParams := contacts.NewCreateContactParams()
	createParams.CreateContact = &models.CreateContact{
		Email: strfmt.Email(email),
		Attributes: map[string]interface{}{
			"foo": "bar",
		},
	}
	createRes, err := sib.Contacts.CreateContact(createParams, nil)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("createRes")
		fmt.Println(createRes)
		t.Fatal(err)
	}
	t.Log(spewConfig.Sdump(createRes.Payload))

	// getInfoParams := contacts.NewGetContactInfoParams()
	// getInfoParams.Email = email
	// getInfoRes, err := sib.Contacts.GetContactInfo(getInfoParams, nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(spewConfig.Sdump(getInfoRes.Payload))

	// updateParams := contacts.NewUpdateContactParams()
	// updateParams.Email = email
	// updateParams.UpdateContact = &models.UpdateContact{
	// 	Attributes: map[string]string{
	// 		"foo": "baz",
	// 	},
	// }
	// _, err = sib.Contacts.UpdateContact(updateParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}
