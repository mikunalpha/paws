package paws_test

import (
	"testing"

	"github.com/mikunalpha/paws"
)

type User struct {
	Email string
	Name  *string
	Age   *int32
}

func TestPaws(t *testing.T) {
	_ = User{
		Email: "user001@email.com",
		Name:  paws.String("user001"),
		Age:   paws.Int32(3),
	}

	var requestEmail *string = nil
	var requestName *string = nil
	var requestAge *int32 = nil
	_ = User{
		Email: paws.MustString(requestEmail, "unknown@email.com"),
		Name:  paws.String(paws.MustString(requestName, "unknown")),
		Age:   requestAge,
	}
}
