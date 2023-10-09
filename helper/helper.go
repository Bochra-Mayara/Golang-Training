package helper

import "strings"
//// capitalize first lettre to import the function
//// you can explort the variable, type, constant and functions

///3 level of scope
///Local :  declare within function, can be use within the function
////package: declare outside all functions,  use everywhere in same package
///Global : declare outside the all functions and uppercase first lettre, can be use everywhere all the packages
/// variable scope: where a defined variable can be accessed

func ValidationInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValName := len(firstName) >= 2 && len(lastName) >= 2 /// return boolean
	isValEmail := strings.Contains(email, "@")
	isValTickets := userTickets > 0 && userTickets <= remainingTickets
	////in Go be can return multiple values
	return isValName, isValEmail, isValTickets

}