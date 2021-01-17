package easy_orm

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)



var EzOrm *EasyOrm

type EasyOrm struct {
	Conn     string
	Db       *gorm.DB
	Enforcer *casbin.Enforcer
	*Config
}

type Config struct {
	GormConfig        *gorm.Config
	Adapter           string        // 类型
	Name              string        // 名称
	Username          string        // 用户名
	Pwd               string        // 密码
	Host              string        // 地址
	Port              int64         // 端口
	CasbinModelPath   string        // casbin 模型规则路径
	SqlitePath        string        // sqlite 模型路径
	Models            []interface{} // 模型数据
	Debug             bool          // 调试
	CasbinTablePrefix string
}

func Init(c *Config) {
	EzOrm = new(EasyOrm)
	EzOrm.Config = &Config{
		GormConfig: &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "tb_",
				SingularTable: false,
			},
		},
		Adapter:           "mysql", // 类型
		Name:              "",
		Username:          "root",      // 用户名
		Pwd:               "",          // 密码
		Host:              "127.0.0.1", // 地址
		Port:              3306,        // 端口
		CasbinModelPath:   "",          // casbin 模型规则路径
		Models:            nil,
		CasbinTablePrefix: "",
	}
	if c != nil {
		EzOrm.Config = c
	}

}

func (db *EasyOrm) getConn() string {
	if db.Config.Adapter == "mysql" {
		return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", db.Config.Username, db.Config.Pwd, db.Config.Host, db.Config.Port, db.Config.Name)
	} else if db.Config.Adapter == "postgres" {
		return fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", db.Config.Username, db.Config.Pwd, db.Config.Host, db.Config.Name)
	} else if db.Config.Adapter == "sqlite3" {
		return filepath.Join(cwd(), db.Config.Name) + ".db"
	} else {
		fmt.Println(errors.New("not supported database adapter"))
	}

	return ""
}

// getGormDb
func (db *EasyOrm) getGormDb() error {
	var err error
	var dialector gorm.Dialector
	if db.Config.Adapter == "mysql" {
		dialector = mysql.Open(db.getConn())
	} else if db.Config.Adapter == "postgres" {
		dialector = postgres.Open(db.getConn())
	} else if db.Config.Adapter == "sqlite3" {
		dialector = sqlite.Open(db.getConn())
	} else {
		fmt.Println(errors.New("not supported database adapter"))
	}

	if db.Config.Debug {
		fmt.Println(fmt.Sprintf("Conn: : %s", db.getConn()))
	}

	//&gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   db.Config.Prefix, // 表名前缀，`User` 的表名应该是 `t_users`
	//		SingularTable: false,            // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
	//	},
	//}
	db.Db, err = gorm.Open(dialector, db.Config.GormConfig)
	if err != nil {
		return err
	}

	err = db.Db.Use(
		dbresolver.Register(
			dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	if err != nil {
		return err
	}
	db.Db.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})
	return nil
}

// getEnforcer
func (db *EasyOrm) getEnforcer() error {

	c, err := gormadapter.NewAdapterByDBUseTableName(db.Db, db.Config.CasbinTablePrefix, "casbin_rule") // Your driver and data source.
	if err != nil {
		return err
	}

	db.Enforcer, err = casbin.NewEnforcer(db.Config.CasbinModelPath, c)
	if err != nil {
		return err
	}

	_ = db.Enforcer.LoadPolicy()

	return nil
}

// cwd 获取项目路径
func cwd() string {
	// 兼容 travis 集成测试
	if os.Getenv("TRAVIS_BUILD_DIR") != "" {
		return os.Getenv("TRAVIS_BUILD_DIR")
	}

	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}
