/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofV3SendRequestRequestData struct {

	PresentationRequest PresentProofV3SendRequestRequestDataPresentationRequest `json:"presentation_request,omitempty"`
}

// AssertPresentProofV3SendRequestRequestDataRequired checks if the required fields are not zero-ed
func AssertPresentProofV3SendRequestRequestDataRequired(obj PresentProofV3SendRequestRequestData) error {
	if err := AssertPresentProofV3SendRequestRequestDataPresentationRequestRequired(obj.PresentationRequest); err != nil {
		return err
	}
	return nil
}

// AssertRecursePresentProofV3SendRequestRequestDataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofV3SendRequestRequestData (e.g. [][]PresentProofV3SendRequestRequestData), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofV3SendRequestRequestDataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofV3SendRequestRequestData, ok := obj.(PresentProofV3SendRequestRequestData)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofV3SendRequestRequestDataRequired(aPresentProofV3SendRequestRequestData)
	})
}
