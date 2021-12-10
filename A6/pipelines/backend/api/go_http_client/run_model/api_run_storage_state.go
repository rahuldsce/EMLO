// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// APIRunStorageState api run storage state
// swagger:model apiRunStorageState
type APIRunStorageState string

const (

	// APIRunStorageStateSTORAGESTATEAVAILABLE captures enum value "STORAGESTATE_AVAILABLE"
	APIRunStorageStateSTORAGESTATEAVAILABLE APIRunStorageState = "STORAGESTATE_AVAILABLE"

	// APIRunStorageStateSTORAGESTATEARCHIVED captures enum value "STORAGESTATE_ARCHIVED"
	APIRunStorageStateSTORAGESTATEARCHIVED APIRunStorageState = "STORAGESTATE_ARCHIVED"
)

// for schema
var apiRunStorageStateEnum []interface{}

func init() {
	var res []APIRunStorageState
	if err := json.Unmarshal([]byte(`["STORAGESTATE_AVAILABLE","STORAGESTATE_ARCHIVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiRunStorageStateEnum = append(apiRunStorageStateEnum, v)
	}
}

func (m APIRunStorageState) validateAPIRunStorageStateEnum(path, location string, value APIRunStorageState) error {
	if err := validate.Enum(path, location, value, apiRunStorageStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this api run storage state
func (m APIRunStorageState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPIRunStorageStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
