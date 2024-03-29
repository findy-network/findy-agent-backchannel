/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofVerifyPresentation200Response struct {

	State *interface{} `json:"state"`

	ThreadId string `json:"thread_id"`
}

// AssertPresentProofVerifyPresentation200ResponseRequired checks if the required fields are not zero-ed
func AssertPresentProofVerifyPresentation200ResponseRequired(obj PresentProofVerifyPresentation200Response) error {
	elements := map[string]interface{}{
		"state": obj.State,
		"thread_id": obj.ThreadId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecursePresentProofVerifyPresentation200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofVerifyPresentation200Response (e.g. [][]PresentProofVerifyPresentation200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofVerifyPresentation200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofVerifyPresentation200Response, ok := obj.(PresentProofVerifyPresentation200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofVerifyPresentation200ResponseRequired(aPresentProofVerifyPresentation200Response)
	})
}
