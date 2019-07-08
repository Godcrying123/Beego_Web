package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"strconv"

	_ "github.com/lib/pq"
	//"github.com/Go-SQL-Driver/MySQL"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	maxIdle, _ := strconv.Atoi(beego.AppConfig.String("maxIdle"))
	maxConn, _ := strconv.Atoi(beego.AppConfig.String("maxConn"))
	if dbport == "" {
		dbport = "5432"
	}
	dsn := "user=" + dbuser + " password=" + dbpassword + " dbname=" + dbname + " host=" + dbhost + " port=" + dbport + " sslmode=disable"
	PgDataRegister(dsn, maxIdle, maxConn)
	//dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	//MySQLDataRegister(dsn, maxIdle, maxConn)
	ModelRegister()
	orm.RunSyncdb("default", false, true)
}

func PgDataRegister(dsn string, maxIdle, maxConn int) {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dsn, maxIdle, maxConn)
}

func MySQLDataRegister(dsn string, maxIdle, maxConn int) {
	orm.RegisterDataBase("default", "mysql", dsn, maxIdle, maxConn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
}

func ModelRegister() {
	orm.RegisterModel(new(Category), new(Topic), new(User), new(Profile))
}
