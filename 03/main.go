package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, ok := contactIndexByName[name]; ok {
		fmt.Println("Contact already exists: " + name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Println("Contact added: " + name)
}

func findContact(name string) *Contact {
	if index, ok := contactIndexByName[name]; ok {
		return &contactList[index]
	}
	return nil
}

func ListContact() {
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for _, contact := range contactList {
		fmt.Println("ID:", contact.ID, ", Name:", contact.Name, ", Email:", contact.Email, ", Phone:", contact.Phone)
	}
}

func main() {
	addContact("fsa", "fasdaa", "fsff")
	addContact("fsf", "fasddf", "favfdsddvddd")
	addContact("fsp", "favfdsddvd", "faaf")
	addContact("fsp", "favfdsd", "ffd")
	addContact("fsv", "vdasd", "ff")
	ListContact()

}
