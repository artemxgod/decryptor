package fernet

import (
	"time"

	"github.com/fernet/fernet-go"
)

type Crypt struct {
	EncryptionKey *fernet.Key
}

func NewCrypt(key string) *Crypt {
	k := fernet.MustDecodeKeys(key)
	return &Crypt{
		EncryptionKey: k[0],
	}
}

func (s *Crypt) DecryptMessage(cipher string) string {
	decryptedMsg := fernet.VerifyAndDecrypt([]byte(cipher), 0*time.Second, []*fernet.Key{s.EncryptionKey})
	return string(decryptedMsg)
}

func (s *Crypt) EncryptMessage(message string) (string, error) {
	EncryptedMsg, err := fernet.EncryptAndSign([]byte(message), s.EncryptionKey)
	return string(EncryptedMsg), err
}
