package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/yuhaya/CloudServer/controller"
	"github.com/yuhaya/CloudServer/models"
	"github.com/yuhaya/CloudServer/tool"
	"log"
	"net/http"
	"reflect"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	CreatTables()
}

func NotFound(w http.ResponseWriter, r *http.Request) {

}

func main() {
	m := pat.New()
	m.Get("/", http.HandlerFunc(Root))
	//数据查询示例
	m.Get("/user", http.HandlerFunc(Dispense(&controller.User{}, "Hello")))

	err := http.ListenAndServe(":9006", m)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Dispense(ct controller.ControllerInterface, ac string) func(w http.ResponseWriter, r *http.Request) {

	t := reflect.TypeOf(ct)
	v := reflect.ValueOf(ct)
	ac_name := "Action" + ac
	method := v.MethodByName(ac_name)

	if !method.IsValid() {
		panic(ac_name + " method not exit")
	}

	ct_name := t.Elem().Name()

	return func(w http.ResponseWriter, r *http.Request) {
		ct.Before(w, r, ac_name, ct_name)
		args := make([]reflect.Value, 0)
		v.MethodByName(ac_name).Call(args)
	}
}

/**
 *建表
 */
func CreatTables() {
	db_gorm := tool.GetGormDB()
	db_gorm.CreateTable(&models.Admins{})
	db_gorm.CreateTable(&models.Attendances{})
	db_gorm.CreateTable(&models.CardReceiver{})
	db_gorm.CreateTable(&models.Cards{})
	db_gorm.CreateTable(&models.ClassTeacher{})
	db_gorm.CreateTable(&models.Classes{})
	db_gorm.CreateTable(&models.Devices{})
	db_gorm.CreateTable(&models.Families{})
	db_gorm.CreateTable(&models.FamilyMember{})
	db_gorm.CreateTable(&models.FamilyRelation{})
	db_gorm.CreateTable(&models.GradeClass{})
	db_gorm.CreateTable(&models.Grades{})
	db_gorm.CreateTable(&models.MemberCard{})
	db_gorm.CreateTable(&models.Schools{})
	db_gorm.CreateTable(&models.Students{})
	db_gorm.CreateTable(&models.System{})
	db_gorm.CreateTable(&models.Teachers{})
	db_gorm.CreateTable(&models.Users{})
}
