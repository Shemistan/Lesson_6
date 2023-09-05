package convert

import (
	"encoding/json"
	"errors"
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

func ApiAuthConvertFromoService(user *models.SAuth) (string) {
	js,_ := json.Marshal(user)
	juser := string(js)
	// juser = 
	return juser
}

func ApiUserConvertToService(juser string)(models.SUser, error) {
	js := []byte(juser)
	user1 := models.SUser{}
	if json.Valid(js) {
		err := json.Unmarshal(js, &user1)
		if err != nil{
			fmt.Println(err)
		}
		return user1, nil
	}else{
		return user1, errors.New("Error Cant Convert json")
	}
}

func ApiUserConvertFromoService(user models.SUser) (string) {
	js,_ := json.Marshal(user)
	juser := string(js)
	// juser = 
	return juser
}

func ApiIdConvertToService(juser string)(models.IdGenerate, error) {
	js := []byte(juser)
	id := models.IdGenerate{}
	if json.Valid(js) {
		err := json.Unmarshal(js, &id)
		if err != nil{
			fmt.Println(err)
		}
		return id, nil
	}else{
		return id, errors.New("Error Cant Convert json")
	}
}

func ApiIdConvertFromoService(id models.IdGenerate) (string) {
	js,_ := json.Marshal(id)
	juser := string(js)
	// juser = 
	return juser
}
