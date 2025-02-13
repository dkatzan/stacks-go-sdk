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

// checks if the AddressSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressSearchResult{}

// AddressSearchResult Address search result
type AddressSearchResult struct {
	// The id used to search this query.
	EntityId string `json:"entity_id"`
	EntityType string `json:"entity_type"`
	Metadata *AddressStxBalance `json:"metadata,omitempty"`
}

type _AddressSearchResult AddressSearchResult

// NewAddressSearchResult instantiates a new AddressSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressSearchResult(entityId string, entityType string) *AddressSearchResult {
	this := AddressSearchResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	return &this
}

// NewAddressSearchResultWithDefaults instantiates a new AddressSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressSearchResultWithDefaults() *AddressSearchResult {
	this := AddressSearchResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *AddressSearchResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *AddressSearchResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *AddressSearchResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *AddressSearchResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AddressSearchResult) GetMetadata() AddressStxBalance {
	if o == nil || IsNil(o.Metadata) {
		var ret AddressStxBalance
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchResult) GetMetadataOk() (*AddressStxBalance, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AddressSearchResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given AddressStxBalance and assigns it to the Metadata field.
func (o *AddressSearchResult) SetMetadata(v AddressStxBalance) {
	o.Metadata = &v
}

func (o AddressSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *AddressSearchResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
		"entity_type",
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

	varAddressSearchResult := _AddressSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressSearchResult)

	if err != nil {
		return err
	}

	*o = AddressSearchResult(varAddressSearchResult)

	return err
}

type NullableAddressSearchResult struct {
	value *AddressSearchResult
	isSet bool
}

func (v NullableAddressSearchResult) Get() *AddressSearchResult {
	return v.value
}

func (v *NullableAddressSearchResult) Set(val *AddressSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressSearchResult(val *AddressSearchResult) *NullableAddressSearchResult {
	return &NullableAddressSearchResult{value: val, isSet: true}
}

func (v NullableAddressSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


