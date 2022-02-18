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
	sql := "INSERT INTO \"account_address\"(\"address\", \"chain_name\", \"first_occurtime\", \"first_occurtx\") VALUES "

	from := "from"
	to := "to"
	var stamp int64 = 1567788990
	i := 5
	index := "0x5"

	sprintf1 := fmt.Sprintf("('%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d])", from, "ETH", stamp, i, hex2dec(index))
	sprintf2 := fmt.Sprintf("('%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d])", to, "ETH", stamp, i, hex2dec(index))

	var bf bytes.Buffer
	bf.WriteString("('")
	bf.WriteString(from)
	bf.WriteString("','")
	bf.WriteString("ETH',TO_TIMESTAMP(")
	bf.WriteString(strconv.FormatInt(stamp, 10))
	bf.WriteString("),ARRAY[")
	bf.WriteString(strconv.FormatInt(int64(i), 10))
	bf.WriteString(",")
	bf.WriteString(strconv.FormatInt(hex2dec(index), 10))
	bf.WriteString("])")

	var bf2 bytes.Buffer
	bf2.WriteString("('")
	bf2.WriteString(to)
	bf2.WriteString("','")
	bf2.WriteString("ETH',TO_TIMESTAMP(")
	bf2.WriteString(strconv.FormatInt(stamp, 10))
	bf2.WriteString("),ARRAY[")
	bf2.WriteString(strconv.FormatInt(int64(i), 10))
	bf2.WriteString(",")
	bf2.WriteString(strconv.FormatInt(hex2dec(index), 10))
	bf2.WriteString("])")

	fmt.Println(sql + sprintf1)
	fmt.Println(sql + bf.String())
	fmt.Println("--------------------------------")

	fmt.Println(sql + sprintf2)
	fmt.Println(sql + bf2.String())

}
