/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofSendPresentationRequestAttribute struct {
	Name         string        `json:"name,omitempty"`
	Value        string        `json:"value,omitempty"`
	Names        []string      `json:"names,omitempty"`
	Restrictions []interface{} `json:"restrictions,omitempty"`
	Revealed     bool          `json:"revealed,omitempty"`
	CredID       string        `json:"cred_id,omitempty"`
}

type PresentProofSendPresentationRequestData struct {
	RequestedAttributes    map[string]PresentProofSendPresentationRequestAttribute `json:"requested_attributes,omitempty"`
	RequestedPredicates    map[string]interface{}                                  `json:"requested_predicates,omitempty"`
	SelfAttestedAttributes map[string]interface{}                                  `json:"self_attested_attributes,omitempty"`
	Comment                string                                                  `json:"comment,omitempty"`
	Name                   string                                                  `json:"name,omitempty"`
	Version                string                                                  `json:"version,omitempty"`
}

// AssertPresentProofSendPresentationRequestDataRequired checks if the required fields are not zero-ed
func AssertPresentProofSendPresentationRequestDataRequired(obj PresentProofSendPresentationRequestData) error {
	return nil
}

// AssertRecursePresentProofSendPresentationRequestDataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofSendPresentationRequestData (e.g. [][]PresentProofSendPresentationRequestData), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofSendPresentationRequestDataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofSendPresentationRequestData, ok := obj.(PresentProofSendPresentationRequestData)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofSendPresentationRequestDataRequired(aPresentProofSendPresentationRequestData)
	})
}
