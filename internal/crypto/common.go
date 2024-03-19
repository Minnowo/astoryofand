package crypto

import (
	"strings"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
)

type EncryptionWriter interface {
	SaveAndEncryptData(string, []byte) (string, error)
	EnsureCanWriteDiskOrExit()
}

func IsPublicKey(data string) bool {

	block, err := armor.Decode(strings.NewReader(data))

	if err != nil {
		return false
	}

	if block.Type != openpgp.PublicKeyType && block.Type != openpgp.PrivateKeyType {
		return false
	}

	return false
}
