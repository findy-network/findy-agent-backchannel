/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofVerifyPresentation200ResponseAllOf struct {

	State *interface{} `json:"state,omitempty"`
}

// AssertPresentProofVerifyPresentation200ResponseAllOfRequired checks if the required fields are not zero-ed
func AssertPresentProofVerifyPresentation200ResponseAllOfRequired(obj PresentProofVerifyPresentation200ResponseAllOf) error {
	return nil
}

// AssertRecursePresentProofVerifyPresentation200ResponseAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofVerifyPresentation200ResponseAllOf (e.g. [][]PresentProofVerifyPresentation200ResponseAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofVerifyPresentation200ResponseAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofVerifyPresentation200ResponseAllOf, ok := obj.(PresentProofVerifyPresentation200ResponseAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofVerifyPresentation200ResponseAllOfRequired(aPresentProofVerifyPresentation200ResponseAllOf)
	})
}
