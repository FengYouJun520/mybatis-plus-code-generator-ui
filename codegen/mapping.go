package codegen

import (
	"encoding/json"

	"gorm.io/gorm"
)

const (
	Db_Varchar   = "varchar"
	Db_Char      = "char"
	Db_Longtext  = "longtext"
	Db_Tinyint   = "tinyint"
	Db_Int       = "int"
	Db_Integer   = "integer"
	Db_Bigint    = "bigint"
	Db_Decimal   = "decimal"
	Db_Float     = "float"
	Db_Double    = "double"
	Db_Timestamp = "timestamp"
	Db_Date      = "date"
	Db_Datetime  = "datetime"
)

// java类型的包路径
const (
	BooleanPackage       = "java.lang.Boolean"
	StringPackage        = "java.lang.String"
	IntegerPackage       = "java.lang.Integer"
	LongPath             = "java.lang.Long"
	BigDecimalPackage    = "java.lang.math.BigDecimal"
	FloatPackage         = "java.lang.Float"
	DoublePackage        = "java.lang.Double"
	LocalDatePackage     = "java.lang.LocalDate"
	LocalDateTimePackage = "java.lang.LocalDateTime"
	SerializablePackage  = "java.io.Serializable"
	TableNamePackage     = "com.baomidou.mybatisplus.annotation.TableName"
	TableIdPackage       = "com.baomidou.mybatisplus.annotation.TableId"
	TableFieldPackage    = "com.baomidou.mybatisplus.annotation.TableField"
	TableLogicPackage    = "com.baomidou.mybatisplus.annotation.TableLogic"
	VersionPackage       = "com.baomidou.mybatisplus.annotation.Version"
	IServicePackage      = "com.baomidou.mybatisplus.extension.service.IService"
	ServiceImplPackage   = "com.baomidou.mybatisplus.extension.service.impl.ServiceImpl"
	BaseMapperPackage    = "com.baomidou.mybatisplus.core.mapper.BaseMapper"
	FieldFillPackage     = "com.baomidou.mybatisplus.annotation.FieldFill"
	IdTypePackage        = "com.baomidou.mybatisplus.annotation.IdType"
)

const (
	FieldFill_DEFAULT       = "DEFAULT"
	FieldFill_INSERT        = "INSERT"
	FieldFill_UPDATE        = "UPDATE"
	FieldFill_INSERT_UPDATE = "INSERT_UPDATE"
)

const (
	Java_String        = "String"
	Java_Boolean       = "Boolean"
	Java_Integer       = "Integer"
	Java_Long          = "Long"
	Java_BigDecimal    = "BigDecimal"
	Java_Float         = "Float"
	Java_Double        = "Double"
	Java_LocalDate     = "LocalDate"
	Java_LocalDateTime = "LocalDateTime"
)

var (
	TableQuerySql = `SELECT
	table_name 'name',
	IFNULL( TABLE_COMMENT, table_name ) 'comment',
	TABLE_SCHEMA 'table_schema'
FROM
	INFORMATION_SCHEMA.TABLES 
WHERE
	UPPER( table_type )= 'BASE TABLE' 
	AND LOWER( table_schema ) = ? 
	AND table_name IN (?);
`
	// 查询列信息sql
	ColumnQuerySql = `SELECT
	COLUMN_NAME 'name',
	column_comment 'comment',
	DATA_TYPE 'data_type',
	IS_NULLABLE 'is_nullable',
	IFNULL( CHARACTER_MAXIMUM_LENGTH, 0 ) 'length',
    COLUMN_KEY 'key_flag' 
FROM
	information_schema.COLUMNS 
WHERE
	table_schema = ? 
	AND table_name = ?;
`
)

// TypeMappings 数据库和java类型的映射集合
var TypeMappings = map[string]string{
	Db_Varchar:   Java_String,
	Db_Char:      Java_String,
	Db_Longtext:  Java_String,
	Db_Tinyint:   Java_Boolean,
	Db_Int:       Java_Integer,
	Db_Integer:   Java_Integer,
	Db_Bigint:    Java_Long,
	Db_Decimal:   Java_BigDecimal,
	Db_Float:     Java_Float,
	Db_Double:    Java_Double,
	Db_Timestamp: Java_Long,
	Db_Date:      Java_LocalDate,
	Db_Datetime:  Java_LocalDateTime,
}

var TypePackages = map[string]string{
	Java_String:        StringPackage,
	Java_Boolean:       BooleanPackage,
	Java_Integer:       IntegerPackage,
	Java_Long:          LongPath,
	Java_BigDecimal:    BigDecimalPackage,
	Java_Float:         FloatPackage,
	Java_Double:        DoublePackage,
	Java_LocalDate:     LocalDatePackage,
	Java_LocalDateTime: LocalDateTimePackage,
}

// TableInfo 数据库表信息
type TableInfo struct {
	Name       string    // 表名
	Comment    string    // 描述
	SchemaName string    // 表所属数据库
	Columns    []*Column // 表所有列信息
}

// Column 列信息
type Column struct {
	Name       string // 列名
	Comment    string // 描述
	DataType   string // 数据库类型
	IsNullable string // 是否可为空
	Length     int    // 类型长度
	KeyFlag    string // 是否是主键PRI
}

func NewTableInfo(name, tableSchema string) *TableInfo {
	return &TableInfo{
		Name:       name,
		SchemaName: tableSchema,
	}
}

func (t *TableInfo) ExecuteColumns(db *gorm.DB) error {
	db = db.Raw(ColumnQuerySql, t.SchemaName, t.Name)
	if db.Error != nil {
		return db.Error
	}
	var columns []*Column

	rows, _ := db.Rows()
	defer rows.Close()

	for rows.Next() {
		column := Column{}
		db.ScanRows(rows, &column)
		columns = append(columns, &column)
	}

	t.Columns = columns
	return nil
}

func (t *TableInfo) String() string {
	data, _ := json.MarshalIndent(t, "", "  ")
	return string(data)
}
