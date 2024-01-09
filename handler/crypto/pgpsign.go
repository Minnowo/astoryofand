package crypto

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/assets"
)

func FailIfPGPDirNotExists() {

	if _, err := os.Stat(assets.PGPOutputDir); os.IsNotExist(err) {

		err := os.Mkdir(assets.PGPOutputDir, os.ModePerm)

		if err != nil {
			log.Error("error: Cannot create ", assets.PGPOutputDir, ", and it does not exist!\n")
			os.Exit(1)
		}
	}
}

func GetNewOrderName() string {

	formattedDateTime := time.Now().Format("2006-01-02_15-04-05")

	uid := uuid.NewString()

	return fmt.Sprintf("%s%s.asc", formattedDateTime, uid)

}

func WritePGPOrder(json []byte) error {

	armor, err_ := helper.EncryptBinaryMessageArmored(string(assets.PublicKeyBytes), json)

	if err_ != nil {
		return err_
	}

	outfile := filepath.Join(assets.PGPOutputDir, GetNewOrderName())

	err := os.WriteFile(outfile, []byte(armor), 0644)

	if err != nil {
		return err
	}

	return nil
}
