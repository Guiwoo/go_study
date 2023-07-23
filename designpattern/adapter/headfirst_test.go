package adapter

import (
	"fmt"
	"testing"
)

func duckTest(duck Duck) {
	duck.quack()
	duck.fly()
}

func Test_01(t *testing.T) {
	var d Duck
	duck := &MallardDuck{}

	w := &WildTurkey{}
	d = &TurkeyAdapter{w}

	duckTest(duck)

	duckTest(d)

}

func Test_DB(t *testing.T) {
	batch := []DbBatch{NewDbBatchMySQLAdapter(), NewDbBatchPostgreSQLAdapter()}

	for _, v := range batch {
		v.BatchInsert()
		v.BatchDelete()
		v.BatchRead()
		v.BatchUpdate()
		fmt.Println("----------------------------------")
	}
}
