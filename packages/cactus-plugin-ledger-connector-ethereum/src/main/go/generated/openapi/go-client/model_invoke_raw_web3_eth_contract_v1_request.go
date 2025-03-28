/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the InvokeRawWeb3EthContractV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvokeRawWeb3EthContractV1Request{}

// InvokeRawWeb3EthContractV1Request struct for InvokeRawWeb3EthContractV1Request
type InvokeRawWeb3EthContractV1Request struct {
	// The application binary interface of the solidity contract
	Abi interface{} `json:"abi"`
	// Deployed solidity contract address
	Address string `json:"address"`
	InvocationType EthContractInvocationWeb3Method `json:"invocationType"`
	// The list of arguments for contract invocation method (send, call, etc...)
	InvocationParams map[string]interface{} `json:"invocationParams,omitempty"`
	// Method of deployed solidity contract to execute
	ContractMethod string `json:"contractMethod"`
	// The list of arguments for deployed solidity contract method
	ContractMethodArgs []interface{} `json:"contractMethodArgs,omitempty"`
}

// NewInvokeRawWeb3EthContractV1Request instantiates a new InvokeRawWeb3EthContractV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvokeRawWeb3EthContractV1Request(abi interface{}, address string, invocationType EthContractInvocationWeb3Method, contractMethod string) *InvokeRawWeb3EthContractV1Request {
	this := InvokeRawWeb3EthContractV1Request{}
	this.Abi = abi
	this.Address = address
	this.InvocationType = invocationType
	this.ContractMethod = contractMethod
	return &this
}

// NewInvokeRawWeb3EthContractV1RequestWithDefaults instantiates a new InvokeRawWeb3EthContractV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvokeRawWeb3EthContractV1RequestWithDefaults() *InvokeRawWeb3EthContractV1Request {
	this := InvokeRawWeb3EthContractV1Request{}
	return &this
}

// GetAbi returns the Abi field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InvokeRawWeb3EthContractV1Request) GetAbi() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvokeRawWeb3EthContractV1Request) GetAbiOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Abi) {
		return nil, false
	}
	return &o.Abi, true
}

// SetAbi sets field value
func (o *InvokeRawWeb3EthContractV1Request) SetAbi(v interface{}) {
	o.Abi = v
}

// GetAddress returns the Address field value
func (o *InvokeRawWeb3EthContractV1Request) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InvokeRawWeb3EthContractV1Request) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InvokeRawWeb3EthContractV1Request) SetAddress(v string) {
	o.Address = v
}

// GetInvocationType returns the InvocationType field value
func (o *InvokeRawWeb3EthContractV1Request) GetInvocationType() EthContractInvocationWeb3Method {
	if o == nil {
		var ret EthContractInvocationWeb3Method
		return ret
	}

	return o.InvocationType
}

// GetInvocationTypeOk returns a tuple with the InvocationType field value
// and a boolean to check if the value has been set.
func (o *InvokeRawWeb3EthContractV1Request) GetInvocationTypeOk() (*EthContractInvocationWeb3Method, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvocationType, true
}

// SetInvocationType sets field value
func (o *InvokeRawWeb3EthContractV1Request) SetInvocationType(v EthContractInvocationWeb3Method) {
	o.InvocationType = v
}

// GetInvocationParams returns the InvocationParams field value if set, zero value otherwise.
func (o *InvokeRawWeb3EthContractV1Request) GetInvocationParams() map[string]interface{} {
	if o == nil || IsNil(o.InvocationParams) {
		var ret map[string]interface{}
		return ret
	}
	return o.InvocationParams
}

// GetInvocationParamsOk returns a tuple with the InvocationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeRawWeb3EthContractV1Request) GetInvocationParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InvocationParams) {
		return map[string]interface{}{}, false
	}
	return o.InvocationParams, true
}

// HasInvocationParams returns a boolean if a field has been set.
func (o *InvokeRawWeb3EthContractV1Request) HasInvocationParams() bool {
	if o != nil && !IsNil(o.InvocationParams) {
		return true
	}

	return false
}

// SetInvocationParams gets a reference to the given map[string]interface{} and assigns it to the InvocationParams field.
func (o *InvokeRawWeb3EthContractV1Request) SetInvocationParams(v map[string]interface{}) {
	o.InvocationParams = v
}

// GetContractMethod returns the ContractMethod field value
func (o *InvokeRawWeb3EthContractV1Request) GetContractMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractMethod
}

// GetContractMethodOk returns a tuple with the ContractMethod field value
// and a boolean to check if the value has been set.
func (o *InvokeRawWeb3EthContractV1Request) GetContractMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractMethod, true
}

// SetContractMethod sets field value
func (o *InvokeRawWeb3EthContractV1Request) SetContractMethod(v string) {
	o.ContractMethod = v
}

// GetContractMethodArgs returns the ContractMethodArgs field value if set, zero value otherwise.
func (o *InvokeRawWeb3EthContractV1Request) GetContractMethodArgs() []interface{} {
	if o == nil || IsNil(o.ContractMethodArgs) {
		var ret []interface{}
		return ret
	}
	return o.ContractMethodArgs
}

// GetContractMethodArgsOk returns a tuple with the ContractMethodArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeRawWeb3EthContractV1Request) GetContractMethodArgsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ContractMethodArgs) {
		return nil, false
	}
	return o.ContractMethodArgs, true
}

// HasContractMethodArgs returns a boolean if a field has been set.
func (o *InvokeRawWeb3EthContractV1Request) HasContractMethodArgs() bool {
	if o != nil && !IsNil(o.ContractMethodArgs) {
		return true
	}

	return false
}

// SetContractMethodArgs gets a reference to the given []interface{} and assigns it to the ContractMethodArgs field.
func (o *InvokeRawWeb3EthContractV1Request) SetContractMethodArgs(v []interface{}) {
	o.ContractMethodArgs = v
}

func (o InvokeRawWeb3EthContractV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvokeRawWeb3EthContractV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Abi != nil {
		toSerialize["abi"] = o.Abi
	}
	toSerialize["address"] = o.Address
	toSerialize["invocationType"] = o.InvocationType
	if !IsNil(o.InvocationParams) {
		toSerialize["invocationParams"] = o.InvocationParams
	}
	toSerialize["contractMethod"] = o.ContractMethod
	if !IsNil(o.ContractMethodArgs) {
		toSerialize["contractMethodArgs"] = o.ContractMethodArgs
	}
	return toSerialize, nil
}

type NullableInvokeRawWeb3EthContractV1Request struct {
	value *InvokeRawWeb3EthContractV1Request
	isSet bool
}

func (v NullableInvokeRawWeb3EthContractV1Request) Get() *InvokeRawWeb3EthContractV1Request {
	return v.value
}

func (v *NullableInvokeRawWeb3EthContractV1Request) Set(val *InvokeRawWeb3EthContractV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInvokeRawWeb3EthContractV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInvokeRawWeb3EthContractV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvokeRawWeb3EthContractV1Request(val *InvokeRawWeb3EthContractV1Request) *NullableInvokeRawWeb3EthContractV1Request {
	return &NullableInvokeRawWeb3EthContractV1Request{value: val, isSet: true}
}

func (v NullableInvokeRawWeb3EthContractV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvokeRawWeb3EthContractV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


