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

// checks if the StxAssetTransactionEvent1AllOfAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StxAssetTransactionEvent1AllOfAsset{}

// StxAssetTransactionEvent1AllOfAsset struct for StxAssetTransactionEvent1AllOfAsset
type StxAssetTransactionEvent1AllOfAsset struct {
	AssetEventType StxAssetTransactionEvent1AllOfAssetAssetEventType `json:"asset_event_type"`
	Sender string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount string `json:"amount"`
	Memo *string `json:"memo,omitempty"`
}

type _StxAssetTransactionEvent1AllOfAsset StxAssetTransactionEvent1AllOfAsset

// NewStxAssetTransactionEvent1AllOfAsset instantiates a new StxAssetTransactionEvent1AllOfAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStxAssetTransactionEvent1AllOfAsset(assetEventType StxAssetTransactionEvent1AllOfAssetAssetEventType, sender string, recipient string, amount string) *StxAssetTransactionEvent1AllOfAsset {
	this := StxAssetTransactionEvent1AllOfAsset{}
	this.AssetEventType = assetEventType
	this.Sender = sender
	this.Recipient = recipient
	this.Amount = amount
	return &this
}

// NewStxAssetTransactionEvent1AllOfAssetWithDefaults instantiates a new StxAssetTransactionEvent1AllOfAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStxAssetTransactionEvent1AllOfAssetWithDefaults() *StxAssetTransactionEvent1AllOfAsset {
	this := StxAssetTransactionEvent1AllOfAsset{}
	return &this
}

// GetAssetEventType returns the AssetEventType field value
func (o *StxAssetTransactionEvent1AllOfAsset) GetAssetEventType() StxAssetTransactionEvent1AllOfAssetAssetEventType {
	if o == nil {
		var ret StxAssetTransactionEvent1AllOfAssetAssetEventType
		return ret
	}

	return o.AssetEventType
}

// GetAssetEventTypeOk returns a tuple with the AssetEventType field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent1AllOfAsset) GetAssetEventTypeOk() (*StxAssetTransactionEvent1AllOfAssetAssetEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetEventType, true
}

// SetAssetEventType sets field value
func (o *StxAssetTransactionEvent1AllOfAsset) SetAssetEventType(v StxAssetTransactionEvent1AllOfAssetAssetEventType) {
	o.AssetEventType = v
}

// GetSender returns the Sender field value
func (o *StxAssetTransactionEvent1AllOfAsset) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent1AllOfAsset) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *StxAssetTransactionEvent1AllOfAsset) SetSender(v string) {
	o.Sender = v
}

// GetRecipient returns the Recipient field value
func (o *StxAssetTransactionEvent1AllOfAsset) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent1AllOfAsset) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *StxAssetTransactionEvent1AllOfAsset) SetRecipient(v string) {
	o.Recipient = v
}

// GetAmount returns the Amount field value
func (o *StxAssetTransactionEvent1AllOfAsset) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent1AllOfAsset) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *StxAssetTransactionEvent1AllOfAsset) SetAmount(v string) {
	o.Amount = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *StxAssetTransactionEvent1AllOfAsset) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent1AllOfAsset) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *StxAssetTransactionEvent1AllOfAsset) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *StxAssetTransactionEvent1AllOfAsset) SetMemo(v string) {
	o.Memo = &v
}

func (o StxAssetTransactionEvent1AllOfAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StxAssetTransactionEvent1AllOfAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_event_type"] = o.AssetEventType
	toSerialize["sender"] = o.Sender
	toSerialize["recipient"] = o.Recipient
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	return toSerialize, nil
}

func (o *StxAssetTransactionEvent1AllOfAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_event_type",
		"sender",
		"recipient",
		"amount",
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

	varStxAssetTransactionEvent1AllOfAsset := _StxAssetTransactionEvent1AllOfAsset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStxAssetTransactionEvent1AllOfAsset)

	if err != nil {
		return err
	}

	*o = StxAssetTransactionEvent1AllOfAsset(varStxAssetTransactionEvent1AllOfAsset)

	return err
}

type NullableStxAssetTransactionEvent1AllOfAsset struct {
	value *StxAssetTransactionEvent1AllOfAsset
	isSet bool
}

func (v NullableStxAssetTransactionEvent1AllOfAsset) Get() *StxAssetTransactionEvent1AllOfAsset {
	return v.value
}

func (v *NullableStxAssetTransactionEvent1AllOfAsset) Set(val *StxAssetTransactionEvent1AllOfAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableStxAssetTransactionEvent1AllOfAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableStxAssetTransactionEvent1AllOfAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStxAssetTransactionEvent1AllOfAsset(val *StxAssetTransactionEvent1AllOfAsset) *NullableStxAssetTransactionEvent1AllOfAsset {
	return &NullableStxAssetTransactionEvent1AllOfAsset{value: val, isSet: true}
}

func (v NullableStxAssetTransactionEvent1AllOfAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStxAssetTransactionEvent1AllOfAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


