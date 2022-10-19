/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ConnectionAcceptRequest200Response struct {

	ConnectionId string `json:"connection_id,omitempty"`

	State ConnectionState `json:"state,omitempty"`
}

// AssertConnectionAcceptRequest200ResponseRequired checks if the required fields are not zero-ed
func AssertConnectionAcceptRequest200ResponseRequired(obj ConnectionAcceptRequest200Response) error {
	if err := AssertConnectionStateRequired(obj.State); err != nil {
		return err
	}
	return nil
}

// AssertRecurseConnectionAcceptRequest200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ConnectionAcceptRequest200Response (e.g. [][]ConnectionAcceptRequest200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConnectionAcceptRequest200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConnectionAcceptRequest200Response, ok := obj.(ConnectionAcceptRequest200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConnectionAcceptRequest200ResponseRequired(aConnectionAcceptRequest200Response)
	})
}
