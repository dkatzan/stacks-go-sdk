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

// checks if the ListOfBurnchainRewardRecipientsAndAmounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOfBurnchainRewardRecipientsAndAmounts{}

// ListOfBurnchainRewardRecipientsAndAmounts struct for ListOfBurnchainRewardRecipientsAndAmounts
type ListOfBurnchainRewardRecipientsAndAmounts struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total int32 `json:"total"`
	Results []BurnchainRewardSlotHolder `json:"results"`
}

type _ListOfBurnchainRewardRecipientsAndAmounts ListOfBurnchainRewardRecipientsAndAmounts

// NewListOfBurnchainRewardRecipientsAndAmounts instantiates a new ListOfBurnchainRewardRecipientsAndAmounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfBurnchainRewardRecipientsAndAmounts(limit int32, offset int32, total int32, results []BurnchainRewardSlotHolder) *ListOfBurnchainRewardRecipientsAndAmounts {
	this := ListOfBurnchainRewardRecipientsAndAmounts{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewListOfBurnchainRewardRecipientsAndAmountsWithDefaults instantiates a new ListOfBurnchainRewardRecipientsAndAmounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfBurnchainRewardRecipientsAndAmountsWithDefaults() *ListOfBurnchainRewardRecipientsAndAmounts {
	this := ListOfBurnchainRewardRecipientsAndAmounts{}
	return &this
}

// GetLimit returns the Limit field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetResults() []BurnchainRewardSlotHolder {
	if o == nil {
		var ret []BurnchainRewardSlotHolder
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ListOfBurnchainRewardRecipientsAndAmounts) GetResultsOk() ([]BurnchainRewardSlotHolder, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ListOfBurnchainRewardRecipientsAndAmounts) SetResults(v []BurnchainRewardSlotHolder) {
	o.Results = v
}

func (o ListOfBurnchainRewardRecipientsAndAmounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOfBurnchainRewardRecipientsAndAmounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *ListOfBurnchainRewardRecipientsAndAmounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
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

	varListOfBurnchainRewardRecipientsAndAmounts := _ListOfBurnchainRewardRecipientsAndAmounts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOfBurnchainRewardRecipientsAndAmounts)

	if err != nil {
		return err
	}

	*o = ListOfBurnchainRewardRecipientsAndAmounts(varListOfBurnchainRewardRecipientsAndAmounts)

	return err
}

type NullableListOfBurnchainRewardRecipientsAndAmounts struct {
	value *ListOfBurnchainRewardRecipientsAndAmounts
	isSet bool
}

func (v NullableListOfBurnchainRewardRecipientsAndAmounts) Get() *ListOfBurnchainRewardRecipientsAndAmounts {
	return v.value
}

func (v *NullableListOfBurnchainRewardRecipientsAndAmounts) Set(val *ListOfBurnchainRewardRecipientsAndAmounts) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfBurnchainRewardRecipientsAndAmounts) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfBurnchainRewardRecipientsAndAmounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfBurnchainRewardRecipientsAndAmounts(val *ListOfBurnchainRewardRecipientsAndAmounts) *NullableListOfBurnchainRewardRecipientsAndAmounts {
	return &NullableListOfBurnchainRewardRecipientsAndAmounts{value: val, isSet: true}
}

func (v NullableListOfBurnchainRewardRecipientsAndAmounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOfBurnchainRewardRecipientsAndAmounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


