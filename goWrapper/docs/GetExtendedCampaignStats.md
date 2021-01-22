# GetExtendedCampaignStats

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalStats** | [***AllOfgetExtendedCampaignStatsGlobalStats**](AllOfgetExtendedCampaignStatsGlobalStats.md) | Overall statistics of the campaign | [default to null]
**CampaignStats** | [**[]AllOfgetExtendedCampaignStatsCampaignStatsItems**](interface{}.md) | List-wise statistics of the campaign. | [default to null]
**MirrorClick** | **int64** | Number of clicks on mirror link | [default to null]
**Remaining** | **int64** | Number of remaning emails to send | [default to null]
**LinksStats** | [***interface{}**](interface{}.md) | Statistics about the number of clicks for the links | [default to null]
**StatsByDomain** | [***map[string]GetCampaignStats**](map.md) |  | [default to null]
**StatsByDevice** | [***GetStatsByDevice**](getStatsByDevice.md) |  | [default to null]
**StatsByBrowser** | [***map[string]GetDeviceBrowserStats**](map.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

