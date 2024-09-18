/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v7.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"fmt"
)

// checks if the MempoolTransactionStatsResponseTxSimpleFeeAveragesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MempoolTransactionStatsResponseTxSimpleFeeAveragesValue{}

// MempoolTransactionStatsResponseTxSimpleFeeAveragesValue struct for MempoolTransactionStatsResponseTxSimpleFeeAveragesValue
type MempoolTransactionStatsResponseTxSimpleFeeAveragesValue struct {
	P25 NullableFloat32 `json:"p25"`
	P50 NullableFloat32 `json:"p50"`
	P75 NullableFloat32 `json:"p75"`
	P95 NullableFloat32 `json:"p95"`
	AdditionalProperties map[string]interface{}
}

type _MempoolTransactionStatsResponseTxSimpleFeeAveragesValue MempoolTransactionStatsResponseTxSimpleFeeAveragesValue

// NewMempoolTransactionStatsResponseTxSimpleFeeAveragesValue instantiates a new MempoolTransactionStatsResponseTxSimpleFeeAveragesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponseTxSimpleFeeAveragesValue(p25 NullableFloat32, p50 NullableFloat32, p75 NullableFloat32, p95 NullableFloat32) *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue {
	this := MempoolTransactionStatsResponseTxSimpleFeeAveragesValue{}
	this.P25 = p25
	this.P50 = p50
	this.P75 = p75
	this.P95 = p95
	return &this
}

// NewMempoolTransactionStatsResponseTxSimpleFeeAveragesValueWithDefaults instantiates a new MempoolTransactionStatsResponseTxSimpleFeeAveragesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseTxSimpleFeeAveragesValueWithDefaults() *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue {
	this := MempoolTransactionStatsResponseTxSimpleFeeAveragesValue{}
	return &this
}

// GetP25 returns the P25 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP25() float32 {
	if o == nil || o.P25.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P25.Get()
}

// GetP25Ok returns a tuple with the P25 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP25Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P25.Get(), o.P25.IsSet()
}

// SetP25 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) SetP25(v float32) {
	o.P25.Set(&v)
}

// GetP50 returns the P50 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP50() float32 {
	if o == nil || o.P50.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P50.Get()
}

// GetP50Ok returns a tuple with the P50 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP50Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P50.Get(), o.P50.IsSet()
}

// SetP50 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) SetP50(v float32) {
	o.P50.Set(&v)
}

// GetP75 returns the P75 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP75() float32 {
	if o == nil || o.P75.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P75.Get()
}

// GetP75Ok returns a tuple with the P75 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP75Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P75.Get(), o.P75.IsSet()
}

// SetP75 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) SetP75(v float32) {
	o.P75.Set(&v)
}

// GetP95 returns the P95 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP95() float32 {
	if o == nil || o.P95.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P95.Get()
}

// GetP95Ok returns a tuple with the P95 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) GetP95Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P95.Get(), o.P95.IsSet()
}

// SetP95 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) SetP95(v float32) {
	o.P95.Set(&v)
}

func (o MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["p25"] = o.P25.Get()
	toSerialize["p50"] = o.P50.Get()
	toSerialize["p75"] = o.P75.Get()
	toSerialize["p95"] = o.P95.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"p25",
		"p50",
		"p75",
		"p95",
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

	varMempoolTransactionStatsResponseTxSimpleFeeAveragesValue := _MempoolTransactionStatsResponseTxSimpleFeeAveragesValue{}

	err = json.Unmarshal(data, &varMempoolTransactionStatsResponseTxSimpleFeeAveragesValue)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponseTxSimpleFeeAveragesValue(varMempoolTransactionStatsResponseTxSimpleFeeAveragesValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "p25")
		delete(additionalProperties, "p50")
		delete(additionalProperties, "p75")
		delete(additionalProperties, "p95")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue struct {
	value *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue
	isSet bool
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue) Get() *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue) Set(val *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue(val *MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue {
	return &NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


