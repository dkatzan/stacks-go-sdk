/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v8.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetPoxCycleSignerStackers200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPoxCycleSignerStackers200ResponseResultsInner{}

// GetPoxCycleSignerStackers200ResponseResultsInner struct for GetPoxCycleSignerStackers200ResponseResultsInner
type GetPoxCycleSignerStackers200ResponseResultsInner struct {
	StackerAddress string `json:"stacker_address"`
	StackedAmount string `json:"stacked_amount"`
	PoxAddress string `json:"pox_address"`
	StackerType GetPoxCycleSignerStackers200ResponseResultsInnerStackerType `json:"stacker_type"`
}

type _GetPoxCycleSignerStackers200ResponseResultsInner GetPoxCycleSignerStackers200ResponseResultsInner

// NewGetPoxCycleSignerStackers200ResponseResultsInner instantiates a new GetPoxCycleSignerStackers200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPoxCycleSignerStackers200ResponseResultsInner(stackerAddress string, stackedAmount string, poxAddress string, stackerType GetPoxCycleSignerStackers200ResponseResultsInnerStackerType) *GetPoxCycleSignerStackers200ResponseResultsInner {
	this := GetPoxCycleSignerStackers200ResponseResultsInner{}
	this.StackerAddress = stackerAddress
	this.StackedAmount = stackedAmount
	this.PoxAddress = poxAddress
	this.StackerType = stackerType
	return &this
}

// NewGetPoxCycleSignerStackers200ResponseResultsInnerWithDefaults instantiates a new GetPoxCycleSignerStackers200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPoxCycleSignerStackers200ResponseResultsInnerWithDefaults() *GetPoxCycleSignerStackers200ResponseResultsInner {
	this := GetPoxCycleSignerStackers200ResponseResultsInner{}
	return &this
}

// GetStackerAddress returns the StackerAddress field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetStackerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackerAddress
}

// GetStackerAddressOk returns a tuple with the StackerAddress field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetStackerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackerAddress, true
}

// SetStackerAddress sets field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) SetStackerAddress(v string) {
	o.StackerAddress = v
}

// GetStackedAmount returns the StackedAmount field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetStackedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackedAmount
}

// GetStackedAmountOk returns a tuple with the StackedAmount field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetStackedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackedAmount, true
}

// SetStackedAmount sets field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) SetStackedAmount(v string) {
	o.StackedAmount = v
}

// GetPoxAddress returns the PoxAddress field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetPoxAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoxAddress
}

// GetPoxAddressOk returns a tuple with the PoxAddress field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetPoxAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoxAddress, true
}

// SetPoxAddress sets field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) SetPoxAddress(v string) {
	o.PoxAddress = v
}

// GetStackerType returns the StackerType field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetStackerType() GetPoxCycleSignerStackers200ResponseResultsInnerStackerType {
	if o == nil {
		var ret GetPoxCycleSignerStackers200ResponseResultsInnerStackerType
		return ret
	}

	return o.StackerType
}

// GetStackerTypeOk returns a tuple with the StackerType field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) GetStackerTypeOk() (*GetPoxCycleSignerStackers200ResponseResultsInnerStackerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackerType, true
}

// SetStackerType sets field value
func (o *GetPoxCycleSignerStackers200ResponseResultsInner) SetStackerType(v GetPoxCycleSignerStackers200ResponseResultsInnerStackerType) {
	o.StackerType = v
}

func (o GetPoxCycleSignerStackers200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPoxCycleSignerStackers200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stacker_address"] = o.StackerAddress
	toSerialize["stacked_amount"] = o.StackedAmount
	toSerialize["pox_address"] = o.PoxAddress
	toSerialize["stacker_type"] = o.StackerType
	return toSerialize, nil
}

func (o *GetPoxCycleSignerStackers200ResponseResultsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stacker_address",
		"stacked_amount",
		"pox_address",
		"stacker_type",
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

	varGetPoxCycleSignerStackers200ResponseResultsInner := _GetPoxCycleSignerStackers200ResponseResultsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPoxCycleSignerStackers200ResponseResultsInner)

	if err != nil {
		return err
	}

	*o = GetPoxCycleSignerStackers200ResponseResultsInner(varGetPoxCycleSignerStackers200ResponseResultsInner)

	return err
}

type NullableGetPoxCycleSignerStackers200ResponseResultsInner struct {
	value *GetPoxCycleSignerStackers200ResponseResultsInner
	isSet bool
}

func (v NullableGetPoxCycleSignerStackers200ResponseResultsInner) Get() *GetPoxCycleSignerStackers200ResponseResultsInner {
	return v.value
}

func (v *NullableGetPoxCycleSignerStackers200ResponseResultsInner) Set(val *GetPoxCycleSignerStackers200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPoxCycleSignerStackers200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPoxCycleSignerStackers200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPoxCycleSignerStackers200ResponseResultsInner(val *GetPoxCycleSignerStackers200ResponseResultsInner) *NullableGetPoxCycleSignerStackers200ResponseResultsInner {
	return &NullableGetPoxCycleSignerStackers200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableGetPoxCycleSignerStackers200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPoxCycleSignerStackers200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


