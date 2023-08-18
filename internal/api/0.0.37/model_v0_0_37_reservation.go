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

// checks if the V0037Reservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037Reservation{}

// V0037Reservation struct for V0037Reservation
type V0037Reservation struct {
	// Allowed accounts
	Accounts *string `json:"accounts,omitempty"`
	// Reserved burst buffer
	BurstBuffer *string `json:"burst_buffer,omitempty"`
	// Number of reserved cores
	CoreCount *int32 `json:"core_count,omitempty"`
	// Number of reserved specialized cores
	CoreSpecCnt *int32 `json:"core_spec_cnt,omitempty"`
	// End time of the reservation
	EndTime *int32 `json:"end_time,omitempty"`
	// List of features
	Features *string `json:"features,omitempty"`
	// Reservation options
	Flags []string `json:"flags,omitempty"`
	// List of groups permitted to use the reserved nodes
	Groups *string `json:"groups,omitempty"`
	// List of licenses
	Licenses *string `json:"licenses,omitempty"`
	// Maximum delay in which jobs outside of the reservation will be permitted to overlap once any jobs are queued for the reservation
	MaxStartDelay *int32 `json:"max_start_delay,omitempty"`
	// Reservationn name
	Name *string `json:"name,omitempty"`
	// Count of nodes reserved
	NodeCount *int32 `json:"node_count,omitempty"`
	// List of reserved nodes
	NodeList *string `json:"node_list,omitempty"`
	// Partition
	Partition *string `json:"partition,omitempty"`
	PurgeCompleted *V0037ReservationPurgeCompleted `json:"purge_completed,omitempty"`
	// Start time of reservation
	StartTime *int32 `json:"start_time,omitempty"`
	// amount of power to reserve in watts
	Watts *int32 `json:"watts,omitempty"`
	// List of TRES
	Tres *string `json:"tres,omitempty"`
	// List of users
	Users *string `json:"users,omitempty"`
}

// NewV0037Reservation instantiates a new V0037Reservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037Reservation() *V0037Reservation {
	this := V0037Reservation{}
	return &this
}

// NewV0037ReservationWithDefaults instantiates a new V0037Reservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037ReservationWithDefaults() *V0037Reservation {
	this := V0037Reservation{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *V0037Reservation) GetAccounts() string {
	if o == nil || IsNil(o.Accounts) {
		var ret string
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *V0037Reservation) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given string and assigns it to the Accounts field.
func (o *V0037Reservation) SetAccounts(v string) {
	o.Accounts = &v
}

// GetBurstBuffer returns the BurstBuffer field value if set, zero value otherwise.
func (o *V0037Reservation) GetBurstBuffer() string {
	if o == nil || IsNil(o.BurstBuffer) {
		var ret string
		return ret
	}
	return *o.BurstBuffer
}

// GetBurstBufferOk returns a tuple with the BurstBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetBurstBufferOk() (*string, bool) {
	if o == nil || IsNil(o.BurstBuffer) {
		return nil, false
	}
	return o.BurstBuffer, true
}

// HasBurstBuffer returns a boolean if a field has been set.
func (o *V0037Reservation) HasBurstBuffer() bool {
	if o != nil && !IsNil(o.BurstBuffer) {
		return true
	}

	return false
}

// SetBurstBuffer gets a reference to the given string and assigns it to the BurstBuffer field.
func (o *V0037Reservation) SetBurstBuffer(v string) {
	o.BurstBuffer = &v
}

// GetCoreCount returns the CoreCount field value if set, zero value otherwise.
func (o *V0037Reservation) GetCoreCount() int32 {
	if o == nil || IsNil(o.CoreCount) {
		var ret int32
		return ret
	}
	return *o.CoreCount
}

// GetCoreCountOk returns a tuple with the CoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetCoreCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CoreCount) {
		return nil, false
	}
	return o.CoreCount, true
}

// HasCoreCount returns a boolean if a field has been set.
func (o *V0037Reservation) HasCoreCount() bool {
	if o != nil && !IsNil(o.CoreCount) {
		return true
	}

	return false
}

// SetCoreCount gets a reference to the given int32 and assigns it to the CoreCount field.
func (o *V0037Reservation) SetCoreCount(v int32) {
	o.CoreCount = &v
}

// GetCoreSpecCnt returns the CoreSpecCnt field value if set, zero value otherwise.
func (o *V0037Reservation) GetCoreSpecCnt() int32 {
	if o == nil || IsNil(o.CoreSpecCnt) {
		var ret int32
		return ret
	}
	return *o.CoreSpecCnt
}

// GetCoreSpecCntOk returns a tuple with the CoreSpecCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetCoreSpecCntOk() (*int32, bool) {
	if o == nil || IsNil(o.CoreSpecCnt) {
		return nil, false
	}
	return o.CoreSpecCnt, true
}

// HasCoreSpecCnt returns a boolean if a field has been set.
func (o *V0037Reservation) HasCoreSpecCnt() bool {
	if o != nil && !IsNil(o.CoreSpecCnt) {
		return true
	}

	return false
}

// SetCoreSpecCnt gets a reference to the given int32 and assigns it to the CoreSpecCnt field.
func (o *V0037Reservation) SetCoreSpecCnt(v int32) {
	o.CoreSpecCnt = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *V0037Reservation) GetEndTime() int32 {
	if o == nil || IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetEndTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *V0037Reservation) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *V0037Reservation) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *V0037Reservation) GetFeatures() string {
	if o == nil || IsNil(o.Features) {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *V0037Reservation) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *V0037Reservation) SetFeatures(v string) {
	o.Features = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0037Reservation) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0037Reservation) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0037Reservation) SetFlags(v []string) {
	o.Flags = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *V0037Reservation) GetGroups() string {
	if o == nil || IsNil(o.Groups) {
		var ret string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetGroupsOk() (*string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *V0037Reservation) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given string and assigns it to the Groups field.
func (o *V0037Reservation) SetGroups(v string) {
	o.Groups = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0037Reservation) GetLicenses() string {
	if o == nil || IsNil(o.Licenses) {
		var ret string
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetLicensesOk() (*string, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0037Reservation) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given string and assigns it to the Licenses field.
func (o *V0037Reservation) SetLicenses(v string) {
	o.Licenses = &v
}

// GetMaxStartDelay returns the MaxStartDelay field value if set, zero value otherwise.
func (o *V0037Reservation) GetMaxStartDelay() int32 {
	if o == nil || IsNil(o.MaxStartDelay) {
		var ret int32
		return ret
	}
	return *o.MaxStartDelay
}

// GetMaxStartDelayOk returns a tuple with the MaxStartDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetMaxStartDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxStartDelay) {
		return nil, false
	}
	return o.MaxStartDelay, true
}

// HasMaxStartDelay returns a boolean if a field has been set.
func (o *V0037Reservation) HasMaxStartDelay() bool {
	if o != nil && !IsNil(o.MaxStartDelay) {
		return true
	}

	return false
}

// SetMaxStartDelay gets a reference to the given int32 and assigns it to the MaxStartDelay field.
func (o *V0037Reservation) SetMaxStartDelay(v int32) {
	o.MaxStartDelay = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0037Reservation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0037Reservation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0037Reservation) SetName(v string) {
	o.Name = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *V0037Reservation) GetNodeCount() int32 {
	if o == nil || IsNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetNodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *V0037Reservation) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *V0037Reservation) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetNodeList returns the NodeList field value if set, zero value otherwise.
func (o *V0037Reservation) GetNodeList() string {
	if o == nil || IsNil(o.NodeList) {
		var ret string
		return ret
	}
	return *o.NodeList
}

// GetNodeListOk returns a tuple with the NodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetNodeListOk() (*string, bool) {
	if o == nil || IsNil(o.NodeList) {
		return nil, false
	}
	return o.NodeList, true
}

// HasNodeList returns a boolean if a field has been set.
func (o *V0037Reservation) HasNodeList() bool {
	if o != nil && !IsNil(o.NodeList) {
		return true
	}

	return false
}

// SetNodeList gets a reference to the given string and assigns it to the NodeList field.
func (o *V0037Reservation) SetNodeList(v string) {
	o.NodeList = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0037Reservation) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0037Reservation) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0037Reservation) SetPartition(v string) {
	o.Partition = &v
}

// GetPurgeCompleted returns the PurgeCompleted field value if set, zero value otherwise.
func (o *V0037Reservation) GetPurgeCompleted() V0037ReservationPurgeCompleted {
	if o == nil || IsNil(o.PurgeCompleted) {
		var ret V0037ReservationPurgeCompleted
		return ret
	}
	return *o.PurgeCompleted
}

// GetPurgeCompletedOk returns a tuple with the PurgeCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetPurgeCompletedOk() (*V0037ReservationPurgeCompleted, bool) {
	if o == nil || IsNil(o.PurgeCompleted) {
		return nil, false
	}
	return o.PurgeCompleted, true
}

// HasPurgeCompleted returns a boolean if a field has been set.
func (o *V0037Reservation) HasPurgeCompleted() bool {
	if o != nil && !IsNil(o.PurgeCompleted) {
		return true
	}

	return false
}

// SetPurgeCompleted gets a reference to the given V0037ReservationPurgeCompleted and assigns it to the PurgeCompleted field.
func (o *V0037Reservation) SetPurgeCompleted(v V0037ReservationPurgeCompleted) {
	o.PurgeCompleted = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *V0037Reservation) GetStartTime() int32 {
	if o == nil || IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *V0037Reservation) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *V0037Reservation) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetWatts returns the Watts field value if set, zero value otherwise.
func (o *V0037Reservation) GetWatts() int32 {
	if o == nil || IsNil(o.Watts) {
		var ret int32
		return ret
	}
	return *o.Watts
}

// GetWattsOk returns a tuple with the Watts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetWattsOk() (*int32, bool) {
	if o == nil || IsNil(o.Watts) {
		return nil, false
	}
	return o.Watts, true
}

// HasWatts returns a boolean if a field has been set.
func (o *V0037Reservation) HasWatts() bool {
	if o != nil && !IsNil(o.Watts) {
		return true
	}

	return false
}

// SetWatts gets a reference to the given int32 and assigns it to the Watts field.
func (o *V0037Reservation) SetWatts(v int32) {
	o.Watts = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0037Reservation) GetTres() string {
	if o == nil || IsNil(o.Tres) {
		var ret string
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetTresOk() (*string, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0037Reservation) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given string and assigns it to the Tres field.
func (o *V0037Reservation) SetTres(v string) {
	o.Tres = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V0037Reservation) GetUsers() string {
	if o == nil || IsNil(o.Users) {
		var ret string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetUsersOk() (*string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V0037Reservation) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given string and assigns it to the Users field.
func (o *V0037Reservation) SetUsers(v string) {
	o.Users = &v
}

func (o V0037Reservation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037Reservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.BurstBuffer) {
		toSerialize["burst_buffer"] = o.BurstBuffer
	}
	if !IsNil(o.CoreCount) {
		toSerialize["core_count"] = o.CoreCount
	}
	if !IsNil(o.CoreSpecCnt) {
		toSerialize["core_spec_cnt"] = o.CoreSpecCnt
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.MaxStartDelay) {
		toSerialize["max_start_delay"] = o.MaxStartDelay
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NodeCount) {
		toSerialize["node_count"] = o.NodeCount
	}
	if !IsNil(o.NodeList) {
		toSerialize["node_list"] = o.NodeList
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.PurgeCompleted) {
		toSerialize["purge_completed"] = o.PurgeCompleted
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Watts) {
		toSerialize["watts"] = o.Watts
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableV0037Reservation struct {
	value *V0037Reservation
	isSet bool
}

func (v NullableV0037Reservation) Get() *V0037Reservation {
	return v.value
}

func (v *NullableV0037Reservation) Set(val *V0037Reservation) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037Reservation) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037Reservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037Reservation(val *V0037Reservation) *NullableV0037Reservation {
	return &NullableV0037Reservation{value: val, isSet: true}
}

func (v NullableV0037Reservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037Reservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


