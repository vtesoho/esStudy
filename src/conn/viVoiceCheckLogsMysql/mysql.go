package viVoiceCheckLogsMysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() (*gorm.DB, error) {

	dsn := "root:aejo#IEBBBsNtlwr726D9%VwJ@OwKR@tcp(10.192.10.11:3306)/aichecklog?charset=utf8mb4&parseTime=True&loc=Local" //测试

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		// NamingStrategy: schema.NamingStrategy{
		// 	SingularTable: true,
		// },
	})

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)

	// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 240)

	DB = db
	// conn,err:=sqlDB.Conn(context.TODO())

	return db, err

	// DB.Select("")

}

func MysqlTabTable(tableName string) *gorm.DB {
	sql := "use " + tableName
	DB.Exec(sql)
	return DB

}

func GetMysqlDB() *gorm.DB {
	return DB
}

func ReTableName(tabname string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(tabname)
	}
}
