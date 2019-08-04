package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"errors"
)

//Signup information will be embedded into DB 
func Login(id string, password string) ( *User, error) {
	orm.Debug = true
	msg := "No result found"
	o := orm.NewOrm()
	o.Using("default")
	// Read User Data & If User ID is not in DB, No result found 
	readTarget := User{ID: id}
	err := o.Read(&readTarget)
	if err == orm.ErrNoRows {
		log.Println("No ID")
		return &User{readTarget.ID, readTarget.Password, readTarget.Email}, errors.New(msg)	
	} 
	// If password is not matched with DB, No result found 
	if password != readTarget.Password {
		beego.Info("No result found.")
		return &User{readTarget.ID, readTarget.Password, readTarget.Email}, errors.New(msg)	
	}    else {
		beego.Info("Login Success")
		return &User{readTarget.ID, readTarget.Password, readTarget.Email}, nil
	}	

}