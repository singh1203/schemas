// Package category provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package category

import "github.com/gofrs/uuid"

type CategoryDefinition struct {
	Id       uuid.UUID              `json:"-"`
	Name     string                 `json:"name" gorm:"name"`
	Metadata map[string]interface{} `json:"metadata"  yaml:"metadata" gorm:"type:bytes;serializer:json"`
}
