package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func Sha2Encrypt(raw string) string {
	hash := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hash[:])
}

func AES256Encrypt(key string, data string) (string, error) {
	text := []byte(data)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	encryptor := cipher.NewCFBEncrypter(block, iv)

	ciphertext := make([]byte, len(text))
	encryptor.XORKeyStream(ciphertext, text)
	return hex.EncodeToString(ciphertext), nil
}

func AES256Decrypt(key string, data string) (string, error) {
	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	decryptor := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	decryptor.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func Http(uri string, method string, ptr interface{}, headers map[string]string, body io.Reader) (err error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(ptr); err != nil {
		return err
	}
	return nil
}

func Get(uri string, headers map[string]string) (data interface{}, err error) {
	err = Http(uri, http.MethodGet, &data, headers, nil)
	return data, err
}

func Post(uri string, headers map[string]string, body interface{}) (data interface{}, err error) {
	var form io.Reader
	if buffer, err := json.Marshal(body); err == nil {
		form = bytes.NewBuffer(buffer)
	}
	err = Http(uri, http.MethodPost, &data, headers, form)
	return data, err
}

func PostForm(uri string, body map[string]interface{}) (data map[string]interface{}, err error) {
	client := &http.Client{}
	form := make(url.Values)
	for key, value := range body {
		form[key] = []string{value.(string)}
	}
	res, err := client.PostForm(uri, form)
	if err != nil {
		return nil, err
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return data, nil
}
