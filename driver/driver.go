package driver

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/funayman/manticore"
	"github.com/go-sql-driver/mysql"
)

func init() {
	sql.Register("manticore", &Driver{})
}

type Driver struct {
	baseDriver *mysql.MySQLDriver
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	d.baseDriver = &mysql.MySQLDriver{}
	return d.baseDriver.Open(name)
}

type Conn struct {
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return nil, manticore.ErrNotSupported
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, manticore.ErrNotImplemented
}
