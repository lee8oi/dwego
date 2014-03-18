// crypt.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

//func main() {
//	var C crypt
//	C.key = []byte("j19cQ2lSY9TnAJtezePsCfJdKtJ5Pdj3")
//	C.test2()
//	//ciph := C.encrypt([]byte("Some crazy text."))
//	//fmt.Println(ciph)
//	//fmt.Println(C.decrypt(ciph))
//}

type crypt struct {
	key []byte
}

func (c *crypt) test() {
	plaintext := []byte("some really really really long plaintext that goes on forever")
	fmt.Printf("%s\n", plaintext)
	ciphertext := c.encrypt(plaintext)
	fmt.Printf("%x\n", ciphertext)
	err := ioutil.WriteFile("ciphtest", ciphertext, 0644)
	if err != nil {
		fmt.Println(err)
	}
	b2, err := ioutil.ReadFile("ciphtest")
	if err != nil {
		fmt.Println(err)
	}
	result := c.decrypt(b2)
	fmt.Printf("%s\n", result)
}

func (c *crypt) test2() error {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		return err
	}
	c.encryptSave("encryptedjsfile", b)
	str, err := c.decryptLoad("encryptedjsfile")
	fmt.Println(str)
	return err
}

func (c *crypt) encryptSave(path string, b []byte) error {
	return ioutil.WriteFile(path, c.encrypt(b), 0644)
}

func (c *crypt) decryptLoad(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return c.decrypt(b), nil
}

func (c *crypt) encrypt(text []byte) []byte {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		panic(err)
	}
	b := c.encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext
}

func (c *crypt) decrypt(text []byte) string {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(c.decodeBase64(string(text)))
}

func (c *crypt) encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (c *crypt) decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func random(size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
