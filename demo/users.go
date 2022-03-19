package Users

import "time"

//build by easyCli

type Users struct {
	Id int
	UserName string
	Password string
	Sex int
	Avatar string
	Age int
	CreateTime time.Time
	UpdateTime time.Time

}