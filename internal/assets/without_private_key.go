//go:build !include_private_key
// +build !include_private_key

package assets

// we never want to include the private key
// unless we are building the program meant to contain it
var PrivateKeyBytes string = ""
