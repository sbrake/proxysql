package main

import (
	"strconv"
	"fmt"
	"github.com/imSQL/proxysql"
)

func main() {
	conn, err := proxysql.NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		fmt.Println(err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		fmt.Println(err)
	}

	i := new(uint64)
	for *i = 0;*i<1000 ;*i++{
		newsrv,err := proxysql.NewQr(strconv.Itoa(int(*i)),0)
		if err != nil {
			fmt.Println(err)
		}

		err = newsrv.AddOneQr(db)
		if err != nil {
			fmt.Println(err)
		}
	}

	
}
