package main

import (
	// "encoding/json"
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

func main() {
	id := models.IdGenerate{
		IdGenerate: 0,
	}
	auth := models.SAuth{
		Login: "nabijonov",
		PasswordHash: "123",
	}
	test := models.SUser{Login: "nabijonov",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	test1 := models.SUser{Login: "qwerty",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	test2 := models.SUser{Login: "qwerty123",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	db := storage.NewIStorage()
	service := service.NewIService(db)
	api := api.NewIApi(service)
	boo, _ := db.Add(&test)
	boo1, _ := db.Add(&test2)
	boo3,_ :=service.Add(&test1,&id)
	fmt.Println(boo)
	fmt.Println(boo1)
	fmt.Println(boo3)
	idg, _ := service.Auth(&auth)
	value, _ := service.GetUser(idg)
	fmt.Println(idg, "id polzova", value)
	// service.GetUsers()
	// service.GetStatistics()
	// service.GetUsers()
	str1 := `{"Login":"qwerty","PasswordHash":"123"}`   //json to service_struct
	api1,_ := api.Auth(str1)
	fmt.Println(api1, "POluchilos")
	
}
