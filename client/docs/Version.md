# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildVersion** | Pointer to **string** | The build version and metadata of this release (e.g. prerelease.1; v0.1.2-prerelease.1 &#x3D;&#x3D; x.y.z-build.metadata). | [optional] 
**ContainerImages** | Pointer to [**[]ContainerImage**](ContainerImage.md) | Container images associated with this version. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Object creation time in RFC 3339 format. | [optional] 
**Id** | Pointer to **string** | Version in semantic versioning format. | [optional] 
**Stable** | Pointer to **bool** | Whether this is a stable version. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Object last updated time in RFC 3339 format. | [optional] 
**XVersion** | Pointer to **int64** | The X version of this release (e.g. 0; v0.1.2-prerelease.1 &#x3D;&#x3D; x.y.z-build.metadata). | [optional] 
**YVersion** | Pointer to **int64** | The Y version of this release (e.g. 1; v0.1.2-prerelease.1 &#x3D;&#x3D; x.y.z-build.metadata). | [optional] 
**ZVersion** | Pointer to **int64** | The Z version of this release (e.g. 2; v0.1.2-prerelease.1 &#x3D;&#x3D; x.y.z-build.metadata). | [optional] 

## Methods

### NewVersion

`func NewVersion() *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildVersion

`func (o *Version) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *Version) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *Version) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *Version) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetContainerImages

`func (o *Version) GetContainerImages() []ContainerImage`

GetContainerImages returns the ContainerImages field if non-nil, zero value otherwise.

### GetContainerImagesOk

`func (o *Version) GetContainerImagesOk() (*[]ContainerImage, bool)`

GetContainerImagesOk returns a tuple with the ContainerImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImages

`func (o *Version) SetContainerImages(v []ContainerImage)`

SetContainerImages sets ContainerImages field to given value.

### HasContainerImages

`func (o *Version) HasContainerImages() bool`

HasContainerImages returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Version) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Version) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Version) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Version) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Version) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Version) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Version) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Version) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStable

`func (o *Version) GetStable() bool`

GetStable returns the Stable field if non-nil, zero value otherwise.

### GetStableOk

`func (o *Version) GetStableOk() (*bool, bool)`

GetStableOk returns a tuple with the Stable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStable

`func (o *Version) SetStable(v bool)`

SetStable sets Stable field to given value.

### HasStable

`func (o *Version) HasStable() bool`

HasStable returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Version) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Version) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Version) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Version) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetXVersion

`func (o *Version) GetXVersion() int64`

GetXVersion returns the XVersion field if non-nil, zero value otherwise.

### GetXVersionOk

`func (o *Version) GetXVersionOk() (*int64, bool)`

GetXVersionOk returns a tuple with the XVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXVersion

`func (o *Version) SetXVersion(v int64)`

SetXVersion sets XVersion field to given value.

### HasXVersion

`func (o *Version) HasXVersion() bool`

HasXVersion returns a boolean if a field has been set.

### GetYVersion

`func (o *Version) GetYVersion() int64`

GetYVersion returns the YVersion field if non-nil, zero value otherwise.

### GetYVersionOk

`func (o *Version) GetYVersionOk() (*int64, bool)`

GetYVersionOk returns a tuple with the YVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYVersion

`func (o *Version) SetYVersion(v int64)`

SetYVersion sets YVersion field to given value.

### HasYVersion

`func (o *Version) HasYVersion() bool`

HasYVersion returns a boolean if a field has been set.

### GetZVersion

`func (o *Version) GetZVersion() int64`

GetZVersion returns the ZVersion field if non-nil, zero value otherwise.

### GetZVersionOk

`func (o *Version) GetZVersionOk() (*int64, bool)`

GetZVersionOk returns a tuple with the ZVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZVersion

`func (o *Version) SetZVersion(v int64)`

SetZVersion sets ZVersion field to given value.

### HasZVersion

`func (o *Version) HasZVersion() bool`

HasZVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


