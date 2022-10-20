/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofV2VerifyPresentation200ResponseAllOf struct {

	State *interface{} `json:"state"`

	Verified bool `json:"verified"`
}

// AssertPresentProofV2VerifyPresentation200ResponseAllOfRequired checks if the required fields are not zero-ed
func AssertPresentProofV2VerifyPresentation200ResponseAllOfRequired(obj PresentProofV2VerifyPresentation200ResponseAllOf) error {
	elements := map[string]interface{}{
		"state": obj.State,
		"verified": obj.Verified,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecursePresentProofV2VerifyPresentation200ResponseAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofV2VerifyPresentation200ResponseAllOf (e.g. [][]PresentProofV2VerifyPresentation200ResponseAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofV2VerifyPresentation200ResponseAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofV2VerifyPresentation200ResponseAllOf, ok := obj.(PresentProofV2VerifyPresentation200ResponseAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofV2VerifyPresentation200ResponseAllOfRequired(aPresentProofV2VerifyPresentation200ResponseAllOf)
	})
}