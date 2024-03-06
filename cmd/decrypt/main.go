package main

import (
	"flag"
	"os"
	"path"
	"path/filepath"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/assets"
)

func dirExists(path string, mkiffail bool) (bool, error) {

	_, err := os.Stat(path)

	if os.IsNotExist(err) {

		if !mkiffail {
			return false, nil
		}

		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return false, err
		}

		_, err = os.Stat(path)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func decryptFilesIn(srcDir, dstDir, passwordStr string) {

	files, err := os.ReadDir(srcDir)

	if err != nil {
		log.Fatal(err)
	}

	var password []byte = []byte(passwordStr)

	if len(password) == 0 {
		password = nil
	}

	log.Infof("Searching files in %s", srcDir)

	for _, file := range files {

		if filepath.Ext(file.Name()) != ".asc" {
			continue
		}

		log.Infof("   %s", file.Name())

		var srcPath, dstPath string
		srcPath = path.Join(srcDir, file.Name())
		dstPath = path.Join(dstDir, file.Name())

		fileContent, err := os.ReadFile(srcPath)

		if err != nil {
			log.Fatal(err)
		}

		data, err := helper.DecryptBinaryMessageArmored(assets.PrivateKeyBytes, password, string(fileContent))

		if err != nil {
			log.Errorf("Could not decrpt file: %v", err)
			continue
		}

		os.WriteFile(dstPath, data, 0644)
	}
}

func main() {

	log.SetHeader("${level}")
	log.SetLevel(log.INFO)

	var inputDir string
	var outputDir string
	var password string

	flag.StringVar(&inputDir, "i", "./enc", "Directory")
	flag.StringVar(&outputDir, "o", "./encd", "Directory")
	flag.StringVar(&password, "p", "", "Password for private key")

	flag.Parse()

	// ensure input dir exists
	var s, err = dirExists(inputDir, false)

	if !s {

		if err != nil {
			log.Error(err)
		}

		log.Fatalf("Input directory %s does not exist!", inputDir)
	}

	// ensure output dir exists / is created
	s, err = dirExists(outputDir, true)

	if !s {

		if err != nil {
			log.Error(err)
		}

		log.Fatal("Output directory does not exist!")
	}

	decryptFilesIn(
		path.Join(inputDir, "pgp_out"),
		outputDir,
		password,
	)

	decryptFilesIn(
		path.Join(inputDir, "uses"),
		outputDir,
		password,
	)

}
