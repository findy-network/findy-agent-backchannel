/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ConnectionCreateInvitationRequest struct {

	Data ConnectionCreateInvitationRequestData `json:"data"`
}

// AssertConnectionCreateInvitationRequestRequired checks if the required fields are not zero-ed
func AssertConnectionCreateInvitationRequestRequired(obj ConnectionCreateInvitationRequest) error {
	elements := map[string]interface{}{
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertConnectionCreateInvitationRequestDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRecurseConnectionCreateInvitationRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ConnectionCreateInvitationRequest (e.g. [][]ConnectionCreateInvitationRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConnectionCreateInvitationRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConnectionCreateInvitationRequest, ok := obj.(ConnectionCreateInvitationRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConnectionCreateInvitationRequestRequired(aConnectionCreateInvitationRequest)
	})
}
