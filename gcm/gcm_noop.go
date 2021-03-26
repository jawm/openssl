// +build !gcm

package gcm

import "crypto/cipher"

func NewUnauthenticatedGCM(isEncrypter bool, key []byte, iv []byte) (cipher.BlockMode, error) {
	return nil, nil
}
