package fullname

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) (_ *FullName, err error) {
	fullName := new(FullName)

	fullName.firstName = firstName
	fullName.lastName = lastName

	return fullName, nil
}
