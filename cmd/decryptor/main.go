package main

import (
	"bufio"
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

	var md MessageDecryptor
	md.decryptor = fernet.NewCrypt(cfg.EncryptionKey)

	fmt.Println("Write down the message to decrypt")
	in := bufio.NewReader(os.Stdin)
	
	if _, err := fmt.Fscan(in, &md.message); err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULT:", md.decryptor.DecryptMessage(md.message))
}
