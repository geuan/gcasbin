package models

import "fmt"

type Role struct {
	RoleId int  `gorm:"column:role_id;primaryKey"`
	RoleName string  `gorm:"column:role_name"`
	RolePid int  `gorm:"column:role_pid;"`
	RoleComment string `gorm:"column:role_comment"`

}
func(this *Role) TableName() string{   // gorm会先校验我们是否有传表名，没有的话就调用TableName(）方法
	return "roles"
}
func(this *Role) String() string{
	return fmt.Sprintf("ID:%d 角色名:%s",this.RoleId,this.RoleName)
}
