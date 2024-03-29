/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DidExchangeCreateRequestResolvableDidRequestData struct {

	TheirPublicDid string `json:"their_public_did"`

	TheirDid string `json:"their_did"`
}

// AssertDidExchangeCreateRequestResolvableDidRequestDataRequired checks if the required fields are not zero-ed
func AssertDidExchangeCreateRequestResolvableDidRequestDataRequired(obj DidExchangeCreateRequestResolvableDidRequestData) error {
	elements := map[string]interface{}{
		"their_public_did": obj.TheirPublicDid,
		"their_did": obj.TheirDid,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseDidExchangeCreateRequestResolvableDidRequestDataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DidExchangeCreateRequestResolvableDidRequestData (e.g. [][]DidExchangeCreateRequestResolvableDidRequestData), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDidExchangeCreateRequestResolvableDidRequestDataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDidExchangeCreateRequestResolvableDidRequestData, ok := obj.(DidExchangeCreateRequestResolvableDidRequestData)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDidExchangeCreateRequestResolvableDidRequestDataRequired(aDidExchangeCreateRequestResolvableDidRequestData)
	})
}
