package main

import (
	// "encoding/json"
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/api"
	// "github.com/Shemistan/Lesson_6/internal/convert"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

func main() {

	db := storage.NewIStorage()
	service := service.NewIService(db)
	api := api.NewIApi(service)

	auth := models.SAuth{
		Login: "nabijonov",
		PasswordHash: "123",
	}
	user1 := models.SUser{Login: "nabijonov",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	user2 := models.SUser{Login: "qwerty",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	user3 := models.SUser{Login: "qwerty123",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	
	id1, _ := db.Add(&user1)
	id2, _ := db.Add(&user2)
	id3,_ :=service.Add(&user3)
	fmt.Println(id1)
	fmt.Println(id2)
	fmt.Println(id3)
	id4, _ := service.Auth(&auth)
	value, _ := service.GetUser(id4)
	_,_ =id4, value
	// fmt.Println(id4, "id polzova", value)
	// service.GetUsers()
	// service.GetStatistics()
	// service.GetUsers()

	str1 := `{"Login":"qwerty","PasswordHash":"123"}`   //json to service_struct
	apiIDUser,_ := api.Auth(str1)
	// fmt.Println(apiIDUser, "POluchilos")
	str2 := `{"Login":"testAddFromAPI","PasswordHash":"123","Name":"anvar","Surname":"nabijonov","Status":"active","Role":"active"}`
	boo1, _ := api.Add(str2)
	_ = boo1
	fmt.Println(apiIDUser, "POLZOVOTEL")
	_, str3 := api.GetUser(apiIDUser)
	fmt.Println("getting User From API", str3)
	api.GetUsers()
	
}
