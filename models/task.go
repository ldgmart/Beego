package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gopherjs/gopherjs/js"
	//"github.com/astaxie/beego/validation"
	"fmt"
	"time"
	"log"
)

// ** Database connect info ** // 
const (
	dbAlias       = "default"
	mysqlUser     = "root"
	mysqlPassword = "wjstks01"
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlDatabase = "test"
	mysqlCharset  = "utf8"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
		mysqlCharset,
	)
)
type User struct {
	ID    string  `orm:"pk"` //Primary Key
	Password string `orm:"size(32)"`
	Email string `orm:"size(64);unique"`
	
}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(User))
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDataBase(dbAlias, "mysql", mysqlCon)
}

// Insert Signup data
func AddDB(id string, password string, email string){
	o := orm.NewOrm()
	o.Using("default")
	var res int64
	var err error
	user := new(User)
	user.ID = id
	user.Password = password
	user.Email = email
	
	if res, err = o.Insert(user); err != nil {
	log.Println(err)
	}
	log.Printf("inserted: %d row", res)

}
func IsID(id string) (bool, error) {
	o := orm.NewOrm()
	o.Using("default")
	readTarget := User{ID: id}
	err := o.Read(&readTarget)
	if err == orm.ErrNoRows {
		return false, nil
	}else{
		return true, nil
	}
}
//Signup information will be embedded into DB 
func Signup(id string, password string, email string) (*User, error) {
	orm.Debug = true
// Sync DB if there is no databases	
/*
	force := true   // Drop table and re-create.
	verbose := true // Print log
	// generate Tables
	if err := orm.RunSyncdb(dbAlias, force, verbose); err != nil {
		log.Println(err)
	}
*/
// SignUP 할 때 기존 사용자 있는지 체크할 것, 있으면 메시지 출력 

    AddDB(id, password, email) // Insert Data into Mysql 
	beego.Info("Signup DB Success")
	return &User{id, password, email}, nil
}