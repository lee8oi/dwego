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
	"os"
	"os/user"
)

var confdir string
var keypath string
var crypt Crypt

const sep = string(os.PathSeparator)

type Crypt struct {
	key []byte
}

func init() {
	//make sure confdir is available and secure
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	confdir = usr.HomeDir + sep + ".dwego"
	keypath = confdir + sep + "keyfile"
	err = os.Mkdir(confdir, 0700)
	if os.IsExist(err) {
		//always (try to) make sure confdir permissions are secure
		err := os.Chmod(confdir, 0700)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = crypt.LoadKey(keypath)
	if err != nil && !os.IsExist(err) {
		crypt.NewKey()
		crypt.SaveKey(keypath)
	}
	//Try these one at a time, reloading program between, to test homedir key file.
	//crypt.testsave("test")
	//crypt.testload("test")
}

//NewKey sets Crypt.key to a randomized 32 byte key to be used in encryption.
func (c *Crypt) NewKey() {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 32)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	c.key = bytes
}

//ObjectToFile converts a data object to json then saves it to file with Crypt.SaveFile.
func (c *Crypt) ObjectToFile(path string, m interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.SaveFile(path, b)
}

//FileToObject loads a file, saved with Crypt.ObjectToFile, into the given object variable.
func (c *Crypt) FileToObject(path string, m interface{}) {
	b, err := c.LoadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(b, m)
	if err != nil {
		fmt.Println(err)
	}
}

//SaveFile encodes & encrypts byte data then writes it to file.
func (c *Crypt) SaveFile(path string, b []byte) error {
	return ioutil.WriteFile(path, c.encrypt(b), 0644)
}

//LoadFile decrypts & decodes a file that has been saved with Crypt.SaveFile.
func (c *Crypt) LoadFile(path string) (b []byte, e error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		e = err
	} else {
		b = c.decrypt(b)
	}
	return
}

//SaveKey encodes current Crypt.key in Base64 and saves it to file.
func (c *Crypt) SaveKey(path string) error {
	s := encodeBase64(c.key)
	return ioutil.WriteFile(path, []byte(s), 0600)
}

//LoadKey decodes a Base64 encoded key file & sets Crypt.key.
func (c *Crypt) LoadKey(path string) (e error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		e = err
	} else {
		c.key = decodeBase64(string(b))
	}
	return
}

//encrypt encodes bytes in base64 then encrypts data with AES.
func (c *Crypt) encrypt(text []byte) []byte {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		panic(err)
	}
	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext
}

//decrypt data that has been encrypted with Crypt.encrypt.
func (c *Crypt) decrypt(text []byte) []byte {
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
	return decodeBase64(string(text))
}

func (c *Crypt) testsave(path string) {
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
	c.ObjectToFile(path, &group)
}

func (c *Crypt) testload(path string) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	var ngroup ColorGroup
	c.FileToObject(path, &ngroup)
	fmt.Println(ngroup.Name)
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
