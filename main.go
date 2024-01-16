package main

import (
	"fmt"
	crud "main/CRUD"
	"main/models"

	"github.com/k0kubun/pp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//made like a project with command options and it is working till the user will not exit from the program
//used switch cases
//added additional functions

func main() {
	connection := "host=localhost user=postgres password=mubina2007 dbname=homework sslmode=disable TimeZone=Asia/Tashkent"
	mydb, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database.")
	}

	if err = mydb.AutoMigrate(&models.People{}, &models.Activity{}); err != nil {
		panic("Cannot create a table.")
	}

	exit := true
	for exit {
		fmt.Printf("\t\t\tMENU\n")
		fmt.Printf("\t\t1.Insert into People.\n")
		fmt.Printf("\t\t2.Insert into Activities.\n")
		fmt.Printf("\t\t3.Update table People.\n")
		fmt.Printf("\t\t4.Update table Activities.\n")
		fmt.Printf("\t\t5.Get People.\n")
		fmt.Printf("\t\t6.Get Activities.\n")
		fmt.Printf("\t\t7.Delete People.\n")
		fmt.Printf("\t\t8.Delete Activities.\n")
		fmt.Printf("\t\t9.Exit.\n")
		var option int
		fmt.Print("Enter a command: ")
		fmt.Scan(&option)
		switch option {
		case 1:
			var people []*models.People
			var n int
			fmt.Printf("How many people you want to insert?: ")
			fmt.Scan(&n)
			for i := 0; i < n; i++ {
				var full_name, phone_number string
				fmt.Printf("Input %d person's full name: ", i+1)
				fmt.Scan(&full_name)
				fmt.Printf("Input %d person's phone_number: ", i+1)
				fmt.Scan(&phone_number)
				person := models.People{
					Full_name:    full_name,
					Phone_number: phone_number,
				}
				var c int
				fmt.Printf("How many activities does %s attend?: ", full_name)
				fmt.Scan(&c)
				var activities []*models.Activity
				for k := 0; k < c; k++ {
					var activityName, supervisor string
					fmt.Printf("Input %d activity's name that %s attend: ", k+1, full_name)
					fmt.Scan(&activityName)
					fmt.Printf("Input %s supervisor name that %s attend: ", activityName, full_name)
					fmt.Scan(&supervisor)
					activity := models.Activity{
						ActivityName: activityName,
						Supervisor:   supervisor,
					}
					activities = append(activities, &activity)
					person.Activity = activities
				}
				people = append(people, &person)
			}
			err := crud.CreatePeople(mydb, people)
			if err != nil {
				panic(err)
			}
			pp.Println("All datas were successfully inserted!")
		case 2:
			var activities []*models.Activity
			var n int
			fmt.Printf("How many activities you want to insert?: ")
			fmt.Scan(&n)
			for i := 0; i < n; i++ {
				var activity_name, supervisor string
				fmt.Printf("Input %d activity name: ", i+1)
				fmt.Scan(&activity_name)
				fmt.Printf("Input %s's supervisor: ", activity_name)
				fmt.Scan(&supervisor)
				activity := models.Activity{
					ActivityName: activity_name,
					Supervisor:   supervisor,
				}
				var c int
				fmt.Printf("How many attendants does %s have?: ", activity_name)
				fmt.Scan(&c)
				var people []*models.People
				for k := 0; k < c; k++ {
					var full_name, phone_number string
					fmt.Printf("Input %d attendant's full name: ", k+1)
					fmt.Scan(&full_name)
					fmt.Printf("Input %d attendant's phone_number: ", k+1)
					fmt.Scan(&phone_number)
					person := models.People{
						Full_name:    full_name,
						Phone_number: phone_number,
					}
					people = append(people, &person)
					activity.People = people
				}
				activities = append(activities, &activity)
			}
			err := crud.CreateActivity(mydb, activities)
			if err != nil {
				panic(err)
			}
			pp.Println("All datas were successfully inserted!")
		case 3:
			x := true
			for x {
				fmt.Println()
				fmt.Printf("1.Update Full name       2.Update Phone Number       3.Exit\n")
				fmt.Printf("Choose an option: ")
				var a int
				fmt.Scan(&a)
				switch a {
				case 1:
					fmt.Printf("Enter ID of a person you want to update: ")
					var id int
					fmt.Scan(&id)
					fmt.Printf("Enter a new full name: ")
					var new_fullName string
					fmt.Scan(&new_fullName)
					err := crud.UpdatePersonFullNameByid(mydb, id, new_fullName)
					if err != nil {
						panic(err)
					}
					fmt.Println("Successfully updated!")
				case 2:
					fmt.Printf("Enter ID of a person you want to update: ")
					var id int
					fmt.Scan(&id)
					fmt.Printf("Enter a new phone number: ")
					var phone string
					fmt.Scan(&phone)
					err := crud.UpdatePersonPhoneNumberById(mydb, id, phone)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println("Successfully updated!")
					}
				case 3:
					x = false
				default:
					pp.Println("Invalid command. Try again.")
				}
			}
		case 4:
			x := true
			for x {
				fmt.Println()
				fmt.Printf("1.Update Activity name        2.Update Supervisor name       3.Exit\n")
				fmt.Printf("Choose an option: ")
				var a int
				fmt.Scan(&a)
				switch a {
				case 1:
					fmt.Printf("Enter ID of an activity you want to update: ")
					var id int
					fmt.Scan(&id)
					fmt.Printf("Enter a new Activity name: ")
					var new_Name string
					fmt.Scan(&new_Name)
					err := crud.UpdateActivityNameById(mydb, id, new_Name)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println("Successfully updated!")
					}
				case 2:
					fmt.Printf("Enter ID of an activity you want to update: ")
					var id int
					fmt.Scan(&id)
					fmt.Printf("Enter a new supervisor name: ")
					var name string
					fmt.Scan(&name)
					err := crud.UpdateActivitySupervisorById(mydb, id, name)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println("Successfully updated!")
					}
				case 3:
					x = false
				default:
					pp.Println("Invalid command. Try again.")
				}
			}
		case 5:
			x := true
			for x {
				fmt.Printf("1.Get person by ID    2.Get person by Full name    3.Get all people   4.Exit\n")
				var c int
				fmt.Println()
				fmt.Printf("Choose a command: ")
				fmt.Scan(&c)
				switch c {
				case 1:
					var id int
					fmt.Printf("Input ID of a person you want to get: ")
					fmt.Scan(&id)
					person, err := crud.GetPeopleById(mydb, id)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println(person)
					}
					fmt.Println()
				case 2:
					var fullName string
					fmt.Printf("Input full name of a person you want to get: ")
					fmt.Scan(&fullName)
					person, err := crud.GetPeopleByFullName(mydb, fullName)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println(person)
					}
					fmt.Println()
				case 3:
					allPeople, err := crud.GetAllPeople(mydb)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println(allPeople)
					}
					fmt.Println()
				case 4:
					x = false
				default:
					pp.Println("Invalid command. Try again!")
				}
			}
		case 6:
			x := true
			for x {
				fmt.Printf("1.Get activity by ID    2.Get activity by name   3.Get all activities  4.Exit\n")
				var c int
				fmt.Println()
				fmt.Printf("Choose a command: ")
				fmt.Scan(&c)
				switch c {
				case 1:
					var id int
					fmt.Printf("Input ID of an activity you want to get: ")
					fmt.Scan(&id)
					activity, err := crud.GetActivityById(mydb, id)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println(activity)
					}
					fmt.Println()
				case 2:
					var name string
					fmt.Printf("Input a name of an activity you want to get: ")
					fmt.Scan(&name)
					activity, err := crud.GetActivityByName(mydb, name)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println(activity)
					}
					fmt.Println()
				case 3:
					allActivities, err := crud.GetAllActivities(mydb)
					if err != nil {
						fmt.Println(err)
					} else {
						pp.Println(allActivities)
					}
					fmt.Println()
				case 4:
					x = false
				default:
					pp.Println("Invalid command. Try again!")
				}
			}
		case 7:
			x := true
			for x {
				fmt.Printf("1.Delete person by ID  2.Exit\n")
				var c int
				fmt.Println()
				fmt.Printf("Choose a command: ")
				fmt.Scan(&c)
				switch c {
				case 1:
					var id int
					fmt.Printf("Input ID of a person you want to delete: ")
					fmt.Scan(&id)
					person, err := crud.DeletePeopleById(mydb, id)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Successfully deleted a person with data: ")
						pp.Println(person)
					}
					fmt.Println()
				case 2:
					x = false
				default:
					pp.Println("Invalid command. Try again!")
				}
			}
		case 8:
			x := true
			for x {
				fmt.Printf("1.Delete activity by ID   2.Exit\n")
				var c int
				fmt.Println()
				fmt.Printf("Choose a command: ")
				fmt.Scan(&c)
				switch c {
				case 1:
					var id int
					fmt.Printf("Input ID of an activity you want to delete: ")
					fmt.Scan(&id)
					activity, err := crud.DeleteActivityById(mydb, id)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Successfuly deleted activity with data: ")
						pp.Println(activity)
					}
					fmt.Println()
				case 2:
					x = false
				default:
					fmt.Println("Invalid command. Try again!")
				}
			}
		case 9:
			pp.Println("Bye.")
			exit = false
		}
	}

}
