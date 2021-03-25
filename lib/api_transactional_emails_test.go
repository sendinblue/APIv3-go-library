package lib

import (
	"context"
	"fmt"
	"testing"
)

func TestTransactionalEmailsApi(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)
	body := SendSmtpEmail{
		HtmlContent: "<html><body><h1>This is my first transactional email </h1></body></html",
		Subject:     "New Subject",
		TemplateId:  int64(5),
	}

	model, resp, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, body)
	if err != nil {
		fmt.Println("===in CreateEmailCampaign error===")
		t.Fatal(err)
	}
	fmt.Println("====CreateEmailCampaign Result:", overview, "     ====    resp:", resp)
}
