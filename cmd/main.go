package main

import (
	"encoding/json"
	"fmt"

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
	// stat := models.StatsUser{
	// 	DeleteUsersCount: 0,
	// 	UpdateCount: 0,
	// 	GetUserCount: 0,
	// 	GetUsersCount: 0,
	// 	GetAuthClick: 0,
	// }
	// fmt.Println(stat)
	
	test := models.SUser{Login: "nabijonov",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	test1 := models.SUser{Login: "qwerty",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	// mapa := map[int32]*models.SUser{}
	db := storage.NewIStorage()
	service := service.NewIService(db)
	boo, _ := db.Add(&test)
	boo1, _ := db.Add(&test1)
	fmt.Println(boo)
	fmt.Println(boo1)
	boo3,_ :=service.Add(&test1,&id)
	fmt.Println(boo3)
	idg, _ := service.Auth(&auth)
	value, _ := service.GetUser(idg)
	fmt.Println(idg, "id polzova", value)
	service.GetUsers()
	service.UpdateUser(idg,&test1)
	service.GetStatistics()
	service.GetUsers()
	jso, _ := json.Marshal(test)
	str :=string(jso)
	fmt.Println("json:",str)

}


// func sum(a, b int) int {
// 	return a + b
// }
