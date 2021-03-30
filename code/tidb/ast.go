package main

import (
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

type visitor struct {
}

func (v visitor) Enter(in ast.Node) (out ast.Node, skipChildren bool) {
	fmt.Println(fmt.Sprintf("----in: %+v", in))
	return in, false
}

func (v visitor) Leave(in ast.Node) (out ast.Node, ok bool) {
	fmt.Println(fmt.Sprintf("----leave: %+v", in))
	return in, true
}


/*
						sql

              /		 |				\
         	/        |        		 \
        fields      table              and
	   /  \          |            /    		  \
	id   name       user_tab    and      	   or
							    /   \	      /    \
  							 >  	!=       <      >
                            /  \    /   \    / \    /  \
 						  id   10 name  yoo age 10 age 50
*/

func main() {
	p := parser.New()
	stmtNodes, _, err := p.Parse("SELECT id, name FROM user_tab WHERE id > 10 AND name != 'yoo' AND (age < 10 OR age > 50)", "", "")
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		for _, sn := range stmtNodes {
			v := &visitor{}
			sn.Accept(v)
		}
	}
}
