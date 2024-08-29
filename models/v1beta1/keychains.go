// Package v1beta1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1beta1

import (
	"database/sql"

	"github.com/gofrs/uuid"
	externalRef1 "github.com/meshery/schemas/models/core"
)

// Keychain defines model for keychain.
type Keychain struct {
	ID uuid.UUID `db:"id" json:"id"`

	// CreatedAt Timestamp when the resource was created.
	CreatedAt externalRef1.CreatedAt `db:"created_at" json:"created_at,omitempty" yaml:"created_at"`
	DeletedAt sql.NullTime           `db:"deleted_at" json:"deleted_at"`
	Name      string                 `db:"name" json:"name"`
	Owner     uuid.UUID              `db:"owner" json:"owner"`

	// UpdatedAt Timestamp when the resource was updated.
	UpdatedAt externalRef1.UpdatedAt `db:"updated_at" json:"updated_at,omitempty" yaml:"updated_at"`
}

// KeychainFilter defines model for keychainFilter.
type KeychainFilter struct {
	RoleId uuid.UUID `db:"role_id" json:"role_id"`
}

// KeychainsKeysMapping defines model for keychainsKeysMapping.
type KeychainsKeysMapping struct {
	// CreatedAt Timestamp when the resource was created.
	CreatedAt  externalRef1.CreatedAt `db:"created_at" json:"created_at,omitempty" yaml:"created_at"`
	DeletedAt  sql.NullTime           `db:"deleted_at" json:"deleted_at"`
	Id         externalRef1.GeneralId `db:"id" json:"id"`
	KeyId      uuid.UUID              `db:"key_id" json:"key_id"`
	KeychainId uuid.UUID              `db:"keychain_id" json:"keychain_id"`

	// UpdatedAt Timestamp when the resource was updated.
	UpdatedAt externalRef1.UpdatedAt `db:"updated_at" json:"updated_at,omitempty" yaml:"updated_at"`
}

// KeychainsKeysMappingPage defines model for keychainsKeysMappingPage.
type KeychainsKeysMappingPage struct {
	KeychainsKeysMappings []KeychainsKeysMapping `json:"keychains_keys_mappings,omitempty"`
	Page                  externalRef1.Number    `json:"page,omitempty"`
	PageSize              externalRef1.Number    `json:"page_size,omitempty"`
	TotalCount            externalRef1.Number    `json:"total_count,omitempty"`
}

// KeychainsPage defines model for keychainsPage.
type KeychainsPage struct {
	Keychains  []Keychain          `json:"keychains,omitempty"`
	Page       externalRef1.Number `json:"page,omitempty"`
	PageSize   externalRef1.Number `json:"page_size,omitempty"`
	TotalCount externalRef1.Number `json:"total_count,omitempty"`
}
