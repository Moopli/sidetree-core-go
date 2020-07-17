/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package batch

import (
	"github.com/trustbloc/sidetree-core-go/pkg/restapi/model"
)

// Operation defines an operation
type Operation struct {

	//Operation type
	Type OperationType `json:"type"`

	Namespace string `json:"namespace"`

	//ID is full ID for this document - includes namespace + unique suffix
	ID string `json:"id"`

	//The unique suffix - encoded hash of the original create document
	UniqueSuffix string `json:"uniqueSuffix"`

	// OperationBuffer is the original operation request
	OperationBuffer []byte `json:"operationBuffer"`

	// Compact JWS - signed data for the operation
	SignedData string `json:"signedData"`

	// operation delta
	Delta *model.DeltaModel `json:"delta"`

	// encoded delta
	EncodedDelta string `json:"encodedDelta"`

	// suffix data
	SuffixData *model.SuffixDataModel `json:"suffixData"`

	// encoded suffix data
	EncodedSuffixData string `json:"encodedSuffixData"`

	//The logical blockchain time that this operation was anchored on the blockchain
	TransactionTime uint64 `json:"transactionTime"`
	//The transaction number of the transaction this operation was batched within
	TransactionNumber uint64 `json:"transactionNumber"`
	//The index this operation was assigned to in the batch
	OperationIndex uint `json:"operationIndex"`
}

// AnchoredOperation defines an anchored operation
type AnchoredOperation struct {

	//Operation type
	Type OperationType `json:"type"`

	//The unique suffix - encoded hash of the original create document
	UniqueSuffix string `json:"uniqueSuffix"`

	// Compact JWS - signed data for the operation
	SignedData string `json:"signedData,omitempty"`

	// encoded delta
	EncodedDelta string `json:"encodedDelta,omitempty"`

	// encoded suffix data
	EncodedSuffixData string `json:"encodedSuffixData,omitempty"`

	//The logical blockchain time that this operation was anchored on the blockchain
	TransactionTime uint64 `json:"transactionTime"`
	//The transaction number of the transaction this operation was batched within
	TransactionNumber uint64 `json:"transactionNumber"`
	//The index this operation was assigned to in the batch
	OperationIndex uint `json:"operationIndex"`
}

// OperationType defines valid values for operation type
type OperationType string

const (

	// OperationTypeCreate captures "create" operation type
	OperationTypeCreate OperationType = "create"

	// OperationTypeUpdate captures "update" operation type
	OperationTypeUpdate OperationType = "update"

	// OperationTypeDeactivate captures "deactivate" operation type
	OperationTypeDeactivate OperationType = "deactivate"

	// OperationTypeRecover captures "recover" operation type
	OperationTypeRecover OperationType = "recover"
)

// OperationInfo contains the unique suffix and namespace as well as the operation buffer
type OperationInfo struct {
	Data         []byte
	UniqueSuffix string
	Namespace    string
}
