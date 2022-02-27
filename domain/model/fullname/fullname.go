package fullname

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/Msksgm/itddd-go-02-valueobject/iterrors"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) (_ *FullName, err error) {
	defer iterrors.Wrap(&err, "fullname.NewFullName(%s, %s)", firstName, lastName)
	fullName := new(FullName)

	// set firstName
	err = fullName.setFirstName(firstName)
	if err != nil {
		return nil, err
	}

	// set lastName
	err = fullName.setLastName(lastName)
	if err != nil {
		return nil, err
	}

	return fullName, nil
}

func (fullName *FullName) setFirstName(firstName string) (err error) {
	if firstName == "" {
		return fmt.Errorf("firstName is required")
	}
	if !ValidateName(firstName) {
		return fmt.Errorf("firstName has an invalid character. letter is only")
	}
	fullName.firstName = firstName
	return nil
}

func (fullName *FullName) setLastName(lastName string) (err error) {
	if lastName == "" {
		return fmt.Errorf("lastName is required")
	}
	if !ValidateName(lastName) {
		return fmt.Errorf("lastName has an invalid character. letter is only")
	}
	fullName.lastName = lastName
	return nil
}

func ValidateName(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(value)
}

func (fullName *FullName) WithChangeFirstName(firstName string) (_ *FullName, err error) {
	changedFullName, err := NewFullName(firstName, fullName.lastName)
	if err != nil {
		return nil, err
	}
	return changedFullName, nil
}

func (fullName *FullName) WithChangeLastName(lastName string) (_ *FullName, err error) {
	changedFullName, err := NewFullName(fullName.firstName, lastName)
	if err != nil {
		return nil, err
	}
	return changedFullName, nil
}

func (fullName *FullName) Equals(otherFullName FullName) bool {
	return reflect.DeepEqual(fullName.firstName, otherFullName.firstName) && reflect.DeepEqual(fullName.lastName, otherFullName.lastName)
}
