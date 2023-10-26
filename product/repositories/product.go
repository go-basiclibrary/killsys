package repositories

import (
	"database/sql"
	"product/common"
	"product/datamodels"
)

type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) error
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table   string
	sqlConn *sql.DB
}

func NewProductManager(table string, sqlConn *sql.DB) IProduct {
	return &ProductManager{table: table, sqlConn: sqlConn}
}

func (p *ProductManager) Conn() error {
	var err error
	if p.sqlConn == nil {
		p.sqlConn, err = common.NewMysqlConn()
	}
	return err
}

func (p *ProductManager) Insert(product *datamodels.Product) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductManager) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductManager) Update(product *datamodels.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductManager) SelectByKey(id int64) (*datamodels.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductManager) SelectAll() ([]*datamodels.Product, error) {
	//TODO implement me
	panic("implement me")
}
