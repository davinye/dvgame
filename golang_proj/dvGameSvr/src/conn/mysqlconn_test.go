package conn

import (
	"testing"
)

func Test_MysqlTestMain(t *testing.T) {
	num := MysqlTestMain(1)
	t.Log(num)
}
