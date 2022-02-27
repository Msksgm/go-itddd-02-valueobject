package main

import (
	"fmt"
	"log"

	"github.com/Msksgm/itddd-go-02-valueobject/domain/model/fullname"
)

func main() {
	nameA, err := fullname.NewFullName("firstName", "lastName")
	if err != nil {
		log.Fatal(err)
	}
	nameB, err := fullname.NewFullName("firstName", "lastName")
	if err != nil {
		log.Fatal(err)
	}
	compareResult := nameA.Equals(*nameB)
	if compareResult {
		fmt.Printf("nameA: %v is equal to nameB: %v", nameA, nameB)
	} else {
		fmt.Printf("nameA: %v is not equal to nameB: %v", nameA, nameB)
	}
}
