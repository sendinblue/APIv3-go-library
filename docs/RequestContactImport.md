# RequestContactImport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Mandatory if fileBody not defined. URL of the file to be imported (no local file). Possible file types: .txt, .csv | [optional] [default to null]
**FileBody** | **string** | Mandatory if fileUrl is not defined. CSV content to be imported. Use semicolon to separate multiple attributes | [optional] [default to null]
**ListIds** | **[]int32** | Manadatory if newList is not defined. Ids of the lists in which to add the contacts | [optional] [default to null]
**NotifyUrl** | **string** | URL that will be called once the export process is finished | [optional] [default to null]
**NewList** | [**RequestContactImportNewList**](RequestContactImportNewList.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


