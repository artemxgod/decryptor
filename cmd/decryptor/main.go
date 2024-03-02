package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/artemxgod/decryptor/config"
	"github.com/artemxgod/decryptor/internal/decryptor"
	"github.com/artemxgod/decryptor/pkg/fernet"
	"github.com/atotto/clipboard"
)

type MessageDecryptor struct {
	message   string
	decryptor decryptor.Decryptor
}

func main() {

	cfgPath, err := configPath()
	if err != nil {
		log.Fatal(err)
	}
	viper, err := config.LoadConfig(cfgPath)
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
		caseEncrypt(in, md)
	case *decrypt:
		caseDecrypt(in, md)
	case *newKey:
		caseNewKey(in)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}

}

func caseEncrypt(in *bufio.Reader, md MessageDecryptor) {
	fmt.Println("Write down the message to encrypt")
	if _, err := fmt.Fscan(in, &md.message); err != nil {
		log.Fatal(err)
	}

	cr, err := md.decryptor.EncryptMessage(md.message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ENCRYPTED: ", cr)
	clipboard.WriteAll(cr)
}

func caseDecrypt(in *bufio.Reader, md MessageDecryptor) {
	fmt.Println("Write down the message to decrypt")
	if _, err := fmt.Fscan(in, &md.message); err != nil {
		log.Fatal(err)
	}

	cr := md.decryptor.DecryptMessage(md.message)
	fmt.Println("DECRYPTED:", cr)
	// clipboard.WriteAll(cr)
}

func caseNewKey(in *bufio.Reader) {
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
		caseEncrypt(in, md)
	case 2:
		caseDecrypt(in, md)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

func configPath() (string, error) {
	// exePath, err := os.Executable()
	// if err != nil {
	// 	return "", err
	// }
	// exeDir := filepath.Dir(exePath)

	// return filepath.Join(exeDir, "config"), nil

	return "/Users/artemgod/github.com/artemxgod/personal/decryptor/config", nil
}
