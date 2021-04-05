package models

import "time"

type User struct {
	Id int
	Name string
	Age int
	Addr string
	Pic string
	Phone string
}

type UserInfo struct {
	Id int
	Name string
	DBACreateTime time.Time
}

type DBXXXXUserInfo struct {   // dbxxxx_user_infos
	Id int
	Name string
}



