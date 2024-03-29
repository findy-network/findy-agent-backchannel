/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectionState : All possible connection state values
type ConnectionState string

// List of ConnectionState
// const (
// 	INVITATION ConnectionState = "invitation"
// 	REQUEST ConnectionState = "request"
// 	RESPONSE ConnectionState = "response"
// 	ACTIVE ConnectionState = "active"
// )

const (
	INVITATION ConnectionState = "invited"
	REQUEST    ConnectionState = "requested"
	RESPONSE   ConnectionState = "responded"
	ACTIVE     ConnectionState = "complete"
)

// AssertConnectionStateRequired checks if the required fields are not zero-ed
func AssertConnectionStateRequired(obj ConnectionState) error {
	return nil
}

// AssertRecurseConnectionStateRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ConnectionState (e.g. [][]ConnectionState), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConnectionStateRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConnectionState, ok := obj.(ConnectionState)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConnectionStateRequired(aConnectionState)
	})
}
