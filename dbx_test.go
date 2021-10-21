package dbx

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	db, err := Connect("mysql", "root:root@tcp(localhost:3306)/go_rbac?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%v\n", db)
}

type Model struct {
	CreateTime time.Time `db:"CREATE_TIME"`
	CreateUser string    `db:"CREATE_USER"`
	UpdateTime time.Time `db:"UPDATE_TIME"`
	UpdateUser string    `db:"UPDATE_USER"`
	Deleted    int       `db:"DELETED"`
}
type User struct {
	Id       int    `db:"ID"`
	Username string `db:"USERNAME"`
	Password string `db:"PASSWORD"`
	RealName string `db:"REAL_NAME"`
	Model
}

func TestDB_Insert(t *testing.T) {
	db, err := Connect("mysql", "root:root@tcp(localhost:3306)/go_rbac?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	u := User{
		Username: "yangsen",
		Password: "123456",
		RealName: "杨森",
	}
	u.CreateUser = "admin"
	u.CreateTime = time.Now()
	u.UpdateUser = "admin"
	u.UpdateTime = time.Now()

	err = db.Insert(&u)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestDB_One(t *testing.T) {
	db, err := Connect("mysql", "root:root@tcp(localhost:3306)/go_rbac?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	cond := db.Cond().Eq("id", 3)

	var u User
	err = db.One(cond, &u)
	if err != nil {
		log.Fatalln(err)
	}

	u.Password = "adddd"
	err = db.Update(&u)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(u)
}

func TestDB_List(t *testing.T) {
	db, err := Connect("mysql", "root:root@tcp(localhost:3306)/go_rbac?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	cond := db.Cond().Gt("create_time", "2021-10-20 14:23:01")

	var u []User
	err = db.List(cond, &u)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(u)
}

func TestDB_Li(t *testing.T) {
	db, err := Connect("mysql", "root:root@tcp(localhost:3306)/go_rbac?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	ids := []int{5, 6}
	cond := db.Cond().In("id", ids)

	var u []User
	err = db.List(cond, &u)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(u)
}