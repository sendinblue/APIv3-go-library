# CreateAttribute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Value of the attribute. Use only if the attribute&#x27;s category is &#x27;calculated&#x27; or &#x27;global&#x27; | [optional] [default to null]
**Enumeration** | [**[]CreateAttributeEnumeration**](CreateAttributeEnumeration.md) | List of values and labels that the attribute can take. Use only if the attribute&#x27;s category is \&quot;category\&quot;. For example, [{\&quot;value\&quot;:1, \&quot;label\&quot;:\&quot;male\&quot;}, {\&quot;value\&quot;:2, \&quot;label\&quot;:\&quot;female\&quot;}] | [optional] [default to null]
**Type_** | **string** | Type of the attribute. Use only if the attribute&#x27;s category is &#x27;normal&#x27;, &#x27;category&#x27; or &#x27;transactional&#x27; ( type &#x27;boolean&#x27; is only available if the category is &#x27;normal&#x27; attribute, type &#x27;id&#x27; is only available if the category is &#x27;transactional&#x27; attribute &amp; type &#x27;category&#x27; is only available if the category is &#x27;category&#x27; attribute ) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

