/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the DiagnoseNodeV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnoseNodeV1Request{}

// DiagnoseNodeV1Request struct for DiagnoseNodeV1Request
type DiagnoseNodeV1Request struct {
	// Optional property specifying which Corda Node should be the one being diagnosed in case the Connector has multiple connections established for different nodes (which is not yet a supported feature, but we want to keep this possibility open for the future).
	NodeIds []string `json:"nodeIds,omitempty"`
}

// NewDiagnoseNodeV1Request instantiates a new DiagnoseNodeV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnoseNodeV1Request() *DiagnoseNodeV1Request {
	this := DiagnoseNodeV1Request{}
	return &this
}

// NewDiagnoseNodeV1RequestWithDefaults instantiates a new DiagnoseNodeV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnoseNodeV1RequestWithDefaults() *DiagnoseNodeV1Request {
	this := DiagnoseNodeV1Request{}
	return &this
}

// GetNodeIds returns the NodeIds field value if set, zero value otherwise.
func (o *DiagnoseNodeV1Request) GetNodeIds() []string {
	if o == nil || IsNil(o.NodeIds) {
		var ret []string
		return ret
	}
	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnoseNodeV1Request) GetNodeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// HasNodeIds returns a boolean if a field has been set.
func (o *DiagnoseNodeV1Request) HasNodeIds() bool {
	if o != nil && !IsNil(o.NodeIds) {
		return true
	}

	return false
}

// SetNodeIds gets a reference to the given []string and assigns it to the NodeIds field.
func (o *DiagnoseNodeV1Request) SetNodeIds(v []string) {
	o.NodeIds = v
}

func (o DiagnoseNodeV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnoseNodeV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodeIds) {
		toSerialize["nodeIds"] = o.NodeIds
	}
	return toSerialize, nil
}

type NullableDiagnoseNodeV1Request struct {
	value *DiagnoseNodeV1Request
	isSet bool
}

func (v NullableDiagnoseNodeV1Request) Get() *DiagnoseNodeV1Request {
	return v.value
}

func (v *NullableDiagnoseNodeV1Request) Set(val *DiagnoseNodeV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnoseNodeV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnoseNodeV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnoseNodeV1Request(val *DiagnoseNodeV1Request) *NullableDiagnoseNodeV1Request {
	return &NullableDiagnoseNodeV1Request{value: val, isSet: true}
}

func (v NullableDiagnoseNodeV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnoseNodeV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

