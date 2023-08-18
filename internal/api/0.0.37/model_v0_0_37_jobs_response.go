/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.37
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0037JobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037JobsResponse{}

// V0037JobsResponse struct for V0037JobsResponse
type V0037JobsResponse struct {
	// slurm errors
	Errors []V0037Error `json:"errors,omitempty"`
	// job descriptions
	Jobs []V0037JobResponseProperties `json:"jobs,omitempty"`
}

// NewV0037JobsResponse instantiates a new V0037JobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037JobsResponse() *V0037JobsResponse {
	this := V0037JobsResponse{}
	return &this
}

// NewV0037JobsResponseWithDefaults instantiates a new V0037JobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037JobsResponseWithDefaults() *V0037JobsResponse {
	this := V0037JobsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0037JobsResponse) GetErrors() []V0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037JobsResponse) GetErrorsOk() ([]V0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0037JobsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0037Error and assigns it to the Errors field.
func (o *V0037JobsResponse) SetErrors(v []V0037Error) {
	o.Errors = v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0037JobsResponse) GetJobs() []V0037JobResponseProperties {
	if o == nil || IsNil(o.Jobs) {
		var ret []V0037JobResponseProperties
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037JobsResponse) GetJobsOk() ([]V0037JobResponseProperties, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0037JobsResponse) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []V0037JobResponseProperties and assigns it to the Jobs field.
func (o *V0037JobsResponse) SetJobs(v []V0037JobResponseProperties) {
	o.Jobs = v
}

func (o V0037JobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037JobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	return toSerialize, nil
}

type NullableV0037JobsResponse struct {
	value *V0037JobsResponse
	isSet bool
}

func (v NullableV0037JobsResponse) Get() *V0037JobsResponse {
	return v.value
}

func (v *NullableV0037JobsResponse) Set(val *V0037JobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037JobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037JobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037JobsResponse(val *V0037JobsResponse) *NullableV0037JobsResponse {
	return &NullableV0037JobsResponse{value: val, isSet: true}
}

func (v NullableV0037JobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037JobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


