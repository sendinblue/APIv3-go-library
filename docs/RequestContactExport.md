# RequestContactExport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportAttributes** | **[]string** | List of all the attributes that you want to export. These attributes must be present in your contact database. For example, [&#x27;fname&#x27;, &#x27;lname&#x27;, &#x27;email&#x27;]. | [optional] [default to null]
**ContactFilter** | [***interface{}**](interface{}.md) | This attribute has been deprecated and will be removed by January 1st, 2021. Only one of the two filter options (contactFilter or customContactFilter) can be passed in the request. Set the filter for the contacts to be exported. For example, {\&quot;blacklisted\&quot;:true} will export all the blacklisted contacts.  | [optional] [default to null]
**CustomContactFilter** | [***RequestContactExportCustomContactFilter**](RequestContactExportCustomContactFilter.md) |  | [optional] [default to null]
**NotifyUrl** | **string** | Webhook that will be called once the export process is finished. For reference, https://help.sendinblue.com/hc/en-us/articles/360007666479 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

