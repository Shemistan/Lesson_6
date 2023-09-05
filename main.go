package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_6/service"
	"github.com/Shemistan/Lesson_6/storage"
)

func main() {
	st := storage.NewStorage()
	ser := service.NewUserService(st)
	ser.Add("test", "test", "test", "123")
	user, _ := ser.Get(0)
	fmt.Println(user)
}
