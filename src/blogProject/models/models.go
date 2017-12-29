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
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string    `orm:"size(5000)"`
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	Mark            string `orm:"index"`
	Class           string
	//attachment      string
}

type Respone struct {
	Id              int64
	TopicId         int64
	RpName          string
	Title           string
	Content         string    `orm:"size(512)"`
	Created         time.Time `orm:"index"`
	ReplyLastUserId int64
}

type User struct {
	Id           int64
	UserName     string
	UserPwd      string
	UserImageURL string `orm:"size(512)"`
	UserBirthday string
	UserEmail    string
	UserPhone    string
	UserAddr     string `orm:"size(512)"`
	SuperUser    bool
}

func init() {
	// 开启 orm 调试模式
	orm.Debug = true
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
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Respone))
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

//Topics Operations just like "add" "del" "mod" ...
func TopicOps(op string, p *Topic) error {
	var err error
	o := orm.NewOrm()

	ptmp := &Topic{
	//Id: p.Id,
	//Title:   p.Title,
	//Content: p.Content,
	//Updated: time.Now(),
	//ReplyTime: time.Now(),
	//Created:   time.Now(),
	}
	//TODO:使用ORM接口插入失败
	//DONE:使用ORM接口失败问题已解决：
	//因为数据库建表时有些字段明确不能为空而插入数据为空
	//_, err := o.Insert(tp)

	switch op {
	case "add":
		ptmp.Created = time.Now()
		ptmp.ReplyTime = time.Now()
		ptmp.Title = p.Title
		ptmp.Content = p.Content
		ptmp.Updated = time.Now()
		_, err = o.Insert(ptmp)
	case "mod":
		_, err = o.Update(p)

		//id, err = strconv.ParseInt(tid, 10, 64)
		//if err != nil {
		//	beego.Error(err)
		//}
		//_, err = o.Raw("UPDATE Topic SET Title=? Content=? Updated=? WHERE Id=?", title, content, time.Now(), id).Exec()
	}

	if err != nil {
		beego.Error(err)
	}

	return err
}

//Category Operations just like "add" "del" "mod" ...
func CategoryOps(op string, p *Category) error {
	var err error
	o := orm.NewOrm()

	ptmp := &Category{
		Title: p.Title,
	}

	switch op {
	case "add":
		ptmp.Created = time.Now()
		_, err = o.Insert(ptmp)
	case "mod":
		_, err = o.Update(ptmp)
	case "del":
		ptmp.Id = p.Id
		_, err = o.Delete(ptmp)

	}

	if err != nil {
		beego.Error(err)
	}

	return err
}

/*
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}

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
}*/

/*
func DelCategory(cid string) error {
	id, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	o := orm.NewOrm()
	cate := &Category{Id: id}
	_, err = o.Delete(cate)
	if err != nil {
		beego.Error(err)
	}

	return err
}*/

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func GetAllTopics(orderByTimeDec bool) ([]*Topic, error) {
	o := orm.NewOrm()
	tps := make([]*Topic, 0)
	qs := o.QueryTable("Topic")
	_, err := qs.All(&tps)
	return tps, err
}

func GetAllReplysById(tid string) ([]*Respone, error) {
	o := orm.NewOrm()
	rps := make([]*Respone, 0)
	qs := o.QueryTable("Respone")
	_, err := qs.Filter("TopicId", tid).All(&rps)

	return rps, err
}

func GetTopicById(tid string) (*Topic, error) {
	o := orm.NewOrm()
	tp := new(Topic)
	qs := o.QueryTable("Topic")
	err := qs.Filter("id", tid).One(tp)

	//exec Raw SQL
	//var id int64
	//id, err = strconv.ParseInt(tid, 10, 64)
	tp.Views++
	o.Update(tp, "Views")
	//_, err = o.Raw("UPDATE Topic SET Views=? WHERE Id=?", tp.Views, id).Exec()
	if err != nil {
		clog.Clogv(clog.Red, "Raw SQL exec fail!")
		beego.Error(err)
	}

	return tp, err
}

func ResponeOps(op string, rp *Respone) error {
	var err error
	o := orm.NewOrm()
	ptmp := &Respone{
		TopicId:         rp.TopicId,
		RpName:          rp.RpName,
		Title:           rp.Title,
		Content:         rp.Content,
		ReplyLastUserId: rp.ReplyLastUserId,
	}

	switch op {
	case "add":
		ptmp.Created = time.Now()
		//_, err = o.Raw("INSERT INTO Respone VALUES(?,?,?,?,?,?,?)", nil, rp.TopicId, rp.RpName, rp.Title, rp.Content, time.Now(), rp.ReplyLastUserId).Exec()
		_, err = o.Insert(ptmp)
	case "del":
		_, err = o.Delete(ptmp)
	}
	if err != nil {
		beego.Error(err)
	}

	return err
}

func IsAccountExist(name string) bool {
	o := orm.NewOrm()
	pu := new(User)
	qs := o.QueryTable("User")
	err := qs.Filter("UserName", name).One(pu)

	if err != nil {
		beego.Error(err)
		return false
	}
	return name == pu.UserName
}

//User Operations just like "add" "del" "mod" ...
func UserOps(op string, pu *User) error {
	var err error
	o := orm.NewOrm()

	switch op {
	case "add":
		//_, err = o.Raw("INSERT INTO User VALUES(?,?,?,?,?,?,?,?,?)", nil, pu.UserName, pu.UserPwd, pu.UserImageURL, pu.UserBirthday, pu.UserEmail, pu.UserPhone, pu.UserAddr, pu.SuperUser).Exec()

		//TODO:使用ORM接口插入失败
		_, err = o.Insert(pu)

	case "del":
	case "mod":
	}

	if err != nil {
		beego.Error(err)
	}
	return err
}
