package crypto

import (
	"fmt"
	"path/filepath"

	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/util"
)

type FakeEncryptionWriter struct {
	AlwaysFailEncrypt       bool
	AlwaysFailWrite         bool
	AlwaysFailCanWriteCheck bool
	OutputDirectory         string
}

const FAIL_MSG string = "FakeEncryptionWriter always fails this!"

func (p *FakeEncryptionWriter) EnsureCanWriteDiskOrExit() {
	if p.AlwaysFailCanWriteCheck {
		log.Fatal("Cannot write to ", p.OutputDirectory, FAIL_MSG)
	}
}

func (p *FakeEncryptionWriter) SaveAndEncryptData(json []byte) (string, error) {

	log.Warn("Using FakeEncryptionWriter! Nothing is actually about to happen.")
	log.Info("Encrypting some raw byte data")
	log.Debug(json)

	if p.AlwaysFailEncrypt {
		return "", fmt.Errorf("Could not encrypt data. %s", FAIL_MSG)
	}

	orderId := util.GetOrderID()

	outfile := filepath.Join(p.OutputDirectory, util.GetNewOrderName(orderId))

	log.Info("Order number: ", orderId)

	if p.AlwaysFailWrite {
		return "", fmt.Errorf("Could not write encrypted data to a file. %s", FAIL_MSG)
	}

	log.Info("Order saved:", outfile)

	return orderId, nil
}
