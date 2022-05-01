package main

import (
	"GoProjectDemo/btc-demo/day4-v5/lib/base58"
	"GoProjectDemo/btc-demo/day4-v5/lib/ripemd160"
	"bytes"
	"fmt"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	// "github.com/btcsuite/btcutil/base58"
	// "golang.org/x/crypto/ripemd160"
	"log"
)

// 这里的钱包是一个结构   每个钱包保存了公钥 私钥对

type Wallet struct {
	Private *ecdsa.PrivateKey

	// 	这里publickey不存原始的公钥  而是存储x y 拼接的字符串   在校验端重新拆分
	PubKey []byte
}

// NewWallet 创建钱包
func NewWallet() *Wallet {
	// 创建曲线
	p256 := elliptic.P256()
	// 生成私钥
	privateKey, err := ecdsa.GenerateKey(p256, rand.Reader)
	if err != nil {
		log.Panic("生成私钥出错,err", err.Error())
	}

	// 生成公钥
	publicKeyOrig := privateKey.PublicKey

	// 拼接  X Y
	pubKey := append(publicKeyOrig.X.Bytes(), publicKeyOrig.Y.Bytes()...)

	return &Wallet{
		Private: privateKey,
		PubKey:  pubKey,
	}
}

// NewAddress 生成地址
func (w *Wallet) NewAddress() (addr string) {
	pubKey := w.PubKey

	rip160HashValue := HashPubKey(pubKey)
	version := byte(00)
	// 拼接version
	payload := append([]byte{version}, rip160HashValue...)

	// checksum
	checkCode := CheckSum(payload)

	// 25字节数据
	payload = append(payload, checkCode...)

	// 25字节数据
	payload = append(payload, checkCode...)

	addr = base58.Encode(payload)
	return
}

func HashPubKey(data []byte) []byte {
	hash := sha256.Sum256(data)

	// 理解成编码器
	rip160Hasher := ripemd160.New()
	_, err := rip160Hasher.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}

	// 返回rip160的哈希结果
	rip160HasherValue := rip160Hasher.Sum(nil)
	return rip160HasherValue
}

func CheckSum(data []byte) []byte {
	// 两次sha256
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])

	// 前4字节校验码
	checkCode := hash2[:4]
	return checkCode
}

func IsValidAddress(address string) bool {
	// 1. 解码
	addressByte := base58.Decode(address)

	if len(addressByte) < 4 {
		return false
	}

	// 2. 取数据
	payload := addressByte[:len(addressByte)-4]
	checksum1 := addressByte[len(addressByte)-4:]

	// 3. 做checksum函数
	checksum2 := CheckSum(payload)

	fmt.Printf("checksum1 : %x\n", checksum1)
	fmt.Printf("checksum2 : %x\n", checksum2)

	// 4. 比较
	return bytes.Equal(checksum1, checksum2)
}
