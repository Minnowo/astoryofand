package assets

import (
	_ "embed"
)

//go:embed go_pub.gpg
var PublicKeyBytes []byte

const PGPOutputDir string = "./pgp_out"
