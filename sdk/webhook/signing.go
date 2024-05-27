package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func checkSignature(signature string, data []byte, secret string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	computedSignature := hex.EncodeToString(h.Sum(nil))
	return signature == computedSignature
}

// CheckHeader Verifies the HMAC signature extracted from the "fpjs-event-signature" header of the incoming request. This is a part of the webhook signing process, which is available only for enterprise customers.
// If you wish to enable it, please contact our support: https://fingerprint.com/support
func CheckHeader(header string, data []byte, secret string) bool {
	signatures := strings.Split(header, ",")

	for _, signature := range signatures {
		parts := strings.Split(signature, "=")
		if len(parts) == 2 && parts[0] == "v1" {
			hash := parts[1]
			if checkSignature(hash, data, secret) {
				return true
			}
		}
	}

	return false
}
