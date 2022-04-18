package main

import (
	"fmt"
	"strings"
)

type Label struct {
	ChainName string
	Address   string
}

func main() {
	var list = []Label{
		{
			"heco", "hecoaddress",
		},
		{
			"eth", "address2",
		},
		{
			"eth", "address3",
		},
		{
			"bsc", "bscaddress1",
		},
		{
			"bsc", "bscaddress1",
		},
		{
			"heco", "hecoaddress",
		},
	}

	chainName2Address := make(map[string][]string, len(list))
	for _, label := range list {
		chainName := strings.TrimSpace(label.ChainName)
		addressList := chainName2Address[chainName]

		addr := strings.TrimSpace(label.Address)
		if !isExist(addressList, addr) {
			addressList = append(addressList, addr)
		}
		chainName2Address[chainName] = addressList
	}

	fmt.Println(chainName2Address)
}

func isExist(list []string, addr string) bool {
	for _, s := range list {
		if s == addr {
			return true
		}
	}
	return false
}
