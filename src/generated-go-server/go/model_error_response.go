// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Tender Management API
 *
 * API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка).
 *
 * API version: 1.0
 */

package openapi

// ErrorResponse - Используется для возвращения ошибки пользователю
type ErrorResponse struct {

	// Описание ошибки в свободной форме
	Reason string `json:"reason"`
}

// AssertErrorResponseRequired checks if the required fields are not zero-ed
func AssertErrorResponseRequired(obj ErrorResponse) error {
	elements := map[string]interface{}{
		"reason": obj.Reason,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertErrorResponseConstraints checks if the values respects the defined constraints
func AssertErrorResponseConstraints(obj ErrorResponse) error {
	return nil
}
