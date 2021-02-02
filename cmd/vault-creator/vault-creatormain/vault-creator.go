package vault_creatormain

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sosedoff/ansible-vault-go"
	"golang.org/x/term"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func Run() {

	// Encrypt secret data
	//str, err := vault.Encrypt("secret", "password")
	// Decrypt secret data
	//str, err := vault.Decrypt("secret", "password")

	fmt.Print("Enter Vault Password:\n -> ")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	check(err)
	pass := string(bytepw)
	fmt.Println(pass)

	err = filepath.Walk("./test",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fmt.Println(path, info.Size())
				encript(path, readfile(path), pass)
				return nil
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

}

func readfile(path string) string {
	content, err := ioutil.ReadFile(path)
	check(err)
	return string(content)
}

func encript(filepath string, secret string, pass string) {
	// Write secret data to file
	err := vault.EncryptFile(filepath, secret, pass)
	check(err)
}

func decript(filepath string, pass string) {
	// Read existing secret
	str, err := vault.DecryptFile(filepath, pass)
	check(err)
	fmt.Println(str)
}
