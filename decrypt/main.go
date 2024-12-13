package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"os"
)

func main() {
	var mypath string
	var code string

	fmt.Println("please enter the path")
	fmt.Scan(&mypath)

	fmt.Println("please enter the keyword")
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
				decryptFile(actpath+"/"+file.Name(), code)
			} else {
				fmt.Println(actpath + file.Name())
				decryptFile(actpath+file.Name(), code)
			}
		}
	}
}
func decryptFile(myfilepath string, code string) {
	// Reading ciphertext file
	cipherText, err := os.ReadFile(myfilepath)
	if err != nil {
		log.Fatal(err)
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

	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	// Writing decryption content
	err = os.WriteFile(myfilepath, plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}
