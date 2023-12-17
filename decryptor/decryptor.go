package decryptor

type Decryptor interface {
	DecryptMessage(string) string
	EncryptMessage(string) (string, error)
}