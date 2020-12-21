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

// KubernetesNodePoolAnnotation map of annotations attached to node pool
type KubernetesNodePoolAnnotation struct {
	// Key of the annotation. String part must consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character.
	Key *string `json:"key,omitempty"`
	// Value of the annotation.
	Value *string `json:"value,omitempty"`
}



// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolAnnotation) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolAnnotation) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key, true
}

// SetKey sets field value
func (o *KubernetesNodePoolAnnotation) SetKey(v string) {
	o.Key = &v
}

// HasKey returns a boolean if a field has been set.
func (o *KubernetesNodePoolAnnotation) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}



// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolAnnotation) GetValue() *string {
	if o == nil {
		return nil
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolAnnotation) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *KubernetesNodePoolAnnotation) SetValue(v string) {
	o.Value = &v
}

// HasValue returns a boolean if a field has been set.
func (o *KubernetesNodePoolAnnotation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}


func (o KubernetesNodePoolAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	

	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodePoolAnnotation struct {
	value *KubernetesNodePoolAnnotation
	isSet bool
}

func (v NullableKubernetesNodePoolAnnotation) Get() *KubernetesNodePoolAnnotation {
	return v.value
}

func (v *NullableKubernetesNodePoolAnnotation) Set(val *KubernetesNodePoolAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodePoolAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodePoolAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodePoolAnnotation(val *KubernetesNodePoolAnnotation) *NullableKubernetesNodePoolAnnotation {
	return &NullableKubernetesNodePoolAnnotation{value: val, isSet: true}
}

func (v NullableKubernetesNodePoolAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodePoolAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


