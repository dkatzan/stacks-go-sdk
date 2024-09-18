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

// checks if the BnsError1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BnsError1{}

// BnsError1 Error
type BnsError1 struct {
	Error string `json:"error"`
}

type _BnsError1 BnsError1

// NewBnsError1 instantiates a new BnsError1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsError1(error_ string) *BnsError1 {
	this := BnsError1{}
	this.Error = error_
	return &this
}

// NewBnsError1WithDefaults instantiates a new BnsError1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsError1WithDefaults() *BnsError1 {
	this := BnsError1{}
	return &this
}

// GetError returns the Error field value
func (o *BnsError1) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *BnsError1) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *BnsError1) SetError(v string) {
	o.Error = v
}

func (o BnsError1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsError1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *BnsError1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varBnsError1 := _BnsError1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsError1)

	if err != nil {
		return err
	}

	*o = BnsError1(varBnsError1)

	return err
}

type NullableBnsError1 struct {
	value *BnsError1
	isSet bool
}

func (v NullableBnsError1) Get() *BnsError1 {
	return v.value
}

func (v *NullableBnsError1) Set(val *BnsError1) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsError1) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsError1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsError1(val *BnsError1) *NullableBnsError1 {
	return &NullableBnsError1{value: val, isSet: true}
}

func (v NullableBnsError1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsError1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


