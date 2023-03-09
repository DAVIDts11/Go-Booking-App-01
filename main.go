//בס"ד
//https://www.youtube.com/watch?v=yyUHQIec83I&t=3788s

package main

import (
	"fmt"
	// "sync"
	"time"
	// "strconv"
	// "strings"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var  remainingtickets uint = conferenceTickets
// var bookings = make([]map[string]string,0)
var bookings = make([]UserData,0)

type UserData struct{ 
	firstName string
	lastName string
	email string
	numverOftickets uint
	
}
// var wg =sync.WaitGroup{}

func main(){

	//Another slice dedclaratiin syntaxes are :
	//var booking = []strting{}    
	//booking := []strting{}

	greetUser()

	for {
		firstName,lastName,email,userTickets :=getUserInput()

		isValidName,isValidEmail,isValidTicketNumber := validateUserInput(firstName,lastName,email,userTickets)
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets,firstName,lastName,email)

			// wg.Add(1)
			go sendTicket(userTickets,firstName,lastName,email)
		
			firstNames := getFirstNames();
			fmt.Printf("The first names of bookings are : %v .\n ",firstNames)
			if  remainingtickets == 0  {
				fmt.Printf("Our conference is booked out . Come back next year . \n") 
				break
			}

	
		}else{
			if !isValidName {
				fmt.Println("Your first or second name is to short ")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of  tickets you entered is invalid")
			}
		}

	}
	// wg.Wait()
}

func greetUser(){
	fmt.Println("Welcme to ",conferenceName," booking application" )
	fmt.Println("We have",conferenceName,"tickets and",remainingtickets,"are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _,booking := range bookings {
		// var names =  strings.Fields(booking)
		// firsNames = append(firsNames,names[0] )
		// firstNames = append(firstNames,booking["firstName"])
		firstNames = append(firstNames,booking.firstName)
	}
	return firstNames
}



func getUserInput()(string,string,string,uint){
	var firstName string 
	var lastName string 
	var email string 
	var userTickets uint 
	
	//ask user input :

	fmt.Println("Enter you first name please :")
	fmt.Scan(&firstName)


	fmt.Println("Enter you last name please :")
	fmt.Scan(&lastName)
	

	fmt.Println("Enter you email please :")
	fmt.Scan(&email)

	fmt.Println("How many tickets doyou want to book? :")
	fmt.Scan(&userTickets)

	return  firstName,lastName,email,userTickets 
}

func bookTicket(userTickets uint,firstName string,lastName string,email string){
	remainingtickets = remainingtickets  - userTickets

	// var userData =make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets),10)
	var userData  = UserData{
		firstName:firstName,
		lastName:lastName,
		email:email,
		numverOftickets:userTickets,
	}


	bookings = append(bookings,userData)
	fmt.Printf("List of bookins is %v  \n",bookings)
	fmt.Printf("")

	fmt.Printf("Thank you  %v %v for booking %v tickets .You will receive a confirmation email at %v .\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v .\n",remainingtickets,conferenceName)

}


func sendTicket(userTickets uint,firstName string,lastName string,email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v .",userTickets,firstName,lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v \nto email addres %v\n",ticket,email)
	fmt.Println("##############")
	// wg.Done()
}