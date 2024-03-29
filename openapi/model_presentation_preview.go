/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentationPreview struct {
	Type string `json:"@type"`

	Comment string `json:"comment"`

	Attributes []PresentationPreviewAttributesInner `json:"attributes"`

	Predicates []PresentationPreviewPredicatesInner `json:"predicates"`
}

// AssertPresentationPreviewRequired checks if the required fields are not zero-ed
func AssertPresentationPreviewRequired(obj PresentationPreview) error {
	elements := map[string]interface{}{
		//"@type":      obj.Type,
		"attributes": obj.Attributes,
		"predicates": obj.Predicates,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Attributes {
		if err := AssertPresentationPreviewAttributesInnerRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Predicates {
		if err := AssertPresentationPreviewPredicatesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecursePresentationPreviewRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentationPreview (e.g. [][]PresentationPreview), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentationPreviewRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentationPreview, ok := obj.(PresentationPreview)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentationPreviewRequired(aPresentationPreview)
	})
}
