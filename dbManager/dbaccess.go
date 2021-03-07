package dbManager

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_demo/configManager"
)

var dbCommand *sql.DB
var dbQuery *sql.DB

func initDB() (err error) {
	mysqlConfig := configManager.GetMysqlConfig()
	writeDSN := mysqlConfig.WriteDB.User + ":" + mysqlConfig.WriteDB.Password + "@(" + mysqlConfig.WriteDB.Host + ":" + mysqlConfig.WriteDB.Port + ")/" + mysqlConfig.WriteDB.DbName
	fmt.Println(writeDSN)
	//db, err := sql.Open("mysql", "root:123456@(127.0.0.1)/test")
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	dbCommand, err = sql.Open(mysqlConfig.WriteDB.DriveName, writeDSN)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = dbCommand.Ping()
	if err != nil {
		return err
	}
	readDSN := mysqlConfig.ReadDB.User + ":" + mysqlConfig.ReadDB.Password + "@(" + mysqlConfig.ReadDB.Host + ":" + mysqlConfig.ReadDB.Port + ")/" + mysqlConfig.ReadDB.DbName
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	dbQuery, err = sql.Open(mysqlConfig.ReadDB.DriveName, readDSN)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = dbQuery.Ping()
	if err != nil {
		return err
	}
	return nil
}

func getDB()  {
	if (dbCommand == nil) {
		err := initDB()
		fmt.Println(err)
	}
}


type Executor func() (sql.Result, error)

func ExecCommand(sqlString string, args ...interface{} )(sql.Result, error) {
	fmt.Println(dbCommand)
	if (dbCommand == nil) {
		err := initDB()
		fmt.Println(err)
	}
	fmt.Println("dbCommand")
	fmt.Println(dbCommand)
	return dbCommand.Exec(sqlString, args...)
}

func ToExecutor(sqlString string, args ...interface{}) Executor {
	return func() (sql.Result, error) {
		return dbCommand.Exec(sqlString, args...)
	}
}

func Execs(executors ...Executor ) error {
	getDB()
	tx, err := dbCommand.Begin()
	if err != nil {
		return err
	}

	for _, executor := range executors {
		_, err := executor()
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func ExecPre(query string, args ...interface{}) (sql.Result, error) {
	statement := query
	stmt, err := dbCommand.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}



func Query(query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(dbQuery)
	if (dbQuery == nil) {
		err := initDB()
		fmt.Println(err)
	}
	fmt.Println("dbQuery")
	fmt.Println(dbQuery)

	return dbQuery.Query(query, args...)
}

func QuerySlice(query string, f func(rows *sql.Rows) (interface{}, error), args ...interface{}) ([]interface{}, error) {
	rows, err := dbQuery.Query(query, args...)
	//defer rows.Close()

	if err != nil {
		return nil, err
	}

	rs := make([]interface{}, 0)

	for rows.Next() {
		r, err := f(rows)
		if err != nil {
			return nil, err
		}

		rs = append(rs, r)
	}
	return rs, nil
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return dbQuery.QueryRow(query, args...)
}

func Single(query string, dest ...interface{}) error{
	return dbQuery.QueryRow(query).Scan(dest...)
}

func Find(query string, id interface{}, dest ...interface{}) error {
	return dbQuery.QueryRow(query, id).Scan(dest...)
}

func FindBy(query string, args []interface{}, dest ...interface{}) error{
	return dbQuery.QueryRow(query, args).Scan(dest...)
}