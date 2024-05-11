package crypto

import (
	"testing"

	"github.com/minnowo/astoryofand/internal/assets"
)

func TestOrderPGPSign(t *testing.T) {

	orderEncryption := &PGPEncryptionWriter{
		PublicKey:       assets.PublicKeyBytes,
		OutputDirectory: t.TempDir(),
	}

	var data []byte = []byte("{ \"this\" : \"is some data\" }")

	if _, err := orderEncryption.SaveAndEncryptData("TEST_ORDER_PGP", data); err != nil {

		t.Errorf("Could not encrypt the order data: %v", err)
	}
}

func TestUsesPGPSign(t *testing.T) {

	usesEncryption := &PGPEncryptionWriter{
		PublicKey:       assets.PublicKeyBytes,
		OutputDirectory: t.TempDir(),
	}

	var data []byte = []byte("{ \"this\" : \"is some data\" }")

	if _, err := usesEncryption.SaveAndEncryptData("TEST_ORDER_PGP", data); err != nil {

		t.Errorf("Could not encrypt the uses data: %v", err)
	}
}
