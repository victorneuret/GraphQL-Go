package ORM

import "github.com/victorneuret/graphql-go/graph/model"

func CreateTodo(todo *model.Todo) {
	mutex.Lock()

	DB.Create(todo)
	DB.Save(todo)

	mutex.Unlock()
}

func GetTodos() []*model.Todo {
	var todosP []*model.Todo

	mutex.Lock()

	var todos []model.Todo
	DB.Find(&todos)

	mutex.Unlock()

	for _, todo := range todos {
		todosP = append(todosP, &todo)
	}
	return todosP
}

func UpdateTodo(input model.UpdateTodo) *model.Todo {
	mutex.Lock()

	var todo model.Todo
	DB.Model(&todo).Where("user_name = ?", input.UserName, "id = ?", input.ID).Update("done", input.Done)
	DB.First(&todo, input.ID).Where("user_name = ?", input.UserName)

	mutex.Unlock()
	return &todo
}
