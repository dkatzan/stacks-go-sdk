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

// checks if the GetBurnchainRewardList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBurnchainRewardList200Response{}

// GetBurnchainRewardList200Response List of burnchain reward recipients and amounts
type GetBurnchainRewardList200Response struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Results []BurnchainReward `json:"results"`
}

type _GetBurnchainRewardList200Response GetBurnchainRewardList200Response

// NewGetBurnchainRewardList200Response instantiates a new GetBurnchainRewardList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBurnchainRewardList200Response(limit int32, offset int32, results []BurnchainReward) *GetBurnchainRewardList200Response {
	this := GetBurnchainRewardList200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Results = results
	return &this
}

// NewGetBurnchainRewardList200ResponseWithDefaults instantiates a new GetBurnchainRewardList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBurnchainRewardList200ResponseWithDefaults() *GetBurnchainRewardList200Response {
	this := GetBurnchainRewardList200Response{}
	return &this
}

// GetLimit returns the Limit field value
func (o *GetBurnchainRewardList200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetBurnchainRewardList200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetBurnchainRewardList200Response) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *GetBurnchainRewardList200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *GetBurnchainRewardList200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *GetBurnchainRewardList200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetResults returns the Results field value
func (o *GetBurnchainRewardList200Response) GetResults() []BurnchainReward {
	if o == nil {
		var ret []BurnchainReward
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetBurnchainRewardList200Response) GetResultsOk() ([]BurnchainReward, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetBurnchainRewardList200Response) SetResults(v []BurnchainReward) {
	o.Results = v
}

func (o GetBurnchainRewardList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBurnchainRewardList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetBurnchainRewardList200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetBurnchainRewardList200Response := _GetBurnchainRewardList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBurnchainRewardList200Response)

	if err != nil {
		return err
	}

	*o = GetBurnchainRewardList200Response(varGetBurnchainRewardList200Response)

	return err
}

type NullableGetBurnchainRewardList200Response struct {
	value *GetBurnchainRewardList200Response
	isSet bool
}

func (v NullableGetBurnchainRewardList200Response) Get() *GetBurnchainRewardList200Response {
	return v.value
}

func (v *NullableGetBurnchainRewardList200Response) Set(val *GetBurnchainRewardList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBurnchainRewardList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBurnchainRewardList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBurnchainRewardList200Response(val *GetBurnchainRewardList200Response) *NullableGetBurnchainRewardList200Response {
	return &NullableGetBurnchainRewardList200Response{value: val, isSet: true}
}

func (v NullableGetBurnchainRewardList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBurnchainRewardList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


