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

// checks if the GetMempoolFeePriorities200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMempoolFeePriorities200Response{}

// GetMempoolFeePriorities200Response struct for GetMempoolFeePriorities200Response
type GetMempoolFeePriorities200Response struct {
	All GetMempoolFeePriorities200ResponseAll `json:"all"`
	TokenTransfer *GetMempoolFeePriorities200ResponseAll `json:"token_transfer,omitempty"`
	ContractCall *GetMempoolFeePriorities200ResponseAll `json:"contract_call,omitempty"`
	SmartContract *GetMempoolFeePriorities200ResponseAll `json:"smart_contract,omitempty"`
}

type _GetMempoolFeePriorities200Response GetMempoolFeePriorities200Response

// NewGetMempoolFeePriorities200Response instantiates a new GetMempoolFeePriorities200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMempoolFeePriorities200Response(all GetMempoolFeePriorities200ResponseAll) *GetMempoolFeePriorities200Response {
	this := GetMempoolFeePriorities200Response{}
	this.All = all
	return &this
}

// NewGetMempoolFeePriorities200ResponseWithDefaults instantiates a new GetMempoolFeePriorities200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMempoolFeePriorities200ResponseWithDefaults() *GetMempoolFeePriorities200Response {
	this := GetMempoolFeePriorities200Response{}
	return &this
}

// GetAll returns the All field value
func (o *GetMempoolFeePriorities200Response) GetAll() GetMempoolFeePriorities200ResponseAll {
	if o == nil {
		var ret GetMempoolFeePriorities200ResponseAll
		return ret
	}

	return o.All
}

// GetAllOk returns a tuple with the All field value
// and a boolean to check if the value has been set.
func (o *GetMempoolFeePriorities200Response) GetAllOk() (*GetMempoolFeePriorities200ResponseAll, bool) {
	if o == nil {
		return nil, false
	}
	return &o.All, true
}

// SetAll sets field value
func (o *GetMempoolFeePriorities200Response) SetAll(v GetMempoolFeePriorities200ResponseAll) {
	o.All = v
}

// GetTokenTransfer returns the TokenTransfer field value if set, zero value otherwise.
func (o *GetMempoolFeePriorities200Response) GetTokenTransfer() GetMempoolFeePriorities200ResponseAll {
	if o == nil || IsNil(o.TokenTransfer) {
		var ret GetMempoolFeePriorities200ResponseAll
		return ret
	}
	return *o.TokenTransfer
}

// GetTokenTransferOk returns a tuple with the TokenTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMempoolFeePriorities200Response) GetTokenTransferOk() (*GetMempoolFeePriorities200ResponseAll, bool) {
	if o == nil || IsNil(o.TokenTransfer) {
		return nil, false
	}
	return o.TokenTransfer, true
}

// HasTokenTransfer returns a boolean if a field has been set.
func (o *GetMempoolFeePriorities200Response) HasTokenTransfer() bool {
	if o != nil && !IsNil(o.TokenTransfer) {
		return true
	}

	return false
}

// SetTokenTransfer gets a reference to the given GetMempoolFeePriorities200ResponseAll and assigns it to the TokenTransfer field.
func (o *GetMempoolFeePriorities200Response) SetTokenTransfer(v GetMempoolFeePriorities200ResponseAll) {
	o.TokenTransfer = &v
}

// GetContractCall returns the ContractCall field value if set, zero value otherwise.
func (o *GetMempoolFeePriorities200Response) GetContractCall() GetMempoolFeePriorities200ResponseAll {
	if o == nil || IsNil(o.ContractCall) {
		var ret GetMempoolFeePriorities200ResponseAll
		return ret
	}
	return *o.ContractCall
}

// GetContractCallOk returns a tuple with the ContractCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMempoolFeePriorities200Response) GetContractCallOk() (*GetMempoolFeePriorities200ResponseAll, bool) {
	if o == nil || IsNil(o.ContractCall) {
		return nil, false
	}
	return o.ContractCall, true
}

// HasContractCall returns a boolean if a field has been set.
func (o *GetMempoolFeePriorities200Response) HasContractCall() bool {
	if o != nil && !IsNil(o.ContractCall) {
		return true
	}

	return false
}

// SetContractCall gets a reference to the given GetMempoolFeePriorities200ResponseAll and assigns it to the ContractCall field.
func (o *GetMempoolFeePriorities200Response) SetContractCall(v GetMempoolFeePriorities200ResponseAll) {
	o.ContractCall = &v
}

// GetSmartContract returns the SmartContract field value if set, zero value otherwise.
func (o *GetMempoolFeePriorities200Response) GetSmartContract() GetMempoolFeePriorities200ResponseAll {
	if o == nil || IsNil(o.SmartContract) {
		var ret GetMempoolFeePriorities200ResponseAll
		return ret
	}
	return *o.SmartContract
}

// GetSmartContractOk returns a tuple with the SmartContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMempoolFeePriorities200Response) GetSmartContractOk() (*GetMempoolFeePriorities200ResponseAll, bool) {
	if o == nil || IsNil(o.SmartContract) {
		return nil, false
	}
	return o.SmartContract, true
}

// HasSmartContract returns a boolean if a field has been set.
func (o *GetMempoolFeePriorities200Response) HasSmartContract() bool {
	if o != nil && !IsNil(o.SmartContract) {
		return true
	}

	return false
}

// SetSmartContract gets a reference to the given GetMempoolFeePriorities200ResponseAll and assigns it to the SmartContract field.
func (o *GetMempoolFeePriorities200Response) SetSmartContract(v GetMempoolFeePriorities200ResponseAll) {
	o.SmartContract = &v
}

func (o GetMempoolFeePriorities200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMempoolFeePriorities200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["all"] = o.All
	if !IsNil(o.TokenTransfer) {
		toSerialize["token_transfer"] = o.TokenTransfer
	}
	if !IsNil(o.ContractCall) {
		toSerialize["contract_call"] = o.ContractCall
	}
	if !IsNil(o.SmartContract) {
		toSerialize["smart_contract"] = o.SmartContract
	}
	return toSerialize, nil
}

func (o *GetMempoolFeePriorities200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"all",
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

	varGetMempoolFeePriorities200Response := _GetMempoolFeePriorities200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMempoolFeePriorities200Response)

	if err != nil {
		return err
	}

	*o = GetMempoolFeePriorities200Response(varGetMempoolFeePriorities200Response)

	return err
}

type NullableGetMempoolFeePriorities200Response struct {
	value *GetMempoolFeePriorities200Response
	isSet bool
}

func (v NullableGetMempoolFeePriorities200Response) Get() *GetMempoolFeePriorities200Response {
	return v.value
}

func (v *NullableGetMempoolFeePriorities200Response) Set(val *GetMempoolFeePriorities200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMempoolFeePriorities200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMempoolFeePriorities200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMempoolFeePriorities200Response(val *GetMempoolFeePriorities200Response) *NullableGetMempoolFeePriorities200Response {
	return &NullableGetMempoolFeePriorities200Response{value: val, isSet: true}
}

func (v NullableGetMempoolFeePriorities200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMempoolFeePriorities200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


