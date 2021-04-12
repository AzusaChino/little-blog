package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

// RestResponse 前端返回封装类
type RestResponse struct {
	Code    uint8
	Data    interface{}
	Message string
}

func Ok(obj interface{}) *RestResponse {
	return &RestResponse{
		Code:    http.StatusOK,
		Data:    obj,
		Message: "ok",
	}
}

func Error(code uint8, msg string) *RestResponse {
	return &RestResponse{Code: code, Message: msg}
}

// DsnUrl 数据库连接地址: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
const DsnUrl = "az:azusa520@tcp(180.76.169.35:3306)/little_blog?charset=utf8mb4&parseTime=True&loc=Local"

// GetDb 获取数据库连接
func GetDb() (*gorm.DB, func(), error) {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DsnUrl, // DSN data source name
		DefaultStringSize:         256,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	// clean db connection for defer
	cleanFunc := func() {
		sqlDb, err := db.DB()
		if sqlDb != nil {
			_ = sqlDb.Close()
		}
		if err != nil {
			panic(err)
		}
	}

	return db, cleanFunc, nil
}
