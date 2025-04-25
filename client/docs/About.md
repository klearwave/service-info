# About

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**CommitHash** | Pointer to **string** | Commit hash of this running version. | [optional] 
**Version** | Pointer to **string** | Running version of this service in semantic versioning format. | [optional] 

## Methods

### NewAbout

`func NewAbout() *About`

NewAbout instantiates a new About object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutWithDefaults

`func NewAboutWithDefaults() *About`

NewAboutWithDefaults instantiates a new About object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *About) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *About) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *About) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *About) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCommitHash

`func (o *About) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *About) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *About) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *About) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetVersion

`func (o *About) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *About) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *About) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *About) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


