// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the Version type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Version{}

// Version struct for Version
type Version struct {
	// The build version and metadata of this release (e.g. prerelease.1; v0.1.2-prerelease.1 == x.y.z-build.metadata).
	BuildVersion *string `json:"build_version,omitempty"`
	// Container images associated with this version.
	ContainerImages []ContainerImage `json:"container_images,omitempty"`
	// Object creation time in RFC 3339 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Version in semantic versioning format.
	Id *string `json:"id,omitempty"`
	// Whether this is a stable version.
	Stable *bool `json:"stable,omitempty"`
	// Object last updated time in RFC 3339 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The X version of this release (e.g. 0; v0.1.2-prerelease.1 == x.y.z-build.metadata).
	XVersion *int64 `json:"x_version,omitempty"`
	// The Y version of this release (e.g. 1; v0.1.2-prerelease.1 == x.y.z-build.metadata).
	YVersion *int64 `json:"y_version,omitempty"`
	// The Z version of this release (e.g. 2; v0.1.2-prerelease.1 == x.y.z-build.metadata).
	ZVersion *int64 `json:"z_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Version Version

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion() *Version {
	this := Version{}
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetBuildVersion returns the BuildVersion field value if set, zero value otherwise.
func (o *Version) GetBuildVersion() string {
	if o == nil || IsNil(o.BuildVersion) {
		var ret string
		return ret
	}
	return *o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetBuildVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BuildVersion) {
		return nil, false
	}
	return o.BuildVersion, true
}

// HasBuildVersion returns a boolean if a field has been set.
func (o *Version) HasBuildVersion() bool {
	if o != nil && !IsNil(o.BuildVersion) {
		return true
	}

	return false
}

// SetBuildVersion gets a reference to the given string and assigns it to the BuildVersion field.
func (o *Version) SetBuildVersion(v string) {
	o.BuildVersion = &v
}

// GetContainerImages returns the ContainerImages field value if set, zero value otherwise.
func (o *Version) GetContainerImages() []ContainerImage {
	if o == nil || IsNil(o.ContainerImages) {
		var ret []ContainerImage
		return ret
	}
	return o.ContainerImages
}

// GetContainerImagesOk returns a tuple with the ContainerImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetContainerImagesOk() ([]ContainerImage, bool) {
	if o == nil || IsNil(o.ContainerImages) {
		return nil, false
	}
	return o.ContainerImages, true
}

// HasContainerImages returns a boolean if a field has been set.
func (o *Version) HasContainerImages() bool {
	if o != nil && !IsNil(o.ContainerImages) {
		return true
	}

	return false
}

// SetContainerImages gets a reference to the given []ContainerImage and assigns it to the ContainerImages field.
func (o *Version) SetContainerImages(v []ContainerImage) {
	o.ContainerImages = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Version) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Version) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Version) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Version) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Version) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Version) SetId(v string) {
	o.Id = &v
}

// GetStable returns the Stable field value if set, zero value otherwise.
func (o *Version) GetStable() bool {
	if o == nil || IsNil(o.Stable) {
		var ret bool
		return ret
	}
	return *o.Stable
}

// GetStableOk returns a tuple with the Stable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetStableOk() (*bool, bool) {
	if o == nil || IsNil(o.Stable) {
		return nil, false
	}
	return o.Stable, true
}

// HasStable returns a boolean if a field has been set.
func (o *Version) HasStable() bool {
	if o != nil && !IsNil(o.Stable) {
		return true
	}

	return false
}

// SetStable gets a reference to the given bool and assigns it to the Stable field.
func (o *Version) SetStable(v bool) {
	o.Stable = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Version) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Version) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Version) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetXVersion returns the XVersion field value if set, zero value otherwise.
func (o *Version) GetXVersion() int64 {
	if o == nil || IsNil(o.XVersion) {
		var ret int64
		return ret
	}
	return *o.XVersion
}

// GetXVersionOk returns a tuple with the XVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetXVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.XVersion) {
		return nil, false
	}
	return o.XVersion, true
}

// HasXVersion returns a boolean if a field has been set.
func (o *Version) HasXVersion() bool {
	if o != nil && !IsNil(o.XVersion) {
		return true
	}

	return false
}

// SetXVersion gets a reference to the given int64 and assigns it to the XVersion field.
func (o *Version) SetXVersion(v int64) {
	o.XVersion = &v
}

// GetYVersion returns the YVersion field value if set, zero value otherwise.
func (o *Version) GetYVersion() int64 {
	if o == nil || IsNil(o.YVersion) {
		var ret int64
		return ret
	}
	return *o.YVersion
}

// GetYVersionOk returns a tuple with the YVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetYVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.YVersion) {
		return nil, false
	}
	return o.YVersion, true
}

// HasYVersion returns a boolean if a field has been set.
func (o *Version) HasYVersion() bool {
	if o != nil && !IsNil(o.YVersion) {
		return true
	}

	return false
}

// SetYVersion gets a reference to the given int64 and assigns it to the YVersion field.
func (o *Version) SetYVersion(v int64) {
	o.YVersion = &v
}

// GetZVersion returns the ZVersion field value if set, zero value otherwise.
func (o *Version) GetZVersion() int64 {
	if o == nil || IsNil(o.ZVersion) {
		var ret int64
		return ret
	}
	return *o.ZVersion
}

// GetZVersionOk returns a tuple with the ZVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetZVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.ZVersion) {
		return nil, false
	}
	return o.ZVersion, true
}

// HasZVersion returns a boolean if a field has been set.
func (o *Version) HasZVersion() bool {
	if o != nil && !IsNil(o.ZVersion) {
		return true
	}

	return false
}

// SetZVersion gets a reference to the given int64 and assigns it to the ZVersion field.
func (o *Version) SetZVersion(v int64) {
	o.ZVersion = &v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Version) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildVersion) {
		toSerialize["build_version"] = o.BuildVersion
	}
	if !IsNil(o.ContainerImages) {
		toSerialize["container_images"] = o.ContainerImages
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Stable) {
		toSerialize["stable"] = o.Stable
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.XVersion) {
		toSerialize["x_version"] = o.XVersion
	}
	if !IsNil(o.YVersion) {
		toSerialize["y_version"] = o.YVersion
	}
	if !IsNil(o.ZVersion) {
		toSerialize["z_version"] = o.ZVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Version) UnmarshalJSON(data []byte) (err error) {
	varVersion := _Version{}

	err = json.Unmarshal(data, &varVersion)

	if err != nil {
		return err
	}

	*o = Version(varVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "build_version")
		delete(additionalProperties, "container_images")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "stable")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "x_version")
		delete(additionalProperties, "y_version")
		delete(additionalProperties, "z_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


