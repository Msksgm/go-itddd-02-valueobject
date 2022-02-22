package fullname

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFullName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		firstName, lastName := "firstName", "lastName"
		got, err := NewFullName(firstName, lastName)
		if err != nil {
			log.Fatal(err)
		}

		want := &FullName{firstName: firstName, lastName: lastName}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(FullName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail firstName is empty", func(t *testing.T) {
		firstName, lastName := "", "lastName"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): firstName is required", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail firstName is invalid", func(t *testing.T) {
		firstName, lastName := "first$Name&", "lastName"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): firstName has an invalid character. letter is only", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail lastName is empty", func(t *testing.T) {
		firstName, lastName := "fistName", ""
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): lastName is required", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail lastName is invalid", func(t *testing.T) {
		firstName, lastName := "firstName", "last$Name&"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): lastName has an invalid character. letter is only", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		firstName, lastName := "firstName", "lastName"
		fullName, err := NewFullName(firstName, lastName)
		if err != nil {
			log.Fatal(err)
		}

		otherFullName := &FullName{firstName: firstName, lastName: lastName}
		if !fullName.Equals(*otherFullName) {
			t.Fatalf("fullName: %v must be equal to otherFullName: %v", fullName, otherFullName)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		firstName, lastName := "firstName", "lastName"
		fullName, err := NewFullName(firstName, lastName)
		if err != nil {
			log.Fatal(err)
		}

		otherFullName := &FullName{firstName: "otherFirstName", lastName: "otherFirstName"}
		if fullName.Equals(*otherFullName) {
			t.Fatalf("fullName: %v must not be equal to otherFullName: %v", fullName, otherFullName)
		}
	})
}
