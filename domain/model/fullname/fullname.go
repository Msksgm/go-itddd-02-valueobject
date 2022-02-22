package fullname

import (
	"fmt"
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
	if firstName == "" {
		return nil, fmt.Errorf("firstName is required")
	}
	if !ValidateName(firstName) {
		return nil, fmt.Errorf("firstName has an invalid character. letter is only")
	}
	fullName.firstName = firstName
	fullName.lastName = lastName

	return fullName, nil
}

func ValidateName(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(value)
}
