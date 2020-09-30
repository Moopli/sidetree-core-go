/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package model

import (
	"github.com/trustbloc/sidetree-core-go/pkg/jws"
	"github.com/trustbloc/sidetree-core-go/pkg/patch"
)

// CreateRequest is the struct for create payload
type CreateRequest struct {
	// operation
	// Required: true
	Operation OperationType `json:"type,omitempty"`

	// Encoded suffix data object
	// Required: true
	SuffixData string `json:"suffix_data"`

	// Encoded delta object
	// Required: true
	Delta string `json:"delta"`
}

// CreateRequestJCS is the struct for create payload JCS
type CreateRequestJCS struct {
	// operation
	// Required: true
	Operation OperationType `json:"type,omitempty"`

	// Encoded suffix data object
	// Required: true
	SuffixData *SuffixDataModel `json:"suffix_data,omitempty"`

	// Encoded delta object
	// Required: true
	Delta *DeltaModel `json:"delta,omitempty"`
}

// SuffixDataModel is part of create request
type SuffixDataModel struct {

	// Hash of the delta object
	DeltaHash string `json:"delta_hash,omitempty"`

	// Commitment hash for the next recovery or deactivate operation
	RecoveryCommitment string `json:"recovery_commitment,omitempty"`
}

// DeltaModel contains patch data (patches used for create, recover, update)
type DeltaModel struct {

	// Commitment hash for the next update operation
	UpdateCommitment string `json:"update_commitment,omitempty"`

	// Patches defines document patches
	Patches []patch.Patch `json:"patches,omitempty"`
}

//UpdateRequest is the struct for update request
type UpdateRequest struct {
	Operation OperationType `json:"type"`

	//The suffix of the DID
	DidSuffix string `json:"did_suffix"`

	// Compact JWS - signature information
	SignedData string `json:"signed_data"`

	// Encoded delta object
	Delta string `json:"delta"`
}

//DeactivateRequest is the struct for deactivating document
type DeactivateRequest struct {
	// operation
	// Required: true
	Operation OperationType `json:"type"`

	//The suffix of the DID
	// Required: true
	DidSuffix string `json:"did_suffix"`

	// Compact JWS - signature information
	SignedData string `json:"signed_data"`
}

// UpdateSignedDataModel defines signed data model for update
type UpdateSignedDataModel struct {
	// The current update key
	UpdateKey *jws.JWK `json:"update_key"`

	// Hash of the unsigned delta object
	DeltaHash string `json:"delta_hash"`
}

// RecoverSignedDataModel defines signed data model for recovery
type RecoverSignedDataModel struct {

	// Hash of the unsigned delta object
	DeltaHash string `json:"delta_hash"`

	// The current recovery key
	RecoveryKey *jws.JWK `json:"recovery_key"`

	// Recovery commitment be used for the next recovery/deactivate
	RecoveryCommitment string `json:"recovery_commitment"`
}

// DeactivateSignedDataModel defines data model for deactivate
type DeactivateSignedDataModel struct {

	//The suffix of the DID
	// Required: true
	DidSuffix string `json:"did_suffix"`

	// The current recovery key
	RecoveryKey *jws.JWK `json:"recovery_key"`
}

// RecoverRequest is the struct for document recovery payload
type RecoverRequest struct {
	// operation
	// Required: true
	Operation OperationType `json:"type"`

	//The suffix of the DID
	// Required: true
	DidSuffix string `json:"did_suffix"`

	// Compact JWS - signature information
	SignedData string `json:"signed_data"`

	// Encoded delta object
	// Required: true
	Delta string `json:"delta"`
}
