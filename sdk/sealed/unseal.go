package sealed

import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk"
	"io"
)

type Algorithm string

var (
	sealHeader = []byte{0x9E, 0x85, 0xDC, 0xED}

	ErrInvalidHeader        = errors.New("invalid sealed data header")
	ErrInvalidKeys          = errors.New("invalid decryption keys")
	ErrInvalidAlgorithm     = errors.New("invalid decryption algorithm")
	ErrInvalidEventResponse = errors.New("invalid event response")
)

const (
	AlgorithmAES256GCM Algorithm = "aes-256-gcm"
)

type DecryptionKey struct {
	Key       []byte
	Algorithm Algorithm
}

// UnsealEventsResponse decrypts the sealed response with the provided keys.
// The SDK will try to decrypt the result with each key until it succeeds.
// In case if all keys fail, AggregatedUnsealError is returned with error details for each key.
// To learn more about sealed results visit: https://dev.fingerprint.com/docs/sealed-client-results
func UnsealEventsResponse(sealed []byte, keys []DecryptionKey) (*sdk.EventResponse, error) {
	unsealed, err := Unseal(sealed, keys)

	if err != nil {
		return nil, err
	}

	var eventResponse sdk.EventResponse

	err = json.Unmarshal(unsealed, &eventResponse)

	if err != nil {
		return nil, err
	}

	if eventResponse.Products == nil {
		return nil, ErrInvalidEventResponse
	}

	return &eventResponse, nil
}

func Unseal(sealed []byte, keys []DecryptionKey) ([]byte, error) {
	if len(sealed) < len(sealHeader) || !bytes.Equal(sealed[:len(sealHeader)], sealHeader) {
		return nil, ErrInvalidHeader
	}

	aggregateError := NewAggregatedUnsealError()

	for _, key := range keys {
		switch key.Algorithm {
		case AlgorithmAES256GCM:
			payload, err := decryptAes256gcm(sealed[len(sealHeader):], key.Key)

			if err != nil {
				aggregateError.Add(NewUnsealError(err, key))

				break
			} else {
				return payload, nil
			}

		default:
			return nil, ErrInvalidAlgorithm
		}
	}

	return nil, aggregateError
}

func decryptAes256gcm(payload, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Join(errors.New("new cipher"), err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Join(errors.New("new GCM"), err)
	}

	if len(payload) < aesgcm.NonceSize() {
		return nil, errors.New("nonce")
	}

	nonce, ciphertext := payload[:aesgcm.NonceSize()], payload[aesgcm.NonceSize():]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.Join(errors.New("aesgcm open"), err)
	}

	payload, err = decompress(plaintext)

	if err != nil {
		return nil, errors.Join(errors.New("decompress"), err)
	}

	return payload, nil
}

func decompress(compressed []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(compressed))
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(err, errors.New("inflated payload read all bytes"))
	}

	return decompressed, nil
}
