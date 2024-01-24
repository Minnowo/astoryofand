//go:build include_private_key
// +build include_private_key

package assets

import (
	_ "embed"
)

//go:embed go_pri.gpg
var PrivateKeyBytes string
