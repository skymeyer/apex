/*
Apex Public API

Public (unauthenticated) API to integrate with your local Apex aquarium controller (AOS 5 compatible).

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// StatusResponse struct for StatusResponse
type StatusResponse struct {
	Istat StatusContainer `json:"istat"`
}

// NewStatusResponse instantiates a new StatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponse(istat StatusContainer) *StatusResponse {
	this := StatusResponse{}
	this.Istat = istat
	return &this
}

// NewStatusResponseWithDefaults instantiates a new StatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseWithDefaults() *StatusResponse {
	this := StatusResponse{}
	return &this
}

// GetIstat returns the Istat field value
func (o *StatusResponse) GetIstat() StatusContainer {
	if o == nil {
		var ret StatusContainer
		return ret
	}

	return o.Istat
}

// GetIstatOk returns a tuple with the Istat field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetIstatOk() (*StatusContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Istat, true
}

// SetIstat sets field value
func (o *StatusResponse) SetIstat(v StatusContainer) {
	o.Istat = v
}

func (o StatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["istat"] = o.Istat
	}
	return json.Marshal(toSerialize)
}

type NullableStatusResponse struct {
	value *StatusResponse
	isSet bool
}

func (v NullableStatusResponse) Get() *StatusResponse {
	return v.value
}

func (v *NullableStatusResponse) Set(val *StatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponse(val *StatusResponse) *NullableStatusResponse {
	return &NullableStatusResponse{value: val, isSet: true}
}

func (v NullableStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
