/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofV2SendRequestRequest struct {

	Data SendDifProofRequestData `json:"data"`
}

// AssertPresentProofV2SendRequestRequestRequired checks if the required fields are not zero-ed
func AssertPresentProofV2SendRequestRequestRequired(obj PresentProofV2SendRequestRequest) error {
	elements := map[string]interface{}{
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSendDifProofRequestDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRecursePresentProofV2SendRequestRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofV2SendRequestRequest (e.g. [][]PresentProofV2SendRequestRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofV2SendRequestRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofV2SendRequestRequest, ok := obj.(PresentProofV2SendRequestRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofV2SendRequestRequestRequired(aPresentProofV2SendRequestRequest)
	})
}
