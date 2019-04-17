package sql

import (
	"fmt"
	"testing"
)

func TestQueryBatchIDSQL(t *testing.T) {
	sql := QueryBatchIDSQL([]string{"624519", "990201"}, "mid")
	fmt.Println(sql)
}

func TestQueryBatchStudentSQL(t *testing.T) {
	sql := QueryBatchStudentSQL([]string{"624519", "990201"})
	fmt.Println(sql)
}
