/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CredentialPreviewAttributesInner struct {

	Name string `json:"name"`

	MimeType string `json:"mime-type,omitempty"`

	Value string `json:"value"`
}

// AssertCredentialPreviewAttributesInnerRequired checks if the required fields are not zero-ed
func AssertCredentialPreviewAttributesInnerRequired(obj CredentialPreviewAttributesInner) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"value": obj.Value,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCredentialPreviewAttributesInnerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CredentialPreviewAttributesInner (e.g. [][]CredentialPreviewAttributesInner), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCredentialPreviewAttributesInnerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCredentialPreviewAttributesInner, ok := obj.(CredentialPreviewAttributesInner)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCredentialPreviewAttributesInnerRequired(aCredentialPreviewAttributesInner)
	})
}
