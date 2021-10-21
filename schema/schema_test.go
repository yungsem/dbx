package schema

import (
	"log"
	"reflect"
	"testing"
	"time"
)

type Model struct {
	CreateTime time.Time
	CreateUser string
	UpdateTime time.Time
	UpdateUser string
}
type User struct {
	Id       int
	Name     string
	Age      int
	RealName string
	Model
}

func TestParse(t *testing.T) {
	u := User{
		Id:   1,
		Name: "yangsen",
		Age:  30,
	}
	u.CreateUser = "1"
	u.CreateTime = time.Now()
	u.UpdateUser = "1"
	u.UpdateTime = time.Now()

	us := []User{}
	v := reflect.TypeOf(&us).Elem().Elem()


	sch, err := ParseColumns2(v)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(sch)
}
