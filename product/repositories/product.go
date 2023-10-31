package repositories

import (
	"database/sql"
	"log"
	"product/common"
	"product/datamodels"
	"strconv"
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

func NewProductManager(table string) IProduct {
	if table == "" {
		table = "product"
	}
	sqlConn, err := common.NewMysqlConn()
	if err != nil {
		log.Fatalf("common newMySQLConn err:%v", err)
		return nil
	}
	return &ProductManager{table: table, sqlConn: sqlConn}
}

func (p *ProductManager) Conn() error {
	var err error
	if p.sqlConn == nil {
		p.sqlConn, err = common.NewMysqlConn()
	}
	if p.table == "" {
		p.table = "product"
	}
	return err
}

func (p *ProductManager) Insert(product *datamodels.Product) (int64, error) {
	sql := "INSERT product SET product_name=?,product_num=?,product_image=?,product_url=?"
	stmt, err := p.sqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (p *ProductManager) Delete(id int64) error {
	sql := "DELETE　FROM product where id=?"
	stmt, err := p.sqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func (p *ProductManager) Update(product *datamodels.Product) error {
	sql := "UPDATE product SET product_name=?,product_num=?,product_image=?,product_url=? where id=?"
	stmt, err := p.sqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl, product.ID)
	return err
}

func (p *ProductManager) SelectByKey(id int64) (productResult *datamodels.Product, err error) {
	sql := "SELECT * FROM product WHERE id =" + strconv.FormatInt(id, 10)

	row, err := p.sqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	res := common.GetResultRow(row)
	if len(res) == 0 {
		return &datamodels.Product{}, nil
	}
	productResult = &datamodels.Product{}
	common.DataToStructByTagSql(res, productResult)

	return
}

func (p *ProductManager) SelectAll() (productArray []*datamodels.Product, err error) {
	sql := "SELECT * FROM product"
	rows, err := p.sqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	res := common.GetResultRows(rows)
	if len(res) == 0 {
		return nil, nil
	}

	for _, v := range res {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		productArray = append(productArray, product)
	}
	return
}
