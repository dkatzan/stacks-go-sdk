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

// checks if the AddressStxBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressStxBalance{}

// AddressStxBalance GET request that returns address balances
type AddressStxBalance struct {
	Balance string `json:"balance"`
	// Total STX balance considering pending mempool transactions
	EstimatedBalance *string `json:"estimated_balance,omitempty"`
	TotalSent string `json:"total_sent"`
	TotalReceived string `json:"total_received"`
	TotalFeesSent string `json:"total_fees_sent"`
	TotalMinerRewardsReceived string `json:"total_miner_rewards_received"`
	// The transaction where the lock event occurred. Empty if no tokens are locked.
	LockTxId string `json:"lock_tx_id"`
	// The amount of locked STX, as string quoted micro-STX. Zero if no tokens are locked.
	Locked string `json:"locked"`
	// The STX chain block height of when the lock event occurred. Zero if no tokens are locked.
	LockHeight int32 `json:"lock_height"`
	// The burnchain block height of when the lock event occurred. Zero if no tokens are locked.
	BurnchainLockHeight int32 `json:"burnchain_lock_height"`
	// The burnchain block height of when the tokens unlock. Zero if no tokens are locked.
	BurnchainUnlockHeight int32 `json:"burnchain_unlock_height"`
	TokenOfferingLocked *AddressTokenOfferingLocked `json:"token_offering_locked,omitempty"`
}

type _AddressStxBalance AddressStxBalance

// NewAddressStxBalance instantiates a new AddressStxBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressStxBalance(balance string, totalSent string, totalReceived string, totalFeesSent string, totalMinerRewardsReceived string, lockTxId string, locked string, lockHeight int32, burnchainLockHeight int32, burnchainUnlockHeight int32) *AddressStxBalance {
	this := AddressStxBalance{}
	this.Balance = balance
	this.TotalSent = totalSent
	this.TotalReceived = totalReceived
	this.TotalFeesSent = totalFeesSent
	this.TotalMinerRewardsReceived = totalMinerRewardsReceived
	this.LockTxId = lockTxId
	this.Locked = locked
	this.LockHeight = lockHeight
	this.BurnchainLockHeight = burnchainLockHeight
	this.BurnchainUnlockHeight = burnchainUnlockHeight
	return &this
}

// NewAddressStxBalanceWithDefaults instantiates a new AddressStxBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressStxBalanceWithDefaults() *AddressStxBalance {
	this := AddressStxBalance{}
	return &this
}

// GetBalance returns the Balance field value
func (o *AddressStxBalance) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *AddressStxBalance) SetBalance(v string) {
	o.Balance = v
}

// GetEstimatedBalance returns the EstimatedBalance field value if set, zero value otherwise.
func (o *AddressStxBalance) GetEstimatedBalance() string {
	if o == nil || IsNil(o.EstimatedBalance) {
		var ret string
		return ret
	}
	return *o.EstimatedBalance
}

// GetEstimatedBalanceOk returns a tuple with the EstimatedBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetEstimatedBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedBalance) {
		return nil, false
	}
	return o.EstimatedBalance, true
}

// HasEstimatedBalance returns a boolean if a field has been set.
func (o *AddressStxBalance) HasEstimatedBalance() bool {
	if o != nil && !IsNil(o.EstimatedBalance) {
		return true
	}

	return false
}

// SetEstimatedBalance gets a reference to the given string and assigns it to the EstimatedBalance field.
func (o *AddressStxBalance) SetEstimatedBalance(v string) {
	o.EstimatedBalance = &v
}

// GetTotalSent returns the TotalSent field value
func (o *AddressStxBalance) GetTotalSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalSent
}

// GetTotalSentOk returns a tuple with the TotalSent field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetTotalSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSent, true
}

// SetTotalSent sets field value
func (o *AddressStxBalance) SetTotalSent(v string) {
	o.TotalSent = v
}

// GetTotalReceived returns the TotalReceived field value
func (o *AddressStxBalance) GetTotalReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetTotalReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalReceived, true
}

// SetTotalReceived sets field value
func (o *AddressStxBalance) SetTotalReceived(v string) {
	o.TotalReceived = v
}

// GetTotalFeesSent returns the TotalFeesSent field value
func (o *AddressStxBalance) GetTotalFeesSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalFeesSent
}

// GetTotalFeesSentOk returns a tuple with the TotalFeesSent field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetTotalFeesSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFeesSent, true
}

// SetTotalFeesSent sets field value
func (o *AddressStxBalance) SetTotalFeesSent(v string) {
	o.TotalFeesSent = v
}

// GetTotalMinerRewardsReceived returns the TotalMinerRewardsReceived field value
func (o *AddressStxBalance) GetTotalMinerRewardsReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalMinerRewardsReceived
}

// GetTotalMinerRewardsReceivedOk returns a tuple with the TotalMinerRewardsReceived field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetTotalMinerRewardsReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMinerRewardsReceived, true
}

// SetTotalMinerRewardsReceived sets field value
func (o *AddressStxBalance) SetTotalMinerRewardsReceived(v string) {
	o.TotalMinerRewardsReceived = v
}

// GetLockTxId returns the LockTxId field value
func (o *AddressStxBalance) GetLockTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockTxId
}

// GetLockTxIdOk returns a tuple with the LockTxId field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetLockTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockTxId, true
}

// SetLockTxId sets field value
func (o *AddressStxBalance) SetLockTxId(v string) {
	o.LockTxId = v
}

// GetLocked returns the Locked field value
func (o *AddressStxBalance) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *AddressStxBalance) SetLocked(v string) {
	o.Locked = v
}

// GetLockHeight returns the LockHeight field value
func (o *AddressStxBalance) GetLockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockHeight
}

// GetLockHeightOk returns a tuple with the LockHeight field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockHeight, true
}

// SetLockHeight sets field value
func (o *AddressStxBalance) SetLockHeight(v int32) {
	o.LockHeight = v
}

// GetBurnchainLockHeight returns the BurnchainLockHeight field value
func (o *AddressStxBalance) GetBurnchainLockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnchainLockHeight
}

// GetBurnchainLockHeightOk returns a tuple with the BurnchainLockHeight field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetBurnchainLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnchainLockHeight, true
}

// SetBurnchainLockHeight sets field value
func (o *AddressStxBalance) SetBurnchainLockHeight(v int32) {
	o.BurnchainLockHeight = v
}

// GetBurnchainUnlockHeight returns the BurnchainUnlockHeight field value
func (o *AddressStxBalance) GetBurnchainUnlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnchainUnlockHeight
}

// GetBurnchainUnlockHeightOk returns a tuple with the BurnchainUnlockHeight field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetBurnchainUnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnchainUnlockHeight, true
}

// SetBurnchainUnlockHeight sets field value
func (o *AddressStxBalance) SetBurnchainUnlockHeight(v int32) {
	o.BurnchainUnlockHeight = v
}

// GetTokenOfferingLocked returns the TokenOfferingLocked field value if set, zero value otherwise.
func (o *AddressStxBalance) GetTokenOfferingLocked() AddressTokenOfferingLocked {
	if o == nil || IsNil(o.TokenOfferingLocked) {
		var ret AddressTokenOfferingLocked
		return ret
	}
	return *o.TokenOfferingLocked
}

// GetTokenOfferingLockedOk returns a tuple with the TokenOfferingLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressStxBalance) GetTokenOfferingLockedOk() (*AddressTokenOfferingLocked, bool) {
	if o == nil || IsNil(o.TokenOfferingLocked) {
		return nil, false
	}
	return o.TokenOfferingLocked, true
}

// HasTokenOfferingLocked returns a boolean if a field has been set.
func (o *AddressStxBalance) HasTokenOfferingLocked() bool {
	if o != nil && !IsNil(o.TokenOfferingLocked) {
		return true
	}

	return false
}

// SetTokenOfferingLocked gets a reference to the given AddressTokenOfferingLocked and assigns it to the TokenOfferingLocked field.
func (o *AddressStxBalance) SetTokenOfferingLocked(v AddressTokenOfferingLocked) {
	o.TokenOfferingLocked = &v
}

func (o AddressStxBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressStxBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	if !IsNil(o.EstimatedBalance) {
		toSerialize["estimated_balance"] = o.EstimatedBalance
	}
	toSerialize["total_sent"] = o.TotalSent
	toSerialize["total_received"] = o.TotalReceived
	toSerialize["total_fees_sent"] = o.TotalFeesSent
	toSerialize["total_miner_rewards_received"] = o.TotalMinerRewardsReceived
	toSerialize["lock_tx_id"] = o.LockTxId
	toSerialize["locked"] = o.Locked
	toSerialize["lock_height"] = o.LockHeight
	toSerialize["burnchain_lock_height"] = o.BurnchainLockHeight
	toSerialize["burnchain_unlock_height"] = o.BurnchainUnlockHeight
	if !IsNil(o.TokenOfferingLocked) {
		toSerialize["token_offering_locked"] = o.TokenOfferingLocked
	}
	return toSerialize, nil
}

func (o *AddressStxBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance",
		"total_sent",
		"total_received",
		"total_fees_sent",
		"total_miner_rewards_received",
		"lock_tx_id",
		"locked",
		"lock_height",
		"burnchain_lock_height",
		"burnchain_unlock_height",
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

	varAddressStxBalance := _AddressStxBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressStxBalance)

	if err != nil {
		return err
	}

	*o = AddressStxBalance(varAddressStxBalance)

	return err
}

type NullableAddressStxBalance struct {
	value *AddressStxBalance
	isSet bool
}

func (v NullableAddressStxBalance) Get() *AddressStxBalance {
	return v.value
}

func (v *NullableAddressStxBalance) Set(val *AddressStxBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressStxBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressStxBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressStxBalance(val *AddressStxBalance) *NullableAddressStxBalance {
	return &NullableAddressStxBalance{value: val, isSet: true}
}

func (v NullableAddressStxBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressStxBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


