/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CredentialGetRevocationStatusById200Response struct {

	Revoked bool `json:"revoked"`
}

// AssertCredentialGetRevocationStatusById200ResponseRequired checks if the required fields are not zero-ed
func AssertCredentialGetRevocationStatusById200ResponseRequired(obj CredentialGetRevocationStatusById200Response) error {
	elements := map[string]interface{}{
		"revoked": obj.Revoked,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCredentialGetRevocationStatusById200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CredentialGetRevocationStatusById200Response (e.g. [][]CredentialGetRevocationStatusById200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCredentialGetRevocationStatusById200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCredentialGetRevocationStatusById200Response, ok := obj.(CredentialGetRevocationStatusById200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCredentialGetRevocationStatusById200ResponseRequired(aCredentialGetRevocationStatusById200Response)
	})
}
