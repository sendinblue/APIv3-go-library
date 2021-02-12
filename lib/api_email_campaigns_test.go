package lib

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
)

func TestCreateEmailCampaign(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)
	s := &CreateEmailCampaignSender{
		Name:  "raunak_jain",
		Email: "shubham.upadhyay@sendinblue.com",
	}
	params := CreateEmailCampaign{
		Sender:      s,
		Name:        "TestCampaign",
		HtmlContent: "Hi There! welcome to SIB",
		Subject:     "Welcome!",
	}
	result, resp, err := sib.EmailCampaignsApi.CreateEmailCampaign(ctx, params)
	if err != nil {
		fmt.Println("===in CreateEmailCampaign error===")
		t.Fatal(err)
	}
	fmt.Println("====CreateEmailCampaign Result:", result, "     ====    resp:", resp)
}

func TestGetEmailCampaigns(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)

	params := &EmailCampaignsApiGetEmailCampaignsOpts{
		Limit: optional.NewInt64(10),
	}
	result, resp, err := sib.EmailCampaignsApi.GetEmailCampaigns(ctx, params)
	if err != nil {
		fmt.Println("===in Get All EmailCampaign error===")
		t.Fatal(err)
	}
	fmt.Println("====GetEmailCampaigns Result:", result, "     ====    resp:", resp)
}
