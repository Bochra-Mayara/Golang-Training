package main

///package fmt
import (
	"DemoAppGolang/helper"
	"fmt"
	"sync"
	"time"
)

    const conferenceTickets int = 50
    var conferenceName = "Go conference"
	var remainingTickets uint = 50
	// var bookings = []string{} //slices
	// var bookings = make([]map[string]string, 0) ///map
	var bookings = make([]UserData, 0) ///struct
    //// in struct can be mix many types of data
	type UserData struct {
		firstName string
		lastName string
		email string
		numberofTickets uint
	}
	var wg = sync.WaitGroup{}

func main() {
	
	greetUsers()
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	
	// var bookings = [50] string{} : array
	
	
		//bookings := [] string{} slice
		firstName, lastName, email, userTickets := getUserInput()
		isValName, isValEmail, isValTickets :=  helper.ValidationInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValName && isValEmail && isValTickets {

			bookTickets(userTickets , firstName , lastName,email)
			wg.Add(1)
			go SendTicket(userTickets , firstName , lastName,email)   /// key "go" to concurrancy start multipl threads

		///Call function FirstNames
        firstNames := getFirstNames()
		fmt.Printf("the first names of bookings are: %v\n", firstNames)


		////check sold out
		if(remainingTickets == 0){
			fmt.Println("our conference is booked out ")
			// break
		}
			
			
		}else{
			if !isValName {
				fmt.Println("firstName or llastName is entered is too short")
			} 
			if !isValEmail {
				fmt.Println("you main doesn't contain @")
			}
			if !isValTickets {
				fmt.Println("number tickets you entered is invalid")
			}
			
		}
		wg.Wait()

		
	

// 	city := "London"
// switch city {
//     case "New York":
// 	/// execute code
// 	case "Singapore", "Hong Kong":
// 		////
//     case "London" , "Berlin":
// 		////
// 	case "Mexico City":
// 		///
	
// 	default:
// 		fmt.Print("No vlid city selected")
// 	////
	
// }
}



func greetUsers()  {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend\n ")
	
}



func getFirstNames() []string  {
	firstNames := []string{}

		///_: ignore the var you don't to use
		for _, booking := range bookings {
			firstNames = append(firstNames, booking.firstName)
		}
	 return firstNames

	
}





func getUserInput()(string , string, string, uint)  {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name :")
	fmt.Scan(&lastName)
	fmt.Println("Enter your mail :")
	fmt.Scan(&email)

	fmt.Println("Enter your numberTickets :")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTickets(userTickets uint, firstName string, lastName string,email string )  {
	    remainingTickets = remainingTickets - userTickets
		/// in the map : key in the same type and value also  , support 1 data type
        //  var userData = make(map[string] string)
		//  /////userData["key"] = value
		//  userData["firstName"] = firstName
		//  userData["lastName"] = lastName
		//  userData["email"] = email
		//  userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)   /// transfer uint number to string
		// bookings = append(bookings, userData)


		////struct App

         var userData = UserData{
			firstName:  firstName ,
			lastName: lastName ,
			email:  email,
			numberOfTickets : userTickets ,
			
		 }
		
		bookings = append(bookings, userData)

		fmt.Printf("List of bookings %v \n", bookings)
		
		// fmt.Println(userName)
		fmt.Printf("Thank you %v %v for booking  %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
	



}


func SendTicket(userTickets uint, firstName string, lastName string, email string)  {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket %v to email address %v \n" , ticket, email)
	fmt.Println("#########")
	wg.Done()
	
}

