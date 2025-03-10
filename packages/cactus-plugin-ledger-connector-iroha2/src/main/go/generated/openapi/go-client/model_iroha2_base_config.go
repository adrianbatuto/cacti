/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
)

// checks if the Iroha2BaseConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Iroha2BaseConfig{}

// Iroha2BaseConfig Iroha V2 connection configuration.
type Iroha2BaseConfig struct {
	Torii Iroha2BaseConfigTorii `json:"torii"`
	AccountId *Iroha2AccountId `json:"accountId,omitempty"`
	SigningCredential *Iroha2BaseConfigSigningCredential `json:"signingCredential,omitempty"`
}

// NewIroha2BaseConfig instantiates a new Iroha2BaseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIroha2BaseConfig(torii Iroha2BaseConfigTorii) *Iroha2BaseConfig {
	this := Iroha2BaseConfig{}
	this.Torii = torii
	return &this
}

// NewIroha2BaseConfigWithDefaults instantiates a new Iroha2BaseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIroha2BaseConfigWithDefaults() *Iroha2BaseConfig {
	this := Iroha2BaseConfig{}
	return &this
}

// GetTorii returns the Torii field value
func (o *Iroha2BaseConfig) GetTorii() Iroha2BaseConfigTorii {
	if o == nil {
		var ret Iroha2BaseConfigTorii
		return ret
	}

	return o.Torii
}

// GetToriiOk returns a tuple with the Torii field value
// and a boolean to check if the value has been set.
func (o *Iroha2BaseConfig) GetToriiOk() (*Iroha2BaseConfigTorii, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Torii, true
}

// SetTorii sets field value
func (o *Iroha2BaseConfig) SetTorii(v Iroha2BaseConfigTorii) {
	o.Torii = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Iroha2BaseConfig) GetAccountId() Iroha2AccountId {
	if o == nil || IsNil(o.AccountId) {
		var ret Iroha2AccountId
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iroha2BaseConfig) GetAccountIdOk() (*Iroha2AccountId, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Iroha2BaseConfig) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given Iroha2AccountId and assigns it to the AccountId field.
func (o *Iroha2BaseConfig) SetAccountId(v Iroha2AccountId) {
	o.AccountId = &v
}

// GetSigningCredential returns the SigningCredential field value if set, zero value otherwise.
func (o *Iroha2BaseConfig) GetSigningCredential() Iroha2BaseConfigSigningCredential {
	if o == nil || IsNil(o.SigningCredential) {
		var ret Iroha2BaseConfigSigningCredential
		return ret
	}
	return *o.SigningCredential
}

// GetSigningCredentialOk returns a tuple with the SigningCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iroha2BaseConfig) GetSigningCredentialOk() (*Iroha2BaseConfigSigningCredential, bool) {
	if o == nil || IsNil(o.SigningCredential) {
		return nil, false
	}
	return o.SigningCredential, true
}

// HasSigningCredential returns a boolean if a field has been set.
func (o *Iroha2BaseConfig) HasSigningCredential() bool {
	if o != nil && !IsNil(o.SigningCredential) {
		return true
	}

	return false
}

// SetSigningCredential gets a reference to the given Iroha2BaseConfigSigningCredential and assigns it to the SigningCredential field.
func (o *Iroha2BaseConfig) SetSigningCredential(v Iroha2BaseConfigSigningCredential) {
	o.SigningCredential = &v
}

func (o Iroha2BaseConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Iroha2BaseConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["torii"] = o.Torii
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.SigningCredential) {
		toSerialize["signingCredential"] = o.SigningCredential
	}
	return toSerialize, nil
}

type NullableIroha2BaseConfig struct {
	value *Iroha2BaseConfig
	isSet bool
}

func (v NullableIroha2BaseConfig) Get() *Iroha2BaseConfig {
	return v.value
}

func (v *NullableIroha2BaseConfig) Set(val *Iroha2BaseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIroha2BaseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIroha2BaseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIroha2BaseConfig(val *Iroha2BaseConfig) *NullableIroha2BaseConfig {
	return &NullableIroha2BaseConfig{value: val, isSet: true}
}

func (v NullableIroha2BaseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIroha2BaseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


