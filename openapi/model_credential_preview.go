/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CredentialPreview struct {
	Type string `json:"@type"`

	Attributes []CredentialPreviewAttributesInner `json:"attributes"`
}

// AssertCredentialPreviewRequired checks if the required fields are not zero-ed
func AssertCredentialPreviewRequired(obj CredentialPreview) error {
	elements := map[string]interface{}{
		//"@type": obj.Type,
		//"attributes": obj.Attributes,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Attributes {
		if err := AssertCredentialPreviewAttributesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseCredentialPreviewRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CredentialPreview (e.g. [][]CredentialPreview), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCredentialPreviewRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCredentialPreview, ok := obj.(CredentialPreview)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCredentialPreviewRequired(aCredentialPreview)
	})
}
