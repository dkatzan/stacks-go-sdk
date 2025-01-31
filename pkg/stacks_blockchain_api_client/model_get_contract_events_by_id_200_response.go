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

// checks if the GetContractEventsById200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContractEventsById200Response{}

// GetContractEventsById200Response List of events
type GetContractEventsById200Response struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Results []TokenTransferTransactionEventsInner `json:"results"`
}

type _GetContractEventsById200Response GetContractEventsById200Response

// NewGetContractEventsById200Response instantiates a new GetContractEventsById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContractEventsById200Response(limit int32, offset int32, results []TokenTransferTransactionEventsInner) *GetContractEventsById200Response {
	this := GetContractEventsById200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Results = results
	return &this
}

// NewGetContractEventsById200ResponseWithDefaults instantiates a new GetContractEventsById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContractEventsById200ResponseWithDefaults() *GetContractEventsById200Response {
	this := GetContractEventsById200Response{}
	return &this
}

// GetLimit returns the Limit field value
func (o *GetContractEventsById200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetContractEventsById200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetContractEventsById200Response) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *GetContractEventsById200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *GetContractEventsById200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *GetContractEventsById200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetResults returns the Results field value
func (o *GetContractEventsById200Response) GetResults() []TokenTransferTransactionEventsInner {
	if o == nil {
		var ret []TokenTransferTransactionEventsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetContractEventsById200Response) GetResultsOk() ([]TokenTransferTransactionEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetContractEventsById200Response) SetResults(v []TokenTransferTransactionEventsInner) {
	o.Results = v
}

func (o GetContractEventsById200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContractEventsById200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetContractEventsById200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
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

	varGetContractEventsById200Response := _GetContractEventsById200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetContractEventsById200Response)

	if err != nil {
		return err
	}

	*o = GetContractEventsById200Response(varGetContractEventsById200Response)

	return err
}

type NullableGetContractEventsById200Response struct {
	value *GetContractEventsById200Response
	isSet bool
}

func (v NullableGetContractEventsById200Response) Get() *GetContractEventsById200Response {
	return v.value
}

func (v *NullableGetContractEventsById200Response) Set(val *GetContractEventsById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContractEventsById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContractEventsById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContractEventsById200Response(val *GetContractEventsById200Response) *NullableGetContractEventsById200Response {
	return &NullableGetContractEventsById200Response{value: val, isSet: true}
}

func (v NullableGetContractEventsById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContractEventsById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


