package main

import "fmt"

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %s could not be found\n", id)
}

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}
	fmt.Println("Returning user from database")
	u.LastSearchUsedCache = false
	return user, nil
}
