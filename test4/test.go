package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	sql := "INSERT INTO \"contract_info\"(\"address\", \"creator\", \"chain_name\", \"created_time\", \"created_txhash\", \"contract_creation_code\", \"nonce\") VALUES "
	sprintf := fmt.Sprintf("('%s','%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d],'%s','%s')", "address", "creator", "eth", 156789, 12, 2, "code", "0x001")

	var bf bytes.Buffer
	bf.WriteString("('")
	bf.WriteString("address")
	bf.WriteString("','")
	bf.WriteString("creator")
	bf.WriteString("','")
	bf.WriteString("eth")
	bf.WriteString("',TO_TIMESTAMP(")
	bf.WriteString(strconv.FormatInt(156789, 10))
	bf.WriteString("),ARRAY[")
	bf.WriteString(strconv.FormatInt(12, 10))
	bf.WriteString(",")
	bf.WriteString(strconv.FormatInt(2, 10))
	bf.WriteString("],'")
	bf.WriteString("code")
	bf.WriteString("','")
	bf.WriteString("0x001")
	bf.WriteString("')")

	fmt.Println(sql + sprintf)
	fmt.Println(sql + bf.String())

	fmt.Println((sql + sprintf) == (sql + bf.String()))

}
