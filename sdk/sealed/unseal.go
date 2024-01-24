package sealed

import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v3/sdk"
	"github.com/pkg/errors"
	"io"
)

type Algorithm string

var (
	sealHeader = []byte{0x9E, 0x85, 0xDC, 0xED}

	ErrInvalidHeader    = errors.New("invalid sealed data header")
	ErrInvalidKeys      = errors.New("invalid decryption keys")
	ErrInvalidAlgorithm = errors.New("invalid decryption algorithm")
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

	return &eventResponse, nil
}

func Unseal(sealed []byte, keys []DecryptionKey) ([]byte, error) {
	if !bytes.Equal(sealed[:len(sealHeader)], sealHeader) {
		return nil, ErrInvalidHeader
	}

	for _, key := range keys {
		switch key.Algorithm {
		case AlgorithmAES256GCM:
			payload, err := decryptAes256gcm(sealed[len(sealHeader):], key.Key)

			if err == nil {
				return payload, nil
			}

			break

		default:
			return nil, ErrInvalidAlgorithm
		}
	}

	return nil, ErrInvalidKeys
}

func decryptAes256gcm(payload, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "new cipher")
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "new GCM")
	}

	if len(payload) < aesgcm.NonceSize() {
		return nil, errors.New("nonce")
	}

	nonce, ciphertext := payload[:aesgcm.NonceSize()], payload[aesgcm.NonceSize():]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.Wrap(err, "aesgcm open")
	}

	payload, err = decompress(plaintext)

	if err != nil {
		return nil, errors.Wrap(err, "decompress")
	}

	return payload, nil
}

func decompress(compressed []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(compressed))
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "inflated payload read all bytes")
	}

	return decompressed, nil
}
