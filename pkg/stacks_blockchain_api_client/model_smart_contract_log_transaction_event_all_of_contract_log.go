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

// checks if the SmartContractLogTransactionEventAllOfContractLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartContractLogTransactionEventAllOfContractLog{}

// SmartContractLogTransactionEventAllOfContractLog struct for SmartContractLogTransactionEventAllOfContractLog
type SmartContractLogTransactionEventAllOfContractLog struct {
	ContractId string `json:"contract_id"`
	Topic string `json:"topic"`
	Value TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue `json:"value"`
}

type _SmartContractLogTransactionEventAllOfContractLog SmartContractLogTransactionEventAllOfContractLog

// NewSmartContractLogTransactionEventAllOfContractLog instantiates a new SmartContractLogTransactionEventAllOfContractLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractLogTransactionEventAllOfContractLog(contractId string, topic string, value TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue) *SmartContractLogTransactionEventAllOfContractLog {
	this := SmartContractLogTransactionEventAllOfContractLog{}
	this.ContractId = contractId
	this.Topic = topic
	this.Value = value
	return &this
}

// NewSmartContractLogTransactionEventAllOfContractLogWithDefaults instantiates a new SmartContractLogTransactionEventAllOfContractLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractLogTransactionEventAllOfContractLogWithDefaults() *SmartContractLogTransactionEventAllOfContractLog {
	this := SmartContractLogTransactionEventAllOfContractLog{}
	return &this
}

// GetContractId returns the ContractId field value
func (o *SmartContractLogTransactionEventAllOfContractLog) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *SmartContractLogTransactionEventAllOfContractLog) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *SmartContractLogTransactionEventAllOfContractLog) SetContractId(v string) {
	o.ContractId = v
}

// GetTopic returns the Topic field value
func (o *SmartContractLogTransactionEventAllOfContractLog) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *SmartContractLogTransactionEventAllOfContractLog) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *SmartContractLogTransactionEventAllOfContractLog) SetTopic(v string) {
	o.Topic = v
}

// GetValue returns the Value field value
func (o *SmartContractLogTransactionEventAllOfContractLog) GetValue() TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SmartContractLogTransactionEventAllOfContractLog) GetValueOk() (*TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SmartContractLogTransactionEventAllOfContractLog) SetValue(v TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue) {
	o.Value = v
}

func (o SmartContractLogTransactionEventAllOfContractLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractLogTransactionEventAllOfContractLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract_id"] = o.ContractId
	toSerialize["topic"] = o.Topic
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *SmartContractLogTransactionEventAllOfContractLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract_id",
		"topic",
		"value",
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

	varSmartContractLogTransactionEventAllOfContractLog := _SmartContractLogTransactionEventAllOfContractLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractLogTransactionEventAllOfContractLog)

	if err != nil {
		return err
	}

	*o = SmartContractLogTransactionEventAllOfContractLog(varSmartContractLogTransactionEventAllOfContractLog)

	return err
}

type NullableSmartContractLogTransactionEventAllOfContractLog struct {
	value *SmartContractLogTransactionEventAllOfContractLog
	isSet bool
}

func (v NullableSmartContractLogTransactionEventAllOfContractLog) Get() *SmartContractLogTransactionEventAllOfContractLog {
	return v.value
}

func (v *NullableSmartContractLogTransactionEventAllOfContractLog) Set(val *SmartContractLogTransactionEventAllOfContractLog) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractLogTransactionEventAllOfContractLog) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractLogTransactionEventAllOfContractLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractLogTransactionEventAllOfContractLog(val *SmartContractLogTransactionEventAllOfContractLog) *NullableSmartContractLogTransactionEventAllOfContractLog {
	return &NullableSmartContractLogTransactionEventAllOfContractLog{value: val, isSet: true}
}

func (v NullableSmartContractLogTransactionEventAllOfContractLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractLogTransactionEventAllOfContractLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


