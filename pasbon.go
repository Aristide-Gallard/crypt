package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	mypath := "D:/école"
	var code string
	fmt.Scan(&code)
	treemaker(mypath, code)
}
func treemaker(actpath string, code string) {
	files, err := os.ReadDir(actpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Type().IsDir() {
			if actpath[len(actpath)-1:] != "/" {
				fmt.Println(actpath + "/" + file.Name())
				treemaker((actpath + "/" + file.Name()), code)
			} else {
				fmt.Println(actpath + file.Name())
				treemaker(actpath+file.Name(), code)
			}
		} else {
			if actpath[len(actpath)-1:] != "/" {
				fmt.Println(actpath + "/" + file.Name())
				encryptFile(actpath+"/"+file.Name(), code)
			} else {
				fmt.Println(actpath + file.Name())
				encryptFile(actpath+file.Name(), code)
			}
		}
	}
}
func encryptFile(myfilepath string, code string) {
	// Reading plaintext file
	plainText, err := os.ReadFile(myfilepath)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Reading key
	key := []byte(code)
	//key, err := os.ReadFile(keypath)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	// Writing ciphertext file
	err = os.WriteFile(myfilepath, cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}
