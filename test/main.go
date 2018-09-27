package main

import (
	"encoding/json"
	"fmt"

	"github.com/wnote/dbcompare"
	_ "github.com/wnote/worm/dialects/mysql"
)

func main() {
	diffResult, err := dbcompare.Compare(dbcompare.CompareConfig{
		Db1Dn: "root:123456@tcp(10.5.23.82:3306)/panther_user?charset=utf8&parseTime=true&loc=Local",
		Db2Dn: "root:123456@tcp(10.5.23.82:3306)/panther_user_bak?charset=utf8&parseTime=true&loc=Local",
	})
	if err != nil {
		fmt.Println(err)
	}
	jsonResult, err := json.Marshal(diffResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonResult))
}
