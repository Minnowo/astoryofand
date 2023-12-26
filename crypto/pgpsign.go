package crypto

import (
	"fmt"
	"minno/astoryofand/assets"
	"os"
	"path/filepath"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/google/uuid"
)

func FailIfPGPDirNotExists() {

	if _, err := os.Stat(assets.PGPOutputDir); os.IsNotExist(err) {

		err := os.Mkdir(assets.PGPOutputDir, os.ModePerm)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: Cannot create %s, and it does not exist!\n", assets.PGPOutputDir)
			os.Exit(1)
		}
	}
}

func GetNewOrderName() string {

	formattedDateTime := time.Now().Format("2006-01-02_15-04-05")

	uid := uuid.NewString()

	return fmt.Sprintf("%s%s.asc", formattedDateTime, uid)

}

func WritePGPOrder(json []byte) bool {

	armor, err_ := helper.EncryptBinaryMessageArmored(string(assets.PublicKeyBytes), json)

	if err_ != nil {
		fmt.Println("Could not encrypt message with armor")
		return false
	}

	outfile := filepath.Join(assets.PGPOutputDir, GetNewOrderName())

	err := os.WriteFile(outfile, []byte(armor), 0644)

	if err != nil {
		fmt.Println("Could not encrypt message with armor")
		return false
	}

	fmt.Println(armor)

	return true
}
