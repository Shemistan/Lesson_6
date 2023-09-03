package convert

import (
	"encoding/json"
	// "errors"
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/models"
)

func ApiAuthConvertToService(juser string)(user models.SAuth) {
	js := []byte(juser)
	user1 := models.SAuth{}
	if json.Valid(js) {
		err := json.Unmarshal(js, &user1)
		// fmt.Println("FromUnmarshal-->", user1,"js",js)
		if err != nil{
			fmt.Println(err)
		}
		// fmt.Println(user1, "NRMALIZATSIYA")
		// fmt.Println(user1.Login,"-----",user1.PasswordHash)
		return user1
	}else{
		return user1
	}
}

func ApiAuthConvertFromoService(user *models.SAuth) (juser string) {
	js,_ := json.Marshal(user)
	juser = string(js)
	// juser = 
	return juser
}
