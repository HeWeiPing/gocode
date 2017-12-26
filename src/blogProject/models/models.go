package models

import (
	"clog"
	//	"database/sql"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
	"time"
)

type DBcfg struct {
	db_dir     string
	db_type    string
	db_host    string
	db_port    string
	db_user    string
	db_pass    string
	db_name    string
	db_maxIdle int
	db_maxConn int
}

var dbCfg *DBcfg

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func init() {
	dbCfg = new(DBcfg)
	dbCfg.db_dir = beego.AppConfig.String("db_dir")
	dbCfg.db_type = beego.AppConfig.String("db_type")
	dbCfg.db_host = beego.AppConfig.String("db_host")
	dbCfg.db_port = beego.AppConfig.String("db_port")
	dbCfg.db_user = beego.AppConfig.String("db_user")
	dbCfg.db_pass = beego.AppConfig.String("db_pass")
	dbCfg.db_name = beego.AppConfig.String("db_name")
	dbCfg.db_maxIdle, _ = strconv.Atoi(beego.AppConfig.String("db_maxIdle"))
	dbCfg.db_maxConn, _ = strconv.Atoi(beego.AppConfig.String("db_maxConn"))

	// 需要在init中注册定义的model
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Topic))
}

// Register database
func RegisterDB() error {
	var err error
	switch dbCfg.db_type {
	case "mysql":
		err = RegDbMySQL()
	case "sqlite3":
		err = RegDbSqlite()
	}
	return err
}

//MySQL
func RegDbMySQL() error {
	dns := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbCfg.db_user, dbCfg.db_pass, dbCfg.db_name)
	clog.Clogv(clog.Green, "DB dns=%s", dns)
	err := orm.RegisterDriver(dbCfg.db_name, orm.DRMySQL)
	if err != nil {
		beego.Error(err)
		return err
	}
	err = orm.RegisterDataBase("default", dbCfg.db_type, dns, dbCfg.db_maxIdle, dbCfg.db_maxConn)
	if err != nil {
		beego.Error(err)
		return err
	}

	return nil
}

//sqlite
func RegDbSqlite() error {
	whereDB := dbCfg.db_dir + dbCfg.db_name
	clog.Clogv(clog.Yellow, "db_dir=%s", dbCfg.db_dir)
	clog.Clogv(clog.Yellow, "where database:%s", whereDB)
	if !com.IsExist(whereDB) {
		os.MkdirAll(dbCfg.db_dir, os.ModePerm)
		os.Create(whereDB)
	}

	err := orm.RegisterDriver(dbCfg.db_type, orm.DRSqlite)
	if err != nil {
		beego.Error(err)
		return err
	}
	err = orm.RegisterDataBase(
		"default",
		dbCfg.db_type,
		whereDB,
		dbCfg.db_maxIdle,
		dbCfg.db_maxConn)
	if err != nil {
		beego.Error(err)
		return err
	}

	return nil
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}
	clog.Clogv(clog.Green, "name=%s", name)

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	//TODO:alway error
	//_, err = o.Insert(cate)
	//if err != nil {
	//	beego.Error(err)
	//	return err
	//}

	_, err = o.Raw("INSERT INTO category VALUES(?,?,?,?,?,?,?)", nil, name, 0, 0, 0, 0, 0).Exec()
	if err != nil {
		clog.Clogv(clog.Red, "Raw SQL exec fail!")
		beego.Error(err)
	}

	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
