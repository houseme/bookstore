package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	bookFieldNames          = builderx.RawFieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBookIdPrefix = "cache:book:id:"
)

type (
	BookModel interface {
		Insert(data Book) (sql.Result, error)
		FindOne(id int64) (*Book, error)
		Update(data Book) error
		Delete(id int64) error
	}

	defaultBookModel struct {
		sqlc.CachedConn
		table string
	}

	Book struct {
		Id    int64  `db:"id"`    // ID
		Book  string `db:"book"`  // Book name
		Price int64  `db:"price"` // book price
	}
)

func NewBookModel(conn sqlx.SqlConn, c cache.CacheConf) BookModel {
	return &defaultBookModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`book`",
	}
}

func (m *defaultBookModel) Insert(data Book) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bookRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Book, data.Price)

	return ret, err
}

func (m *defaultBookModel) FindOne(id int64) (*Book, error) {
	bookIdKey := fmt.Sprintf("%s%v", cacheBookIdPrefix, id)
	var resp Book
	err := m.QueryRow(&resp, bookIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) Update(data Book) error {
	bookIdKey := fmt.Sprintf("%s%v", cacheBookIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bookRowsWithPlaceHolder)
		return conn.Exec(query, data.Book, data.Price, data.Id)
	}, bookIdKey)
	return err
}

func (m *defaultBookModel) Delete(id int64) error {

	bookIdKey := fmt.Sprintf("%s%v", cacheBookIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bookIdKey)
	return err
}

func (m *defaultBookModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBookIdPrefix, primary)
}

func (m *defaultBookModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookRows, m.table)
	return conn.QueryRow(v, query, primary)
}
