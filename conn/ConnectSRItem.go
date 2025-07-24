/*
 *@author  chengkenli
 *@project srload
 *@package conn
 *@file    ConnnectSRItem
 *@date    2025/7/11 15:52
 */

package conn

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func StarItem(user, password, host string) (*gorm.DB, error) {
	newLogger := logger.New(nil,
		logger.Config{
			SlowThreshold: time.Second * 1000, // 控制慢SQL阈值
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?parseTime=true&charset=utf8mb4&loc=Local",
		user,
		password,
		host,
		9030,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("Error >:%s", err.Error()))
		return nil, err
	}
	return db, err
}

func StarItemldap(user, password, host string) (*gorm.DB, error) {
	newLogger := logger.New(nil,
		logger.Config{
			SlowThreshold: time.Second * 1000, // 控制慢SQL阈值
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local&tls=true&allowCleartextPasswords=true&allowAllFiles=true&rewriteBatchedStatements=true",
		user,
		password,
		host,
		9030)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("Error >:%s", err.Error()))
		return nil, err
	}
	return db, err
}
