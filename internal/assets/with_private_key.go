//go:build include_private_key
// +build include_private_key

package assets

import (
	_ "embed"
)

//go:embed go_pri.gpg
var PrivateKeyBytes string

//go:embed go_pri_password.gpg
var PrivateKeyPassword string
