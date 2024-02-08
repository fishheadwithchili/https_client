package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

// 生成RSA公钥和私钥到指定路径
func GenerateRSAKeyPair(publicKeyPath, privateKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// 保存私钥
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})
	err = ioutil.WriteFile(privateKeyPath, privateKeyPEM, 0644)
	if err != nil {
		return err
	}

	// 保存公钥
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})
	err = ioutil.WriteFile(publicKeyPath, publicKeyPEM, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 使用RSA公钥和盐加密字符串
func EncryptStringWithRSA(publicKeyPath, content, salt string) (string, error) {
	// 读取公钥文件
	publicKeyPEM, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return "", err
	}

	// 解码公钥
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the public key")
	}

	// 解析公钥
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// 将盐和内容拼接后加密
	data := append([]byte(salt), []byte(content)...)
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), data)
	if err != nil {
		return "", err
	}

	// 返回Base64编码的加密结果
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// 使用RSA私钥和盐解密字符串
func DecryptStringWithRSA(privateKeyPath, content, salt string) (string, error) {
	// 读取私钥文件
	privateKeyPEM, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}

	// 解码私钥
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the private key")
	}

	// 解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// 解码Base64的加密结果
	encrypted, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}

	// 使用私钥解密后截取盐和内容
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted)
	if err != nil {
		return "", err
	}

	return string(decrypted[len(salt):]), nil
}

func Test() {
	err := GenerateRSAKeyPair("./public_key.pem", "./private_key.pem")
	if err != nil {
		panic(err)
	}
	var content = "hellogood mo123fsafdsafsadfsad"
	var salt = "鹽"
	secret, err := EncryptStringWithRSA("./public_key.pem", content, salt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("from %v secret--> %v\n", len(content), len(secret))
	output, err := DecryptStringWithRSA("./private_key.pem", secret, salt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("output--> %s\n", output)

}
