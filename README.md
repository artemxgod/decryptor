# Decryptor CLI

### Flags

- `--en` encrypts the message
- `--de` decrypts the message
- `--new` sets new encryption key and then encrypts/decrypts the message

> in main.go set your absolute config path

### Features
- Encrypted message will save automaticly to your clipboard

### Build app:
- `go build -o decryptor ./cmd/decryptor`

### to run executable from any folder
#### for mac:
- sudo cp "your_path_to_executable/decryptor" /usr/local/bin
- chmod +x /usr/local/bin/decryptor