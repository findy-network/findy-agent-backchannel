/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SendOfferToStartCredentialIssuanceFlowData struct {
	Comment string `json:"comment,omitempty"`

	CredentialPreview CredentialPreview `json:"credential_preview"`

	CredDefId string `json:"cred_def_id,omitempty"`

	ConnectionId string `json:"connection_id"`
}

// AssertSendOfferToStartCredentialIssuanceFlowDataRequired checks if the required fields are not zero-ed
func AssertSendOfferToStartCredentialIssuanceFlowDataRequired(obj SendOfferToStartCredentialIssuanceFlowData) error {
	elements := map[string]interface{}{
		//"credential_preview": obj.CredentialPreview,
		//"connection_id": obj.ConnectionId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCredentialPreviewRequired(obj.CredentialPreview); err != nil {
		return err
	}
	return nil
}

// AssertRecurseSendOfferToStartCredentialIssuanceFlowDataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SendOfferToStartCredentialIssuanceFlowData (e.g. [][]SendOfferToStartCredentialIssuanceFlowData), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSendOfferToStartCredentialIssuanceFlowDataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSendOfferToStartCredentialIssuanceFlowData, ok := obj.(SendOfferToStartCredentialIssuanceFlowData)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSendOfferToStartCredentialIssuanceFlowDataRequired(aSendOfferToStartCredentialIssuanceFlowData)
	})
}
