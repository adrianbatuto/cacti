/*
Hyperledger Cactus Plugin - Object Store - IPFS 

Contains/describes the Hyperledger Cactus Object Store IPFS plugin.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-object-store-ipfs

import (
	"encoding/json"
)

// checks if the HasObjectRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasObjectRequestV1{}

// HasObjectRequestV1 struct for HasObjectRequestV1
type HasObjectRequestV1 struct {
	// The key to check for presence in the object store.
	Key string `json:"key"`
}

// NewHasObjectRequestV1 instantiates a new HasObjectRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasObjectRequestV1(key string) *HasObjectRequestV1 {
	this := HasObjectRequestV1{}
	this.Key = key
	return &this
}

// NewHasObjectRequestV1WithDefaults instantiates a new HasObjectRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasObjectRequestV1WithDefaults() *HasObjectRequestV1 {
	this := HasObjectRequestV1{}
	return &this
}

// GetKey returns the Key field value
func (o *HasObjectRequestV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *HasObjectRequestV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *HasObjectRequestV1) SetKey(v string) {
	o.Key = v
}

func (o HasObjectRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasObjectRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableHasObjectRequestV1 struct {
	value *HasObjectRequestV1
	isSet bool
}

func (v NullableHasObjectRequestV1) Get() *HasObjectRequestV1 {
	return v.value
}

func (v *NullableHasObjectRequestV1) Set(val *HasObjectRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableHasObjectRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableHasObjectRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasObjectRequestV1(val *HasObjectRequestV1) *NullableHasObjectRequestV1 {
	return &NullableHasObjectRequestV1{value: val, isSet: true}
}

func (v NullableHasObjectRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasObjectRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


