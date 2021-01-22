package newtest

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateEmailCampaign(t *testing.T) {
	var ctx context.Context
	var aa = APIClient{
		cfg: NewConfiguration(),
	}
	//aa.cfg = NewConfiguration()
	fmt.Println(aa)
	sib := NewAPIClient(aa.cfg)
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
