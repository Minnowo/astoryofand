package crypto

import (
	"os"
	"path/filepath"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/util"
)

type PGPEncryptionWriter struct {
	PublicKey       string
	OutputDirectory string
}

func (p *PGPEncryptionWriter) EnsureCanWriteDiskOrExit() {

	var fileInfo os.FileInfo
	var err error

	fileInfo, err = os.Stat(p.OutputDirectory)

	if os.IsNotExist(err) {

		if err = os.MkdirAll(p.OutputDirectory, os.ModePerm); err != nil {
			log.Fatal("error: Cannot create ", p.OutputDirectory, ", and it does not exist!\n")
		}

		fileInfo, err = os.Stat(p.OutputDirectory)
	}

	if err != nil {
		log.Fatal(err)
	}

	if !fileInfo.IsDir() {
		log.Fatal(p.OutputDirectory, " is not a directory")
	}

	if mode := fileInfo.Mode(); mode.Perm()&0200 == 0 {
		log.Fatal(p.OutputDirectory, " does not have write permissions")
	} else {
		log.Debug(p.OutputDirectory, " has file permissions ", mode)
	}
}

func (p *PGPEncryptionWriter) SaveAndEncryptData(uuid string, json []byte) (string, error) {

	armor, err := helper.EncryptBinaryMessageArmored(p.PublicKey, json)

	if err != nil {
		return "", err
	}

	outfile := filepath.Join(p.OutputDirectory, util.GetNewOrderName(uuid))

	err = os.WriteFile(outfile, []byte(armor), 0644)

	if err != nil {
		return "", err
	}

	return uuid, nil
}
