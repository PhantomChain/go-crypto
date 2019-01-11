// This file is part of PHANTOM Go Crypto.
//
// (c) PhantomChain <info@phantom.org>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package crypto

import (
	"github.com/PhantomChain/go-crypto/crypto/base58"
)

func AddressFromPassphrase(passphrase string) (string, error) {
	privateKey, err := PrivateKeyFromPassphrase(passphrase)

	if err != nil {
		return "", err
	}

	return privateKey.ToAddress(), nil
}

func AddressToBytes(address string) ([]byte, error) {
	bytes, err := base58.Decode(address)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func ValidateAddress(address string) (bool, error) {
	bytes, err := AddressToBytes(address)

	if err != nil {
		return false, err
	}

	return Byte2Hex(GetNetwork().Version) == Hex2Byte(bytes[:1]), nil
}
