# ContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitHash** | Pointer to **string** | Commit hash related to the image. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Object creation time in RFC 3339 format. | [optional] 
**Id** | Pointer to **int64** | Database ID of the stored object. | [optional] 
**Image** | Pointer to **string** | Full container image including the registry, repository and tag. | [optional] 
**ImageName** | Pointer to **string** | Container image name without the registry, tag or sha256sum information. | [optional] 
**ImageRegistry** | Pointer to **string** | Container image registry without the image name, tag or sha256sum information. | [optional] 
**ImageTag** | Pointer to **string** | Container image tag without the registry, image name or sha256 information. | [optional] 
**Sha256sum** | Pointer to **string** | SHA256 sum of the container image. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Object last updated time in RFC 3339 format. | [optional] 
**Versions** | Pointer to [**[]Version**](Version.md) | Versions associated with this container image. | [optional] 

## Methods

### NewContainerImage

`func NewContainerImage() *ContainerImage`

NewContainerImage instantiates a new ContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageWithDefaults

`func NewContainerImageWithDefaults() *ContainerImage`

NewContainerImageWithDefaults instantiates a new ContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitHash

`func (o *ContainerImage) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ContainerImage) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ContainerImage) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ContainerImage) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContainerImage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContainerImage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContainerImage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContainerImage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ContainerImage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerImage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerImage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ContainerImage) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerImage) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerImage) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ContainerImage) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageName

`func (o *ContainerImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ContainerImage) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegistry

`func (o *ContainerImage) GetImageRegistry() string`

GetImageRegistry returns the ImageRegistry field if non-nil, zero value otherwise.

### GetImageRegistryOk

`func (o *ContainerImage) GetImageRegistryOk() (*string, bool)`

GetImageRegistryOk returns a tuple with the ImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistry

`func (o *ContainerImage) SetImageRegistry(v string)`

SetImageRegistry sets ImageRegistry field to given value.

### HasImageRegistry

`func (o *ContainerImage) HasImageRegistry() bool`

HasImageRegistry returns a boolean if a field has been set.

### GetImageTag

`func (o *ContainerImage) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ContainerImage) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ContainerImage) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *ContainerImage) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetSha256sum

`func (o *ContainerImage) GetSha256sum() string`

GetSha256sum returns the Sha256sum field if non-nil, zero value otherwise.

### GetSha256sumOk

`func (o *ContainerImage) GetSha256sumOk() (*string, bool)`

GetSha256sumOk returns a tuple with the Sha256sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256sum

`func (o *ContainerImage) SetSha256sum(v string)`

SetSha256sum sets Sha256sum field to given value.

### HasSha256sum

`func (o *ContainerImage) HasSha256sum() bool`

HasSha256sum returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ContainerImage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContainerImage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContainerImage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContainerImage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersions

`func (o *ContainerImage) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ContainerImage) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ContainerImage) SetVersions(v []Version)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ContainerImage) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *ContainerImage) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *ContainerImage) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


