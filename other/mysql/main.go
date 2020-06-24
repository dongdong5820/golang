package main

import (
	"database/sql"
	"fmt"
	"time"
	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
)
// 数据建表语句
//CREATE TABLE `userinfo` (
//`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
//`name` VARCHAR(128) DEFAULT NULL,
//`age` INT(10) NOT NULL DEFAULT '0',
//`score` DOUBLE DEFAULT NULL,
//`last_modify_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
//PRIMARY KEY (`id`)
//) ENGINE=INNODB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

// 定义user结构体
type User struct {
	ID int64 `db:"id"`
	Name sql.NullString `db:"name"`
	Age int `db:"age"`
}
// 定义mysql账号密码等信息
const (
	USERNAME = "homestead"
	PASSWORD = "secret"
	NETWORK = "tcp"
	SERVER = "192.168.10.10"
	PORT = 3306
	DATABASE = "homestead"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	//fmt.Println(dsn)
	// 获取数据对象DB
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed, err: %v\n", err)
		return
	}
	// 设置db一个超时时间
	DB.SetConnMaxLifetime(100 * time.Second)
	// 设置最大打开的连接数，默认值为0表示不限制
	DB.SetMaxOpenConns(100)
	//  设置空闲的连接数
	DB.SetMaxIdleConns(16)

	// 查询单条记录
	queryOne(DB)
	// 查询多条记录
	queryMulti(DB)
	// 插入数据
	insertData(DB)
	// 修改行
	updateData(DB)
	// 删除行
	deleteData(DB)
}
// 查询单行
func queryOne(DB *sql.DB)  {
	user := new(User)
	// 单行查询。db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当Scan()方法调用之后把连接释放回到连接池
	row := DB.QueryRow("select id,name,age from userinfo where id=?", 1)
	// 如果行不存在，则scan()返回错误，需要处理异常。成功则绑定数据到结构体上
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Println(*user)
}
// 查询多行
func queryMulti(DB *sql.DB)  {
	user := new(User)
	rows, err := DB.Query("select id,name,age from userinfo where id > ?", 1)
	defer func(){
		if rows !=nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed, err: %v\n", err)
		return
	}
	// rows.Next()，循环获取所有数据
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed, err: %v\n", err)
			return
		}
		fmt.Print(*user)
	}
	fmt.Println()
}
// 插入数据
func insertData(DB *sql.DB)  {
	// exec()会将连接立马返回给连接池，但它3返回的result对象会应用该连接，
	// 数据库的insert，update采用exec方法，返回的error可知操作失败的原因，
	// 返回的ret可进一步查询本次插入数据影响的行数RowsAffected和最后插入的id-LastInsertId
	result, err := DB.Exec("INSERT INTO userinfo (`name`,`age`,`score`) VALUES (?,?,?)", "python", 28, 78)
	if err != nil {
		fmt.Printf("Insert failed, err: %v\n", err)
		return
	}
	// 最后插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get LastInsertId failed, err: %v\n", err)
		return
	}
	fmt.Println("LastInsertId: ", lastInsertId)
	// 本次插入数据影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAfected failed, err: %v\n", err)
		return
	}
	fmt.Println("RowsAffected: ", rowsAffected)
}
// 修改行
func updateData(DB *sql.DB)  {
	result, err := DB.Exec("UPDATE `userinfo` SET age=? WHERE id=?", 30, 3)
	if err != nil {
		fmt.Printf("Update failed, err: %v\n", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get rowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Println("RowsAffected: ", rowsAffected)
}
// 删除行
func deleteData(DB *sql.DB)  {
	result,err := DB.Exec("DELETE FROM `userinfo` WHERE id=?", 1)
	if err != nil {
		fmt.Printf("Delete failed, err: %v\n", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get rowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Println("RowsAffected: ", rowsAffected)
}