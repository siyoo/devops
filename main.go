package main

import (
	"fmt"
	"github.com/pingcap/tidb/parser"
)

func main()  {
	p := parser.New()
	stmtNodes, err := p.Parse("SELECT id FROM user_tab WHERE id > 10 AND name != 'yoo' AND (age < 10 OR age < 50)", "", "")
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		for _, sn := range stmtNodes {
			fmt.Println(sn.Text())
		}
	}
}
