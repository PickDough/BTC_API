package dal

import (
	"User_Service/domain"
	"os"
	"testing"
)

func TestFileRepository_AddWillCreateFile(t *testing.T) {
	file := "users_tmp.data"
	repo := FileRepository{FileLocation: file}

	err := repo.Add(domain.User{Email: "addTest", Password: "addTest"})
	if err != nil {
		t.Errorf(err.Error())
	}

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("expected file at %s", file)
		}
	} else {
		_ = os.Remove(file)
	}

}

func TestFileRepository_GetWillReturnUserFromFile(t *testing.T) {
	file := "users_tmp.data"
	addRepo := FileRepository{FileLocation: file}
	addUser := domain.User{Email: "getTest", Password: "getTest"}

	err := addRepo.Add(addUser)
	if err != nil {
		t.Errorf(err.Error())
	}
	//Recreate repository because previous would return from cash and not file
	getRepo := FileRepository{FileLocation: file}

	getUser, err := getRepo.Get(addUser.Email)
	if err != nil {
		t.Errorf(err.Error())
	}

	if getUser == nil || getUser.Email != addUser.Email {
		t.Errorf("added user wasn't found")
	}

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("expected file at %s", file)
		}
	} else {
		_ = os.Remove(file)
	}
}
