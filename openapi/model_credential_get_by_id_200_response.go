/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CredentialGetById200Response struct {

	Referent string `json:"referent"`

	SchemaId string `json:"schema_id"`

	CredDefId string `json:"cred_def_id"`

	CredentialId string `json:"credential_id"`

	Credential map[string]interface{} `json:"credential"`
}

// AssertCredentialGetById200ResponseRequired checks if the required fields are not zero-ed
func AssertCredentialGetById200ResponseRequired(obj CredentialGetById200Response) error {
	elements := map[string]interface{}{
		"referent": obj.Referent,
		"schema_id": obj.SchemaId,
		"cred_def_id": obj.CredDefId,
		"credential_id": obj.CredentialId,
		"credential": obj.Credential,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCredentialGetById200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CredentialGetById200Response (e.g. [][]CredentialGetById200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCredentialGetById200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCredentialGetById200Response, ok := obj.(CredentialGetById200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCredentialGetById200ResponseRequired(aCredentialGetById200Response)
	})
}