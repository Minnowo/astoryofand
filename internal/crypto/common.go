package crypto

type EncryptionWriter interface {
	SaveAndEncryptData([]byte) (string, error)
	EnsureCanWriteDiskOrExit()
}
