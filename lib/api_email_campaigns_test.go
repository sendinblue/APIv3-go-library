package lib

import (
	"context"
	"fmt"
	"testing"
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
		Sender:                s,
		Name:                  "TestCampaign",
		HtmlContent:           "Hi There! welcome to SIB",
		Subject:               "Welcome!",
		AttachmentUrl:         "https://attachment.domain.com/myAttachmentFromUrl.jpg",
		InlineImageActivation: false,
		ReplyTo:               "replyto@domain.com",
		TemplateId:            int64(10),
		ToField:               "John Doe",
		Recipients: &CreateEmailCampaignRecipients{
			ListIds:          []int64{2},
			ExclusionListIds: []int64{1, 3},
		},
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

	body := UploadImageToGallery{
		ImageUrl: "https://example.net/upload-file.jpg",
		Name:     "Example Image",
	}
	resp, err := sib.EmailCampaignsApi.UploadImageToGallery(ctx, body)
	if err != nil {
		fmt.Println("===in Get All EmailCampaign error===")
		t.Fatal(err)
	}
	fmt.Println("====GetEmailCampaigns Result:", "     ====    resp:", resp)
}
