package models

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/jaykof/chitchat/config"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	//config := LoadConfig()	// 加载全局配置实例
	driver := ViperConfig.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		ViperConfig.Db.User, ViperConfig.Db.Password, ViperConfig.Db.Address, ViperConfig.Db.Database)
	// 数据库连接字符串中增加 loc=Local ，避免出现时间传递到mysql里出现时区偏差8小时问题
	//Db, err = sql.Open("mysql", "root@/chitchat?charset=utf8&parseTime=true&loc=Local")
	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// create a random UUID with from RFC 4122
// adapted from http://gitbub.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1A
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
