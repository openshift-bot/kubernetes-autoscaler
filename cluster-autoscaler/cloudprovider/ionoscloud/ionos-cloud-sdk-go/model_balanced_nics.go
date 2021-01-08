/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// BalancedNics struct for BalancedNics
type BalancedNics struct {
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	// The type of object that has been created
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
	// Array of items in that collection
	Items *[]Nic `json:"items,omitempty"`
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BalancedNics) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *BalancedNics) SetId(v string) {
	o.Id = &v
}

// HasId returns a boolean if a field has been set.
func (o *BalancedNics) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}



// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *BalancedNics) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *BalancedNics) SetType(v Type) {
	o.Type = &v
}

// HasType returns a boolean if a field has been set.
func (o *BalancedNics) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}



// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BalancedNics) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Href, true
}

// SetHref sets field value
func (o *BalancedNics) SetHref(v string) {
	o.Href = &v
}

// HasHref returns a boolean if a field has been set.
func (o *BalancedNics) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}



// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []Nic will be returned
func (o *BalancedNics) GetItems() *[]Nic {
	if o == nil {
		return nil
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetItemsOk() (*[]Nic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *BalancedNics) SetItems(v []Nic) {
	o.Items = &v
}

// HasItems returns a boolean if a field has been set.
func (o *BalancedNics) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}


func (o BalancedNics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	
	return json.Marshal(toSerialize)
}

type NullableBalancedNics struct {
	value *BalancedNics
	isSet bool
}

func (v NullableBalancedNics) Get() *BalancedNics {
	return v.value
}

func (v *NullableBalancedNics) Set(val *BalancedNics) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancedNics) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancedNics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancedNics(val *BalancedNics) *NullableBalancedNics {
	return &NullableBalancedNics{value: val, isSet: true}
}

func (v NullableBalancedNics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancedNics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


