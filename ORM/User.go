package ORM

import "github.com/victorneuret/graphql-go/graph/model"

func UserByName(name string) *model.User {
	mutex.Lock()

	var user model.User
	DB.First(&user, "name = ?", name)

	mutex.Unlock()
	return &user
}

func CreateUser(user *model.User) {
	mutex.Lock()

	DB.Create(user)

	mutex.Unlock()
}

func GetUsers() []*model.User {
	var usersP []*model.User

	mutex.Lock()

	var users []model.User
	DB.Find(&users)

	mutex.Unlock()

	for _, user := range users {
		usersP = append(usersP, &user)
	}
	return usersP
}
