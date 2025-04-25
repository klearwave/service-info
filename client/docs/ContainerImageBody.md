# ContainerImageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**CommitHash** | Pointer to **string** | Commit hash related to the image. | [optional] 
**Image** | Pointer to **string** | Full container image including the registry, repository and tag. | [optional] 
**Sha256sum** | Pointer to **string** | SHA256 sum of the container image. | [optional] 

## Methods

### NewContainerImageBody

`func NewContainerImageBody() *ContainerImageBody`

NewContainerImageBody instantiates a new ContainerImageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageBodyWithDefaults

`func NewContainerImageBodyWithDefaults() *ContainerImageBody`

NewContainerImageBodyWithDefaults instantiates a new ContainerImageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *ContainerImageBody) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ContainerImageBody) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ContainerImageBody) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ContainerImageBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCommitHash

`func (o *ContainerImageBody) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ContainerImageBody) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ContainerImageBody) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ContainerImageBody) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetImage

`func (o *ContainerImageBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerImageBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerImageBody) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ContainerImageBody) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSha256sum

`func (o *ContainerImageBody) GetSha256sum() string`

GetSha256sum returns the Sha256sum field if non-nil, zero value otherwise.

### GetSha256sumOk

`func (o *ContainerImageBody) GetSha256sumOk() (*string, bool)`

GetSha256sumOk returns a tuple with the Sha256sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256sum

`func (o *ContainerImageBody) SetSha256sum(v string)`

SetSha256sum sets Sha256sum field to given value.

### HasSha256sum

`func (o *ContainerImageBody) HasSha256sum() bool`

HasSha256sum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


