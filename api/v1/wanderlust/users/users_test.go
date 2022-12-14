package users

import (
	"encoding/json"
	"testing"

	"github.com/mezmerizxd/go-app/pkg/data/users"
)

func TestMakeUser(t *testing.T) {
    new_user := users.User{
		Firstname: "test_1",
		Lastname: "test_2",
		Username: "test_3",
		Email: "test_4",
		Description: "",
	}

	want := usersUser{Username: "test_3", Description: ""}
    user := makeUser(&new_user)

	wantJSON, err1 := json.Marshal(want)
	if err1 != nil {
		t.Fatalf("Something went wrong when encoding JSON, %s", err1)
	}
	userJSON, err2 := json.Marshal(user)
	if err2 != nil {
		t.Fatalf("Something went wrong when encoding JSON, %s", err2)
	}

	for i := 0; i < len(wantJSON); i++ {
		if userJSON[i] != wantJSON[i] {
			t.Fatalf("Want: %s\nGot: %s", wantJSON, userJSON)
		}
	}
}