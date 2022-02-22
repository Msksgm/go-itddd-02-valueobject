package fullname

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFullName(t *testing.T) {
	firstName, lastName := "firstName", "lastName"
	got, err := NewFullName(firstName, lastName)
	if err != nil {
		log.Fatal(err)
	}

	want := &FullName{firstName: firstName, lastName: lastName}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(FullName{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
