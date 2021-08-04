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

// ConsulLinkedService struct for ConsulLinkedService
type ConsulLinkedService struct {
	CAFile   *string `json:"CAFile,omitempty"`
	CertFile *string `json:"CertFile,omitempty"`
	KeyFile  *string `json:"KeyFile,omitempty"`
	Name     *string `json:"Name,omitempty"`
	SNI      *string `json:"SNI,omitempty"`
}

// NewConsulLinkedService instantiates a new ConsulLinkedService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulLinkedService() *ConsulLinkedService {
	this := ConsulLinkedService{}
	return &this
}

// NewConsulLinkedServiceWithDefaults instantiates a new ConsulLinkedService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulLinkedServiceWithDefaults() *ConsulLinkedService {
	this := ConsulLinkedService{}
	return &this
}

// GetCAFile returns the CAFile field value if set, zero value otherwise.
func (o *ConsulLinkedService) GetCAFile() string {
	if o == nil || o.CAFile == nil {
		var ret string
		return ret
	}
	return *o.CAFile
}

// GetCAFileOk returns a tuple with the CAFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulLinkedService) GetCAFileOk() (*string, bool) {
	if o == nil || o.CAFile == nil {
		return nil, false
	}
	return o.CAFile, true
}

// HasCAFile returns a boolean if a field has been set.
func (o *ConsulLinkedService) HasCAFile() bool {
	if o != nil && o.CAFile != nil {
		return true
	}

	return false
}

// SetCAFile gets a reference to the given string and assigns it to the CAFile field.
func (o *ConsulLinkedService) SetCAFile(v string) {
	o.CAFile = &v
}

// GetCertFile returns the CertFile field value if set, zero value otherwise.
func (o *ConsulLinkedService) GetCertFile() string {
	if o == nil || o.CertFile == nil {
		var ret string
		return ret
	}
	return *o.CertFile
}

// GetCertFileOk returns a tuple with the CertFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulLinkedService) GetCertFileOk() (*string, bool) {
	if o == nil || o.CertFile == nil {
		return nil, false
	}
	return o.CertFile, true
}

// HasCertFile returns a boolean if a field has been set.
func (o *ConsulLinkedService) HasCertFile() bool {
	if o != nil && o.CertFile != nil {
		return true
	}

	return false
}

// SetCertFile gets a reference to the given string and assigns it to the CertFile field.
func (o *ConsulLinkedService) SetCertFile(v string) {
	o.CertFile = &v
}

// GetKeyFile returns the KeyFile field value if set, zero value otherwise.
func (o *ConsulLinkedService) GetKeyFile() string {
	if o == nil || o.KeyFile == nil {
		var ret string
		return ret
	}
	return *o.KeyFile
}

// GetKeyFileOk returns a tuple with the KeyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulLinkedService) GetKeyFileOk() (*string, bool) {
	if o == nil || o.KeyFile == nil {
		return nil, false
	}
	return o.KeyFile, true
}

// HasKeyFile returns a boolean if a field has been set.
func (o *ConsulLinkedService) HasKeyFile() bool {
	if o != nil && o.KeyFile != nil {
		return true
	}

	return false
}

// SetKeyFile gets a reference to the given string and assigns it to the KeyFile field.
func (o *ConsulLinkedService) SetKeyFile(v string) {
	o.KeyFile = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConsulLinkedService) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulLinkedService) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConsulLinkedService) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConsulLinkedService) SetName(v string) {
	o.Name = &v
}

// GetSNI returns the SNI field value if set, zero value otherwise.
func (o *ConsulLinkedService) GetSNI() string {
	if o == nil || o.SNI == nil {
		var ret string
		return ret
	}
	return *o.SNI
}

// GetSNIOk returns a tuple with the SNI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulLinkedService) GetSNIOk() (*string, bool) {
	if o == nil || o.SNI == nil {
		return nil, false
	}
	return o.SNI, true
}

// HasSNI returns a boolean if a field has been set.
func (o *ConsulLinkedService) HasSNI() bool {
	if o != nil && o.SNI != nil {
		return true
	}

	return false
}

// SetSNI gets a reference to the given string and assigns it to the SNI field.
func (o *ConsulLinkedService) SetSNI(v string) {
	o.SNI = &v
}

func (o ConsulLinkedService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CAFile != nil {
		toSerialize["CAFile"] = o.CAFile
	}
	if o.CertFile != nil {
		toSerialize["CertFile"] = o.CertFile
	}
	if o.KeyFile != nil {
		toSerialize["KeyFile"] = o.KeyFile
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SNI != nil {
		toSerialize["SNI"] = o.SNI
	}
	return json.Marshal(toSerialize)
}

type NullableConsulLinkedService struct {
	value *ConsulLinkedService
	isSet bool
}

func (v NullableConsulLinkedService) Get() *ConsulLinkedService {
	return v.value
}

func (v *NullableConsulLinkedService) Set(val *ConsulLinkedService) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulLinkedService) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulLinkedService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulLinkedService(val *ConsulLinkedService) *NullableConsulLinkedService {
	return &NullableConsulLinkedService{value: val, isSet: true}
}

func (v NullableConsulLinkedService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulLinkedService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
