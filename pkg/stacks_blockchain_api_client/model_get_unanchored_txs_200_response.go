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

// checks if the GetUnanchoredTxs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUnanchoredTxs200Response{}

// GetUnanchoredTxs200Response struct for GetUnanchoredTxs200Response
type GetUnanchoredTxs200Response struct {
	Total int32 `json:"total"`
	Results []GetTransactionList200ResponseResultsInner `json:"results"`
}

type _GetUnanchoredTxs200Response GetUnanchoredTxs200Response

// NewGetUnanchoredTxs200Response instantiates a new GetUnanchoredTxs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUnanchoredTxs200Response(total int32, results []GetTransactionList200ResponseResultsInner) *GetUnanchoredTxs200Response {
	this := GetUnanchoredTxs200Response{}
	this.Total = total
	this.Results = results
	return &this
}

// NewGetUnanchoredTxs200ResponseWithDefaults instantiates a new GetUnanchoredTxs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUnanchoredTxs200ResponseWithDefaults() *GetUnanchoredTxs200Response {
	this := GetUnanchoredTxs200Response{}
	return &this
}

// GetTotal returns the Total field value
func (o *GetUnanchoredTxs200Response) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetUnanchoredTxs200Response) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetUnanchoredTxs200Response) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *GetUnanchoredTxs200Response) GetResults() []GetTransactionList200ResponseResultsInner {
	if o == nil {
		var ret []GetTransactionList200ResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetUnanchoredTxs200Response) GetResultsOk() ([]GetTransactionList200ResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetUnanchoredTxs200Response) SetResults(v []GetTransactionList200ResponseResultsInner) {
	o.Results = v
}

func (o GetUnanchoredTxs200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUnanchoredTxs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetUnanchoredTxs200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"results",
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

	varGetUnanchoredTxs200Response := _GetUnanchoredTxs200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetUnanchoredTxs200Response)

	if err != nil {
		return err
	}

	*o = GetUnanchoredTxs200Response(varGetUnanchoredTxs200Response)

	return err
}

type NullableGetUnanchoredTxs200Response struct {
	value *GetUnanchoredTxs200Response
	isSet bool
}

func (v NullableGetUnanchoredTxs200Response) Get() *GetUnanchoredTxs200Response {
	return v.value
}

func (v *NullableGetUnanchoredTxs200Response) Set(val *GetUnanchoredTxs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUnanchoredTxs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUnanchoredTxs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUnanchoredTxs200Response(val *GetUnanchoredTxs200Response) *NullableGetUnanchoredTxs200Response {
	return &NullableGetUnanchoredTxs200Response{value: val, isSet: true}
}

func (v NullableGetUnanchoredTxs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUnanchoredTxs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


