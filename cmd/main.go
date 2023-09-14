package main

import (
	"fmt"
	"strconv"

	apiInternal "github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
	serviceInternal "github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

func main() {
	db := storage.New()
	service := serviceInternal.New(db)
	api := apiInternal.New(service)
	Server(api)
}

func Server(api apiInternal.IApi) {
	var answer string
	fmt.Println("Hello, dear user!")
	for answer != "exit" {
		fmt.Printf("\nWhat would you like to do?\n")
		fmt.Println("Choose one of the following:")
		fmt.Println("1. Authenticate")
		fmt.Println("2. Update your profile")
		fmt.Println("3. Get information about user using id")
		fmt.Println("4. Get information about all users")
		fmt.Println("5. Delete an account using id")
		fmt.Println("6. Get statistics about clicks")
		fmt.Println("Write a digit from 1 to 6")
		fmt.Printf("If you want to exit, type \"exit\" to exit\n")

		fmt.Scan(&answer)
		switch answer {

		case "1":
			var login, password, nextAction string
			for nextAction != "home" {
				fmt.Printf("\nPlease, note that if there is no user with this login\n")
				fmt.Println("A new user will be created")
				fmt.Println("Please, enter your login and password:")

				fmt.Scan(&login, &password)
				id, err := api.Auth(login, password)
				if err != nil {
					fmt.Println("Login is correct, but password is wrong")
					fmt.Println("Please enter your login and password again")
					continue
				}

				fmt.Printf("Your id is %d!", id)
				fmt.Println("What would you like to do next?")
				fmt.Printf("Type \"con\" to continue\n")
				fmt.Printf("Type \"home\" to go back to main page\n")

				fmt.Scan(&nextAction)
				if nextAction == "con" {
					continue
				}
			}
			fmt.Println("Redirecting to main page...")

		case "2":
			var accountInfo models.Account
			var id int64
			var nextAction string
			for nextAction != "home" {
				fmt.Println("Please enter your account id:")

				fmt.Scan(&id)

				fmt.Println("Please enter your name:")

				fmt.Scan(&accountInfo.UserInfo.Name)

				fmt.Println("Please enter your surname:")

				fmt.Scan(&accountInfo.UserInfo.Surname)

				fmt.Println("Is your profile active?(yes/no)")

				var answer string

				fmt.Scan(&answer)
				if answer == "yes" {
					accountInfo.UserInfo.Active = true
				} else {
					accountInfo.UserInfo.Active = false
				}

				fmt.Println("Please enter your role:")

				fmt.Scan(&accountInfo.UserInfo.Role)
				api.UpdateUser(id, accountInfo)

				fmt.Println("To see your profile you need")
				fmt.Println("to choose 3rd option on home menu")
				fmt.Printf("\nWhat would you like to do next?\n")
				fmt.Printf("Type \"con\" to continue\n")
				fmt.Printf("Type \"home\" to go back to main page\n")

				fmt.Scan(&nextAction)

				if nextAction == "con" {
					continue
				}
				break
			}
			fmt.Println("Redirecting to main page...")

		case "3":
			var id string
			var nextAction string
			for nextAction != "home" {
				fmt.Println("Please enter account id(type 'exit' to exit):")

				fmt.Scan(&id)
				if id == "exit" {
					break
				}

				id, _ := strconv.ParseInt(id, 10, 64)
				model, err := api.GetUser(id)
				if err != nil {
					fmt.Println("Wrong ID! Please enter user id again")
					continue
				}

				fmt.Printf("Name: %s\n", model.Name)
				fmt.Printf("Surname: %s\n", model.Surname)
				if model.Active {
					fmt.Println("Is active: yes")
				} else {
					fmt.Println("Is active: no")
				}
				fmt.Printf("Role: %s\n", model.Role)
				fmt.Printf("Registration date: %s\n", model.RegistrationDate)
				fmt.Printf("Update date: %s\n", model.UpdateDate)
				fmt.Println("What would you like to do next?")
				fmt.Printf("Type \"con\" to continue\n")
				fmt.Printf("Type \"home\" to go back to main page\n")

				fmt.Scan(&nextAction)
				if nextAction == "con" {
					continue
				}
			}
			fmt.Println("Redirecting to main page...")

		case "4":
			listOfUsers, err := api.GetUsers()

			if err != nil {
				fmt.Println()
				fmt.Println("There are no records!")
				fmt.Println("Redirecting to main page...")
				continue
			}

			for _, val := range listOfUsers {
				fmt.Printf("\nName: %s\n", val.Name)
				fmt.Printf("Surname: %s\n", val.Surname)
				if val.Active {
					fmt.Println("Is active: yes")
				} else {
					fmt.Println("Is active: no")
				}
				fmt.Printf("Role: %s\n", val.Role)
				fmt.Printf("Registration date: %s\n", val.RegistrationDate)
				fmt.Printf("Update date: %s\n", val.UpdateDate)
			}
			fmt.Println("\nRedirecting to main page...")

		case "5":
			var id int64
			var nextAction string
			for nextAction != "home" {
				fmt.Println("Please enter account id:")

				fmt.Scan(&id)

				err := api.DeleteUser(id)
				if err != nil {
					fmt.Println("Wrong ID! Please enter user id again")
					continue
				}
				fmt.Printf("User with id %d was deleted\n", id)

				stat := api.GetStatistics()

				fmt.Printf("Number of users deleted: %d\n", stat["DeletedUsers"])
				fmt.Println("What would you like to do next?")
				fmt.Printf("Type \"con\" to continue\n")
				fmt.Printf("Type \"home\" to go back to main page\n")

				fmt.Scan(&nextAction)

				if nextAction == "con" {
					continue
				}
				break
			}
			fmt.Println("Redirecting to main page...")

		case "6":
			var nextAction string
			for nextAction != "home" {
				stats := api.GetStatistics()

				fmt.Printf("\"Authenticate\" number of clicks: %d\n", stats["Auth"])
				fmt.Printf("\"Updating profile\" number of clicks: %d\n", stats["UpdateUser"])
				fmt.Printf("\"Getting user info\" number of clicks: %d\n", stats["GetUser"])
				fmt.Printf("\"Getting info of all users\" number of clicks: %d\n", stats["GetUsers"])
				fmt.Printf("\"Delete user\" number of clicks: %d\n", stats["DeleteUser"])
				fmt.Printf("\"Getting stats\" number of clicks: %d\n", stats["GetStatistics"])
				fmt.Printf("Number of deleted users: %d\n", stats["DeletedUsers"])
				fmt.Println("What would you like to do next?")
				fmt.Printf("Type \"con\" to continue\n")
				fmt.Printf("Type \"home\" to go back to main page\n")

				fmt.Scan(&nextAction)

				if nextAction == "con" {
					continue
				}
				break
			}
			fmt.Println("Redirecting to main page...")
		default:
			if answer != "exit" {
				fmt.Println()
				fmt.Println("Please enter a valid action")
			}
		}

	}

}
