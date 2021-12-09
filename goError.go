package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := getData()
	if err != nil {
		fmt.Println("original error==", errors.Cause(err))
		//("%+v", err) 才会打印 堆栈路径，使用 %v %s 不行。
		fmt.Printf("\nstack track== %+v", err)
		return
	}

}

func getData() error {
	err := readDB()
	if err != nil {
		// 不能再次使用Wrap，否则会打印多次堆栈路径。需要err携带上下文信息使用WithMessage
		return errors.WithMessage(err, "getData")
	}
	return nil

}

func readDB() error {
	err := sql.ErrNoRows
	if err != nil {
		// 最下层使用Wrap可以打印出更加完整的堆栈路径
		return errors.Wrap(err, "readDB")

	}
	return nil
}
