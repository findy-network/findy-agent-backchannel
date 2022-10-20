/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofSendPresentation200ResponseAllOf struct {

	State *interface{} `json:"state,omitempty"`
}

// AssertPresentProofSendPresentation200ResponseAllOfRequired checks if the required fields are not zero-ed
func AssertPresentProofSendPresentation200ResponseAllOfRequired(obj PresentProofSendPresentation200ResponseAllOf) error {
	return nil
}

// AssertRecursePresentProofSendPresentation200ResponseAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofSendPresentation200ResponseAllOf (e.g. [][]PresentProofSendPresentation200ResponseAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofSendPresentation200ResponseAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofSendPresentation200ResponseAllOf, ok := obj.(PresentProofSendPresentation200ResponseAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofSendPresentation200ResponseAllOfRequired(aPresentProofSendPresentation200ResponseAllOf)
	})
}
