/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v7.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SmartContractTransactionSmartContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartContractTransactionSmartContract{}

// SmartContractTransactionSmartContract struct for SmartContractTransactionSmartContract
type SmartContractTransactionSmartContract struct {
	// The Clarity version of the contract, only specified for versioned contract transactions, otherwise null
	ClarityVersion NullableFloat32 `json:"clarity_version"`
	// Contract identifier formatted as `<principaladdress>.<contract_name>`
	ContractId string `json:"contract_id"`
	// Clarity code of the smart contract being deployed
	SourceCode string `json:"source_code"`
}

type _SmartContractTransactionSmartContract SmartContractTransactionSmartContract

// NewSmartContractTransactionSmartContract instantiates a new SmartContractTransactionSmartContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractTransactionSmartContract(clarityVersion NullableFloat32, contractId string, sourceCode string) *SmartContractTransactionSmartContract {
	this := SmartContractTransactionSmartContract{}
	this.ClarityVersion = clarityVersion
	this.ContractId = contractId
	this.SourceCode = sourceCode
	return &this
}

// NewSmartContractTransactionSmartContractWithDefaults instantiates a new SmartContractTransactionSmartContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractTransactionSmartContractWithDefaults() *SmartContractTransactionSmartContract {
	this := SmartContractTransactionSmartContract{}
	return &this
}

// GetClarityVersion returns the ClarityVersion field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SmartContractTransactionSmartContract) GetClarityVersion() float32 {
	if o == nil || o.ClarityVersion.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ClarityVersion.Get()
}

// GetClarityVersionOk returns a tuple with the ClarityVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmartContractTransactionSmartContract) GetClarityVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClarityVersion.Get(), o.ClarityVersion.IsSet()
}

// SetClarityVersion sets field value
func (o *SmartContractTransactionSmartContract) SetClarityVersion(v float32) {
	o.ClarityVersion.Set(&v)
}

// GetContractId returns the ContractId field value
func (o *SmartContractTransactionSmartContract) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransactionSmartContract) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *SmartContractTransactionSmartContract) SetContractId(v string) {
	o.ContractId = v
}

// GetSourceCode returns the SourceCode field value
func (o *SmartContractTransactionSmartContract) GetSourceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransactionSmartContract) GetSourceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCode, true
}

// SetSourceCode sets field value
func (o *SmartContractTransactionSmartContract) SetSourceCode(v string) {
	o.SourceCode = v
}

func (o SmartContractTransactionSmartContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractTransactionSmartContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clarity_version"] = o.ClarityVersion.Get()
	toSerialize["contract_id"] = o.ContractId
	toSerialize["source_code"] = o.SourceCode
	return toSerialize, nil
}

func (o *SmartContractTransactionSmartContract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clarity_version",
		"contract_id",
		"source_code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSmartContractTransactionSmartContract := _SmartContractTransactionSmartContract{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractTransactionSmartContract)

	if err != nil {
		return err
	}

	*o = SmartContractTransactionSmartContract(varSmartContractTransactionSmartContract)

	return err
}

type NullableSmartContractTransactionSmartContract struct {
	value *SmartContractTransactionSmartContract
	isSet bool
}

func (v NullableSmartContractTransactionSmartContract) Get() *SmartContractTransactionSmartContract {
	return v.value
}

func (v *NullableSmartContractTransactionSmartContract) Set(val *SmartContractTransactionSmartContract) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractTransactionSmartContract) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractTransactionSmartContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractTransactionSmartContract(val *SmartContractTransactionSmartContract) *NullableSmartContractTransactionSmartContract {
	return &NullableSmartContractTransactionSmartContract{value: val, isSet: true}
}

func (v NullableSmartContractTransactionSmartContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractTransactionSmartContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


