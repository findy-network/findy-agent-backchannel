/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AgentStartRequest struct {

	Data AgentStartRequestData `json:"data"`
}

// AssertAgentStartRequestRequired checks if the required fields are not zero-ed
func AssertAgentStartRequestRequired(obj AgentStartRequest) error {
	elements := map[string]interface{}{
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAgentStartRequestDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRecurseAgentStartRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AgentStartRequest (e.g. [][]AgentStartRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAgentStartRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAgentStartRequest, ok := obj.(AgentStartRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAgentStartRequestRequired(aAgentStartRequest)
	})
}
