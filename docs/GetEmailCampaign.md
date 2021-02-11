# GetEmailCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestSent** | **bool** | Retrieved the status of test email sending. (true&#x3D;Test email has been sent  false&#x3D;Test email has not been sent) | [default to null]
**Header** | **string** | Header of the campaign | [default to null]
**Footer** | **string** | Footer of the campaign | [default to null]
**Sender** | [***GetExtendedCampaignOverviewSender**](getExtendedCampaignOverview_sender.md) |  | [default to null]
**ReplyTo** | **string** | Email defined as the \&quot;Reply to\&quot; of the campaign | [default to null]
**ToField** | **string** | Customisation of the \&quot;to\&quot; field of the campaign | [optional] [default to null]
**HtmlContent** | **string** | HTML content of the campaign | [default to null]
**ShareLink** | **string** | Link to share the campaign on social medias | [optional] [default to null]
**Tag** | **string** | Tag of the campaign | [optional] [default to null]
**CreatedAt** | **string** | Creation UTC date-time of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**ModifiedAt** | **string** | UTC date-time of last modification of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**InlineImageActivation** | **bool** | Status of inline image. inlineImageActivation &#x3D; false means image can’t be embedded, &amp; inlineImageActivation &#x3D; true means image can be embedded, in the email. | [optional] [default to null]
**MirrorActive** | **bool** | Status of mirror links in campaign. mirrorActive &#x3D; false means mirror links are deactivated, &amp; mirrorActive &#x3D; true means mirror links are activated, in the campaign | [optional] [default to null]
**Recurring** | **bool** | FOR TRIGGER ONLY ! Type of trigger campaign.recurring &#x3D; false means contact can receive the same Trigger campaign only once, &amp; recurring &#x3D; true means contact can receive the same Trigger campaign several times | [optional] [default to null]
**SentDate** | **string** | Sent UTC date-time of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ). Only available if &#x27;status&#x27; of the campaign is &#x27;sent&#x27; | [optional] [default to null]
**ReturnBounce** | **int64** | Total number of non-delivered campaigns for a particular campaign id. | [optional] [default to null]
**Recipients** | [***interface{}**](interface{}.md) |  | [default to null]
**Statistics** | [***interface{}**](interface{}.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

