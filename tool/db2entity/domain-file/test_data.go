package domainfile

const (
	target = "example"

	database = "test"

	testStructName = "Test"

	testTable = "test"

	userStructName = "User"

	userTable = "user"

	boubctx = "boubctx"

	delField = "is_del"
)

var (
	Cols = make([]Column, 0)
)

func init() {
	col1 := Column{
		ColumnName:    "user_name",
		DataType:      "varchar",
		IsNullAble:    yesNull,
		ColumnComment: "user name",
	}
	Cols = append(Cols, col1)

	col2 := Column{
		ColumnName: "id",
		ColumnKey:  pri,
		DataType:   "int",
		IsNullAble: noNull,
	}
	Cols = append(Cols, col2)

	col3 := Column{
		ColumnName: "update_time",
		DataType:   "timestamp",
		IsNullAble: noNull,
		Extra:      upCurrentTimestamp,
	}
	Cols = append(Cols, col3)
}
