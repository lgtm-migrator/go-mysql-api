package mysql

import (
	"gopkg.in/doug-martin/goqu.v4"
)

// SQL return sqls by sql builder
type SQL struct {
	sqlBuilder *goqu.Database
	dbMeta     *DataBaseMetadata
}

func (s *SQL) getPriKeyNameOf(tableName string) string {
	return s.dbMeta.GetTableMeta(tableName).GetPrimaryColumn().ColumnName
}

// GetByTable with filter
func (s *SQL) GetByTable(tableName string, mWhere map[string]interface{}, opt QueryOption) (sql string, err error) {
	builder := s.sqlBuilder.From(tableName)
	for k, v := range mWhere {
		builder = builder.Where(goqu.Ex{k: v})
	}
	builder = configBuilder(builder, opt)
	sql, _, err = builder.ToSql()
	return
}

// GetByTableAndID for specific record in table
func (s *SQL) GetByTableAndID(tableName string, id interface{}, opt QueryOption) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(tableName)
	builder := s.sqlBuilder.From(tableName).Where(goqu.Ex{priKeyName: id})
	builder = configBuilder(builder, opt)
	sql, _, err = builder.ToSql()
	return sql, err
}

// UpdateByTable for update specific record by id
func (s *SQL) UpdateByTable(tableName string, record map[string]interface{}) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(tableName)
	builder := s.sqlBuilder.From(tableName).Where(goqu.Ex{priKeyName: record[priKeyName]})
	sql, _, err = builder.ToUpdateSql(record)
	return
}

// InsertByTable and record map
func (s *SQL) InsertByTable(tableName string, record map[string]interface{}) (sql string, err error) {
	sql = goqu.From(tableName).Insert(record).Sql
	return
}

// DeleteByTable by where
func (s *SQL) DeleteByTable(tableName string, mWhere map[string]interface{}) (sql string, err error) {
	builder := goqu.From(tableName)
	for k, v := range mWhere {
		builder = builder.Where(goqu.Ex{k: v})
	}
	sql = builder.Delete().Sql
	return
}

func configBuilder(builder *goqu.Dataset, opt QueryOption) (rs *goqu.Dataset) {
	rs = builder
	if opt.limit != 0 {
		rs = rs.Limit(uint(opt.limit))
	}
	if opt.offset != 0 {
		rs = rs.Offset(uint(opt.offset))
	}
	if opt.fields != nil {
		rs = rs.Select(opt.fields...)
	}
	return
}