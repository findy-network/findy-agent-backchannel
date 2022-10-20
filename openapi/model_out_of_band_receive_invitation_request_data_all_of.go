/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OutOfBandReceiveInvitationRequestDataAllOf struct {

	UseExistingConnection bool `json:"use_existing_connection"`

	MediatorConnectionId string `json:"mediator_connection_id,omitempty"`
}

// AssertOutOfBandReceiveInvitationRequestDataAllOfRequired checks if the required fields are not zero-ed
func AssertOutOfBandReceiveInvitationRequestDataAllOfRequired(obj OutOfBandReceiveInvitationRequestDataAllOf) error {
	elements := map[string]interface{}{
		"use_existing_connection": obj.UseExistingConnection,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseOutOfBandReceiveInvitationRequestDataAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of OutOfBandReceiveInvitationRequestDataAllOf (e.g. [][]OutOfBandReceiveInvitationRequestDataAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseOutOfBandReceiveInvitationRequestDataAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aOutOfBandReceiveInvitationRequestDataAllOf, ok := obj.(OutOfBandReceiveInvitationRequestDataAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertOutOfBandReceiveInvitationRequestDataAllOfRequired(aOutOfBandReceiveInvitationRequestDataAllOf)
	})
}
