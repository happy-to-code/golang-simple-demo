package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

// 演示如何使用ecdsa生成公钥 私钥
// 签名校验
func main() {
	// 创建曲线
	p256 := elliptic.P256()
	// 生成私钥
	privateKey, err := ecdsa.GenerateKey(p256, rand.Reader)
	if err != nil {
		panic(err)
	}

	// 生成公钥
	publicKey := privateKey.PublicKey

	// 	--------------------
	data := "hello world"

	// 生成hash  散列值
	hash := sha256.Sum256([]byte(data))

	// 	签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic("sign err:", err)
	}

	// 	把r s 进行序列化传输
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("signature:%v\n", signature)
	fmt.Printf("signature:%s\n", string(signature))

	// 校验签名
	// 	1
	verify1 := ecdsa.Verify(&publicKey, hash[:], r, s)
	fmt.Println("verify1:", verify1)

	fmt.Println("==============================")
	// 	2
	r1 := big.Int{}
	s1 := big.Int{}
	// 	拆分signature  前办部分给  r                        后半部分给s
	r1.SetBytes(signature[0 : len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])
	verify2 := ecdsa.Verify(&publicKey, hash[:], &r1, &s1)
	fmt.Println("verify2:", verify2)
}
