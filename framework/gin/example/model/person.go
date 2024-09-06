package model

import (
	"errors"
	"fmt"
)

// xorm默认的名称映射是驼峰转下划线，所以当结构体字段名转下划线后和数据库字段名不一致的时候，
// 需要使用单引号指定数据库中的字段名称，比如这里的UserName
type Person struct {
	UserId   int    `xorm:"not null pk autoincr int64 'user_id'" json:"userId,omitempty"`
	UserName string `xorm:"varchar(260) 'username' comment('标题')" json:"userName,omitempty"`
	Sex      string `xorm:"varchar(260) comment('标题')" json:"sex,omitempty"`
	Email    string `xorm:"varchar(260) comment('标题')" json:"email,omitempty"`
}

func AddPerson(person *Person) error {
	ok, err := db.Exist(&Person{UserName: person.UserName})
	if err != nil {
		return err
	}
	if ok {
		return errors.New(MErrEsisted)
	}
	_, err = db.Insert(person)
	if err != nil {
		return err
	}
	return nil
}

func DeletePerson(id int) error {
	ok, err := db.Exist(&Person{UserId: id})
	if err != nil {
		return err
	}
	if ok {
		return errors.New(MErrNotFind)
	}
	count, err := db.ID(id).Delete(&Person{})
	if err != nil || count <= 0 {
		return errors.New(MErrDelete)
	}
	return nil
}

func UpdatePerson(data *Person) error {
	id, err := db.ID(data.UserId).Update(data)
	if err != nil || id <= 0 {
		fmt.Println(err.Error())
		return errors.New(MErrUpdate)
	}
	return nil
}

func GetPersonById(id int) (info Person, errs error) {
	ok, err := db.ID(id).Get(&info)
	if !ok || err != nil {
		errs = errors.New(MErrNotFind)
		return
	}
	return
}

func GetPersonAll() (info []Person, err error) {
	err = db.Where("username!=''").Find(&info)
	return
}
