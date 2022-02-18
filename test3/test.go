package main

import (
	"bytes"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type TxHistory struct {
	Txhash         string     `json:"txhash"`
	From           string     `json:"from"`
	To             string     `json:"to"`
	Amount         *big.Float `json:"amount"`
	OriginalAmount string     `json:"original_amount"`
	Token          string     `json:"token"`
	GasUsed        int64      `json:"gas_used"`
	GasPrice       int64      `json:"gas_price"`
	ChainName      string     `json:"chain_name"`
	CreatedTime    int64      `json:"created_time"`
	BlockNumber    int64      `json:"block_number"`
	Status         int        `json:"status"`
}

func main() {
	var txHistorys = []TxHistory{
		{
			Txhash:         "hash",
			From:           "form",
			To:             "to",
			Amount:         new(big.Float),
			OriginalAmount: "0x12",
			Token:          "token1",
			GasUsed:        130,
			GasPrice:       15,
			ChainName:      "ETH",
			CreatedTime:    1567899766,
			BlockNumber:    1567,
			Status:         1,
		},
		{
			Txhash:         "hash2",
			From:           "form2",
			To:             "to2",
			Amount:         new(big.Float),
			OriginalAmount: "0x122",
			Token:          "token12",
			GasUsed:        13008,
			GasPrice:       150,
			ChainName:      "ETH",
			CreatedTime:    1567899788,
			BlockNumber:    1568,
			Status:         1,
		},
		{
			Txhash:         "hash3",
			From:           "form2",
			To:             "to2",
			Amount:         new(big.Float),
			OriginalAmount: "0x122",
			Token:          "token12",
			GasUsed:        13008,
			GasPrice:       150,
			ChainName:      "ETH",
			CreatedTime:    1567899788,
			BlockNumber:    1568,
			Status:         1,
		},
	}

	sql := "INSERT INTO \"tx_history\"(\"txhash\", \"from\", \"to\", \"amount\", \"token\", \"gas_used\", \"gas_price\", \"chain_name\", \"created_time\", \"block_number\", \"status\") VALUES "

	offTxHistory := cutOffTxHistory(txHistorys)
	fmt.Println("offTxhistory:", offTxHistory)

	for _, histories := range offTxHistory {
		var s = ""
		for i, history := range histories {
			//sprintf := fmt.Sprintf("('%s','%s','%s',%f,'%s',%d,%d,'%s',TO_TIMESTAMP(%d),%d,%d)",
			//	history.Txhash, history.From, history.To, history.Amount, history.Token, history.GasUsed,
			//	history.GasPrice, history.ChainName, history.CreatedTime, history.BlockNumber, history.Status)

			var buffer bytes.Buffer
			buffer.WriteString("('")
			buffer.WriteString(history.Txhash)
			buffer.WriteString("','")
			buffer.WriteString(history.From)
			buffer.WriteString("','")
			buffer.WriteString(history.To)
			buffer.WriteString("',")
			buffer.WriteString(history.Amount.String())
			buffer.WriteString(",'")
			buffer.WriteString(history.Token)
			buffer.WriteString("',")
			buffer.WriteString(strconv.FormatInt(history.GasUsed, 10))
			buffer.WriteString(",")
			buffer.WriteString(strconv.FormatInt(history.GasPrice, 10))
			buffer.WriteString(",'")
			buffer.WriteString(history.ChainName)
			buffer.WriteString("',")
			buffer.WriteString("TO_TIMESTAMP(")
			buffer.WriteString(strconv.FormatInt(history.CreatedTime, 10))
			buffer.WriteString("),")
			buffer.WriteString(strconv.FormatInt(history.BlockNumber, 10))
			buffer.WriteString(",")
			buffer.WriteString(strconv.FormatInt(int64(history.Status), 10))
			buffer.WriteString(")")
			buffer.WriteString(",")
			if i == 0 {
				s = buffer.String()
			} else {
				s = s + buffer.String()
			}
		}

		newSql := sql + s[0:strings.LastIndex(s, ",")]

		fmt.Println("sql:\n", newSql)
		fmt.Println("-----------------------------------------")
	}

}

const size = 3

func cutOffTxHistory(historys []TxHistory) [][]TxHistory {
	var txHiss [][]TxHistory

	mod := len(historys) % size
	k := len(historys) / size

	var end int
	if mod == 0 {
		end = k
	} else {
		end = k + 1
	}

	for i := 0; i < end; i++ {
		if i != k {
			txHiss = append(txHiss, historys[i*size:(i+1)*size])
		} else {
			txHiss = append(txHiss, historys[i*size:])
		}
	}

	return txHiss
}
