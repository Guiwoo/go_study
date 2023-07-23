package adapter

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

/**
기존 시스템 인터페이스 vs 업체에서 제공하는 인터페이스 가 다르다...
*/

type Duck interface {
	quack()
	fly()
}

type MallardDuck struct {
}

func (m *MallardDuck) quack() {
	fmt.Println("quakckkkckkckck")
}
func (m *MallardDuck) fly() {
	fmt.Println("오리는 날고있습니다.")
}

type Turkey interface {
	gobble()
	fly()
}
type WildTurkey struct {
}

func (w *WildTurkey) gobble() {
	fmt.Println("goooooooblelbelbel")
}
func (w *WildTurkey) fly() {
	fmt.Println("짤븡날개로 못나는디;;;")
}

type TurkeyAdapter struct {
	t Turkey
}

func (t *TurkeyAdapter) quack() {
	t.t.gobble()
}
func (t *TurkeyAdapter) fly() {
	t.t.fly()
}

/**
어댑터 패턴은 특정 클래스 인터페이스를 클라이언트에서 요구하는 다른 인터페이스로 변환 합니다. 인터페이스가 호환되지않아 같이 쓸 수 없었던 클래스를 사용할 수 있게 도와준다.
*/

type Iterator interface {
	hasNext()
	next()
	remove()
}
type Enumeration interface {
	hasMoreElements()
	nextElement()
}

type EnumerationIterator struct {
	i Iterator
}

func (e *EnumerationIterator) hasMoreElements() {
	e.i.hasNext()
}
func (e *EnumerationIterator) nextElement() {
	e.i.next()
}
func (e *EnumerationIterator) remove() {
	log.Errorf("unsupported error")
}

type PostgreSQL interface {
	InsertColumn()
	DeleteColumn()
	UpdateColumn()
	ReadColumn()
}

type PostgreSQL_V15 struct {
	db string
}

func (p PostgreSQL_V15) InsertColumn() {
	fmt.Println("Insert by postgresql version 15")
}

func (p PostgreSQL_V15) DeleteColumn() {
	fmt.Println("Delete by posgresql version 15")
}

func (p PostgreSQL_V15) UpdateColumn() {
	fmt.Println("Update by postgresql version 15")
}

func (p PostgreSQL_V15) ReadColumn() {
	fmt.Println("Read by postgresql version 15")
}

var _ PostgreSQL = (*PostgreSQL_V15)(nil)

func NewPostgreSQL() PostgreSQL {
	return &PostgreSQL_V15{"postgresql version 15"}
}

type MySQL interface {
	InsertData()
	DeleteData()
	UpdateData()
	ReadData()
}

type MySQL_V8 struct {
	db string
}

func (m MySQL_V8) InsertData() {
	fmt.Println("Insert by mysql version 8")
}

func (m MySQL_V8) DeleteData() {
	fmt.Println("Delete by mysql version 8")
}

func (m MySQL_V8) UpdateData() {
	fmt.Println("Update by mysql version 8")
}

func (m MySQL_V8) ReadData() {
	fmt.Println("Read  by mysql version 8")
}

var _ MySQL = (*MySQL_V8)(nil)

func NewMySQL() MySQL {
	return &MySQL_V8{"mysql version 8"}
}

type DbBatch interface {
	BatchInsert()
	BatchDelete()
	BatchUpdate()
	BatchRead()
}

type MySQL_V8_Batch_Adapter struct {
	mysql MySQL
}

func (m MySQL_V8_Batch_Adapter) BatchInsert() {
	m.mysql.InsertData()
}

func (m MySQL_V8_Batch_Adapter) BatchDelete() {
	m.mysql.DeleteData()
}

func (m MySQL_V8_Batch_Adapter) BatchUpdate() {
	m.mysql.UpdateData()
}

func (m MySQL_V8_Batch_Adapter) BatchRead() {
	m.mysql.ReadData()
}

var _ DbBatch = (*MySQL_V8_Batch_Adapter)(nil)

func NewDbBatchMySQLAdapter() DbBatch {
	mysql := NewMySQL()
	return &MySQL_V8_Batch_Adapter{mysql}
}

type PostgreSQL_V15_Batch_Adapter struct {
	postgres PostgreSQL
}

func (p PostgreSQL_V15_Batch_Adapter) BatchInsert() {
	p.postgres.InsertColumn()
}

func (p PostgreSQL_V15_Batch_Adapter) BatchDelete() {
	p.postgres.DeleteColumn()
}

func (p PostgreSQL_V15_Batch_Adapter) BatchUpdate() {
	p.postgres.UpdateColumn()
}

func (p PostgreSQL_V15_Batch_Adapter) BatchRead() {
	p.postgres.ReadColumn()
}

var _ DbBatch = (*PostgreSQL_V15_Batch_Adapter)(nil)

func NewDbBatchPostgreSQLAdapter() DbBatch {
	postgres := NewPostgreSQL()
	return &PostgreSQL_V15_Batch_Adapter{postgres}
}
