package main

import (
	"bytes"
	"fmt"

	"strconv"
)

func hex2dec(hex string) int64 {
	val := hex[2:]
	n, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		return -1
	}
	return n
}
func main() {
	addr := "address"
	var stamp int64 = 156779950
	var i int64 = 18
	// 新增旷工地址
	minerSql := "INSERT INTO \"account_address\"(\"address\", \"chain_name\", \"first_occurtime\", \"first_occurtx\", \"tags\") VALUES "
	sprintf3 := fmt.Sprintf("('%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d],ARRAY['%s'])", addr, "ETH", stamp, i, -1, "miner")

	var bf bytes.Buffer
	bf.WriteString("('")
	bf.WriteString(addr)
	bf.WriteString("','ETH',TO_TIMESTAMP(")
	bf.WriteString(strconv.FormatInt(stamp, 10))
	bf.WriteString("),ARRAY[")
	bf.WriteString(strconv.FormatInt(i, 10))
	bf.WriteString(",")
	bf.WriteString(strconv.FormatInt(-1, 10))
	bf.WriteString("],ARRAY['")
	bf.WriteString("miner'])")

	fmt.Println(minerSql + sprintf3)
	fmt.Println(minerSql + bf.String())

}
