# VersionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**ContainerImages** | Pointer to [**[]ContainerImageBody**](ContainerImageBody.md) | Container images associated with this version. | [optional] 
**Id** | Pointer to **string** | Version in semantic versioning format. | [optional] 

## Methods

### NewVersionBody

`func NewVersionBody() *VersionBody`

NewVersionBody instantiates a new VersionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionBodyWithDefaults

`func NewVersionBodyWithDefaults() *VersionBody`

NewVersionBodyWithDefaults instantiates a new VersionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *VersionBody) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *VersionBody) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *VersionBody) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *VersionBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetContainerImages

`func (o *VersionBody) GetContainerImages() []ContainerImageBody`

GetContainerImages returns the ContainerImages field if non-nil, zero value otherwise.

### GetContainerImagesOk

`func (o *VersionBody) GetContainerImagesOk() (*[]ContainerImageBody, bool)`

GetContainerImagesOk returns a tuple with the ContainerImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImages

`func (o *VersionBody) SetContainerImages(v []ContainerImageBody)`

SetContainerImages sets ContainerImages field to given value.

### HasContainerImages

`func (o *VersionBody) HasContainerImages() bool`

HasContainerImages returns a boolean if a field has been set.

### SetContainerImagesNil

`func (o *VersionBody) SetContainerImagesNil(b bool)`

 SetContainerImagesNil sets the value for ContainerImages to be an explicit nil

### UnsetContainerImages
`func (o *VersionBody) UnsetContainerImages()`

UnsetContainerImages ensures that no value is present for ContainerImages, not even an explicit nil
### GetId

`func (o *VersionBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VersionBody) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


