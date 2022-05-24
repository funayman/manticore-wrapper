package manticore

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

type manticoreWrapper struct {
	db *sql.DB
}

func (m *manticoreWrapper) Begin() (*sql.Tx, error) {
	return nil, ErrNotImplemented
}

func (m *manticoreWrapper) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return nil, ErrNotImplemented
}

func (m *manticoreWrapper) Close() error {
	return m.db.Close()
}

func (m *manticoreWrapper) Conn(ctx context.Context) (*sql.Conn, error) {
	return m.db.Conn(ctx)
}

func (m *manticoreWrapper) Driver() driver.Driver {
	return m.db.Driver()
}

func (m *manticoreWrapper) Exec(query string, args ...any) (sql.Result, error) {
	return nil, ErrNotImplemented
}

func (m *manticoreWrapper) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return nil, ErrNotImplemented
}

func (m *manticoreWrapper) Ping() error {
	return m.db.Ping()
}

func (m *manticoreWrapper) PingContext(ctx context.Context) error {
	return m.db.PingContext(ctx)
}

func (m *manticoreWrapper) Prepare(query string) (*sql.Stmt, error) {
	return nil, ErrNotSupported
}

func (m *manticoreWrapper) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, ErrNotSupported
}

func (m *manticoreWrapper) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, ErrNotImplemented
}

func (m *manticoreWrapper) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return nil, ErrNotImplemented
}

func (m *manticoreWrapper) QueryRow(query string, args ...any) *sql.Row {
	return nil
}

func (m *manticoreWrapper) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return nil
}

func (m *manticoreWrapper) SetConnMaxIdleTime(d time.Duration) {
	m.db.SetConnMaxIdleTime(d)
}

func (m *manticoreWrapper) SetConnMaxLifetime(d time.Duration) {
	m.db.SetConnMaxLifetime(d)
}

func (m *manticoreWrapper) SetMaxIdleConns(n int) {
	m.db.SetMaxIdleConns(n)
}

func (m *manticoreWrapper) SetMaxOpenConns(n int) {
	m.db.SetMaxOpenConns(n)
}

func (m *manticoreWrapper) Stats() sql.DBStats {
	return m.db.Stats()
}

func Wrap(db *sql.DB) DB {
	return &manticoreWrapper{db: db}
}
