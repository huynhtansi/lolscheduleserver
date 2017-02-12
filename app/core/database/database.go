package database

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/huynhtansi/lolscheduleserver/app/config"
)

var SQL *gorm.DB

func Connect(conf config.DatabaseConfig) error {
	db, err := DatabaseConfiguration(conf)
	if err != nil {
		return err
	}
	SQL = db
	return nil
}

func DatabaseConfiguration(conf config.DatabaseConfig) (*gorm.DB, error)  {
	dsn := ""	//Data Source Name
	switch conf.DriverName {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Name)
	case "postgres":
		dsn = fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable ", conf.Host, conf.Name, conf.Username ,conf.Password )
	default:
		return nil, fmt.Errorf("ORM：%s", conf.DriverName)
	}
	fmt.Println(dsn)
	db, err := gorm.Open(conf.DriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("Fail to connect to database %v", err)
	}
	//defer db.Close()
	return db, nil
}