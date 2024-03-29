/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofSendRequest200Response struct {

	State *interface{} `json:"state"`

	ThreadId string `json:"thread_id"`
}

// AssertPresentProofSendRequest200ResponseRequired checks if the required fields are not zero-ed
func AssertPresentProofSendRequest200ResponseRequired(obj PresentProofSendRequest200Response) error {
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

// AssertRecursePresentProofSendRequest200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofSendRequest200Response (e.g. [][]PresentProofSendRequest200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofSendRequest200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofSendRequest200Response, ok := obj.(PresentProofSendRequest200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofSendRequest200ResponseRequired(aPresentProofSendRequest200Response)
	})
}
