/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package docutil

import (
	"crypto"
	"errors"
	"fmt"
	"hash"

	"github.com/multiformats/go-multihash"
)

const sha2_256 = 18

// ComputeMultihash will compute the hash for the supplied bytes using multihash code.
func ComputeMultihash(multihashCode uint, bytes []byte) ([]byte, error) {
	h, err := GetHash(multihashCode)
	if err != nil {
		return nil, err
	}

	if _, hashErr := h.Write(bytes); hashErr != nil {
		return nil, hashErr
	}
	hash := h.Sum(nil)

	return multihash.Encode(hash, uint64(multihashCode))
}

// GetHash will return hash based on specified multihash code.
func GetHash(multihashCode uint) (h hash.Hash, err error) {
	switch multihashCode {
	case sha2_256:
		h = crypto.SHA256.New()
	default:
		err = fmt.Errorf("algorithm not supported, unable to compute hash")
	}

	return h, err
}

// IsSupportedMultihash checks to see if the given encoded hash has been hashed using valid multihash code.
func IsSupportedMultihash(encodedMultihash string) bool {
	code, err := GetMultihashCode(encodedMultihash)
	if err != nil {
		return false
	}

	return multihash.ValidCode(code)
}

// IsComputedUsingHashAlgorithm checks to see if the given encoded hash has been hashed using multihash code.
func IsComputedUsingHashAlgorithm(encodedMultihash string, code uint64) bool {
	mhCode, err := GetMultihashCode(encodedMultihash)
	if err != nil {
		return false
	}

	return mhCode == code
}

// GetMultihashCode returns multihash code from encoded multihash.
func GetMultihashCode(encodedMultihash string) (uint64, error) {
	multihashBytes, err := DecodeString(encodedMultihash)
	if err != nil {
		return 0, err
	}

	mh, err := multihash.Decode(multihashBytes)
	if err != nil {
		return 0, err
	}

	return mh.Code, nil
}

// IsValidHash compares encoded content with encoded multihash.
func IsValidHash(encodedContent, encodedMultihash string) error {
	content, err := DecodeString(encodedContent)
	if err != nil {
		return err
	}

	code, err := GetMultihashCode(encodedMultihash)
	if err != nil {
		return err
	}

	computedMultihash, err := ComputeMultihash(uint(code), content)
	if err != nil {
		return err
	}

	encodedComputedMultihash := EncodeToString(computedMultihash)

	if encodedComputedMultihash != encodedMultihash {
		return errors.New("supplied hash doesn't match original content")
	}

	return nil
}
