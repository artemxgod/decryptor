package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/artemxgod/decryptor/config"
	"github.com/artemxgod/decryptor/decryptor"
	"github.com/artemxgod/decryptor/pkg/fernet"
)

type MessageDecryptor struct {
	message   string
	decryptor decryptor.Decryptor
}

func main() {
	viper, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ParseConfig(viper)
	if err != nil {
		log.Fatal(err)
	}

	encrypt := flag.Bool("en", false, "ecrypts message")
	decrypt := flag.Bool("de", false, "decrypts message")
	newKey := flag.Bool("new", false, "write new key crypt key")

	flag.Parse()

	var md MessageDecryptor
	md.decryptor = fernet.NewCrypt(cfg.EncryptionKey)
	in := bufio.NewReader(os.Stdin)

	switch {
	case *encrypt:
		CaseEncrypt(in, md)
	case *decrypt:
		CaseDecrypt(in, md)
	case *newKey:
		CaseNewKey(in)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}

}

func CaseEncrypt(in *bufio.Reader, md MessageDecryptor) {
	fmt.Println("Write down the message to encrypt")
	if _, err := fmt.Fscan(in, &md.message); err != nil {
		log.Fatal(err)
	}

	cr, err := md.decryptor.EncryptMessage(md.message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ENCRYPTED: ", cr)
}

func CaseDecrypt(in *bufio.Reader, md MessageDecryptor) {
	fmt.Println("Write down the message to decrypt")
	if _, err := fmt.Fscan(in, &md.message); err != nil {
		log.Fatal(err)
	}

	cr := md.decryptor.DecryptMessage(md.message)
	fmt.Println("DECRYPTED:", cr)
}

func CaseNewKey(in *bufio.Reader) {
	var key string
	var choose int
	var md MessageDecryptor
	fmt.Println("Write down the encryption key:")
	// in.Reset(os.Stdin)
	if _, err := fmt.Fscan(in, &key); err != nil {
		log.Fatal(err)
	}
	md.decryptor = fernet.NewCrypt(key)



	fmt.Println("1. Encryption\n2. Decryption")
	if _, err := fmt.Fscan(in, &choose); err != nil {
		log.Fatal(err)
	}

	switch choose {
	case 1:
		CaseEncrypt(in, md)
	case 2:
		CaseDecrypt(in, md)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}