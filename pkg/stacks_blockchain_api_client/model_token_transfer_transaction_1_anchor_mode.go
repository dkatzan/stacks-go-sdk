/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v8.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"fmt"
)


// TokenTransferTransaction1AnchorMode `on_chain_only`: the transaction MUST be included in an anchored block, `off_chain_only`: the transaction MUST be included in a microblock, `any`: the leader can choose where to include the transaction.
type TokenTransferTransaction1AnchorMode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TokenTransferTransaction1AnchorMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TokenTransferTransaction1AnchorMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TokenTransferTransaction1AnchorMode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableTokenTransferTransaction1AnchorMode struct {
	value *TokenTransferTransaction1AnchorMode
	isSet bool
}

func (v NullableTokenTransferTransaction1AnchorMode) Get() *TokenTransferTransaction1AnchorMode {
	return v.value
}

func (v *NullableTokenTransferTransaction1AnchorMode) Set(val *TokenTransferTransaction1AnchorMode) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransaction1AnchorMode) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransaction1AnchorMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransaction1AnchorMode(val *TokenTransferTransaction1AnchorMode) *NullableTokenTransferTransaction1AnchorMode {
	return &NullableTokenTransferTransaction1AnchorMode{value: val, isSet: true}
}

func (v NullableTokenTransferTransaction1AnchorMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransaction1AnchorMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


