/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testclient

import (
	"encoding/json"
)

// CheckRestart struct for CheckRestart
type CheckRestart struct {
	Grace          *int64 `json:"Grace,omitempty"`
	IgnoreWarnings *bool  `json:"IgnoreWarnings,omitempty"`
	Limit          *int32 `json:"Limit,omitempty"`
}

// NewCheckRestart instantiates a new CheckRestart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckRestart() *CheckRestart {
	this := CheckRestart{}
	return &this
}

// NewCheckRestartWithDefaults instantiates a new CheckRestart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckRestartWithDefaults() *CheckRestart {
	this := CheckRestart{}
	return &this
}

// GetGrace returns the Grace field value if set, zero value otherwise.
func (o *CheckRestart) GetGrace() int64 {
	if o == nil || o.Grace == nil {
		var ret int64
		return ret
	}
	return *o.Grace
}

// GetGraceOk returns a tuple with the Grace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRestart) GetGraceOk() (*int64, bool) {
	if o == nil || o.Grace == nil {
		return nil, false
	}
	return o.Grace, true
}

// HasGrace returns a boolean if a field has been set.
func (o *CheckRestart) HasGrace() bool {
	if o != nil && o.Grace != nil {
		return true
	}

	return false
}

// SetGrace gets a reference to the given int64 and assigns it to the Grace field.
func (o *CheckRestart) SetGrace(v int64) {
	o.Grace = &v
}

// GetIgnoreWarnings returns the IgnoreWarnings field value if set, zero value otherwise.
func (o *CheckRestart) GetIgnoreWarnings() bool {
	if o == nil || o.IgnoreWarnings == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreWarnings
}

// GetIgnoreWarningsOk returns a tuple with the IgnoreWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRestart) GetIgnoreWarningsOk() (*bool, bool) {
	if o == nil || o.IgnoreWarnings == nil {
		return nil, false
	}
	return o.IgnoreWarnings, true
}

// HasIgnoreWarnings returns a boolean if a field has been set.
func (o *CheckRestart) HasIgnoreWarnings() bool {
	if o != nil && o.IgnoreWarnings != nil {
		return true
	}

	return false
}

// SetIgnoreWarnings gets a reference to the given bool and assigns it to the IgnoreWarnings field.
func (o *CheckRestart) SetIgnoreWarnings(v bool) {
	o.IgnoreWarnings = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CheckRestart) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRestart) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CheckRestart) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CheckRestart) SetLimit(v int32) {
	o.Limit = &v
}

func (o CheckRestart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Grace != nil {
		toSerialize["Grace"] = o.Grace
	}
	if o.IgnoreWarnings != nil {
		toSerialize["IgnoreWarnings"] = o.IgnoreWarnings
	}
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableCheckRestart struct {
	value *CheckRestart
	isSet bool
}

func (v NullableCheckRestart) Get() *CheckRestart {
	return v.value
}

func (v *NullableCheckRestart) Set(val *CheckRestart) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckRestart) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckRestart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckRestart(val *CheckRestart) *NullableCheckRestart {
	return &NullableCheckRestart{value: val, isSet: true}
}

func (v NullableCheckRestart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckRestart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
