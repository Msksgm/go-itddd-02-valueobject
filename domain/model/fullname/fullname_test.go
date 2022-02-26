package fullname

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	firstName = "firstName"
	lastName  = "lastName"
)

func TestNewFullName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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
		firstName := ""
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): firstName is required", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail firstName is invalid", func(t *testing.T) {
		firstName := "first$Name&"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): firstName has an invalid character. letter is only", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail lastName is empty", func(t *testing.T) {
		lastName := ""
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): lastName is required", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail lastName is invalid", func(t *testing.T) {
		lastName := "last$Name&"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): lastName has an invalid character. letter is only", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestWithChangeFirstName(t *testing.T) {
	fullName, err := NewFullName(firstName, lastName)
	if err != nil {
		log.Fatal(err)
	}
	newFirstName := "newFirstName"
	got, err := fullName.WithChangeFirstName(newFirstName)
	if err != nil {
		log.Fatal(err)
	}
	want := &FullName{firstName: newFirstName, lastName: lastName}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(FullName{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestWithChangeLastName(t *testing.T) {
	fullName, err := NewFullName(firstName, lastName)
	if err != nil {
		log.Fatal(err)
	}
	newLastName := "newLastName"
	got, err := fullName.WithChangeLastName(newLastName)
	if err != nil {
		log.Fatal(err)
	}
	want := &FullName{firstName: firstName, lastName: newLastName}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(FullName{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
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
