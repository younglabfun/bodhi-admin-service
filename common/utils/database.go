package utils

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"strings"
	"time"
)

func InitMySQL(host, user, password, database string, port int64) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:         "",
			SingularTable:       true,
			NameReplacer:        nil,
			NoLowerCase:         false,
			IdentifierMaxLength: 0,
		},
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		logx.Errorf("Init MySQL connect err: %v", err)
		return nil, err
	}
	return db, nil
}

func CheckMySQLConf(host, user, password string, port int64) (*sql.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/information_schema",
		user, password, host, port)
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ClearTables(db *gorm.DB, tables []string) {
	for _, table := range tables {
		sql := fmt.Sprintf("DROP TABLE %s", table)
		db.Exec(sql)
	}
	fmt.Println("drop all tables")
}

func RunSql(db *gorm.DB, sqlFile string) error {
	_, err := os.Stat(sqlFile)
	if os.IsNotExist(err) {
		return err
	}

	f, _ := os.ReadFile(sqlFile)
	sqls := strings.Split(string(f), ";")
	for _, sql := range sqls {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("run sql ", len(sqls))
	return nil
}
