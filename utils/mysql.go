package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Link struct {
	gorm.Model
	Url string 		`gorm:"unique_index"`
	Keyword string	`gorm:"unique_index"`
}

type MysqlConn struct {
	DB *gorm.DB
}
/**
创建数据库连接
*/
func NewDB(user, password, ip string, port int, database string) (*MysqlConn, string) {
	mysqlConn := new(MysqlConn)
	var err error
	mysqlConn.DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, ip, port, database))
	if err != nil {
		panic(err)
	}
	mysqlConn.DB.AutoMigrate(&Link{})
	return mysqlConn, mysqlConn.getRecordsNum()
}
/**
获取最大短键
 */
func (mysqlConn *MysqlConn) getRecordsNum() string {
	type MaxStruct struct {
		Max string `json:"max"`
	}
	var result MaxStruct
	err := mysqlConn.DB.Raw("SELECT MAX( links.keyword ) AS max FROM links").Scan(&result).Error
	if err != nil {
		return ""
	}
	return result.Max
}
/**
从数据库中查询长网址对应的短键
*/
func (mysqlConn *MysqlConn) GetKeyword(url string) string {
	var link Link
	if mysqlConn.DB.First(&link, "url = ?", url).Error != nil {
		return ""
	}
	return link.Keyword
}
/**
更新数据库
*/
func (mysqlConn *MysqlConn) Update(keyword, url string) bool {
	link := Link{Keyword: keyword, Url: url}
	if nil != mysqlConn.DB.Create(&link).Error {
		return false
	}
 	return true
}
/**
从数据库中查询短键对应的长网址
*/
func (mysqlConn *MysqlConn) GetUrlFromDB(keyword string) string {
	var link Link
	if mysqlConn.DB.First(&link, "keyword = ?", keyword).Error != nil {
		return ""
	}
	return link.Url
}

