// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"

	"github.com/victorneuret/graphql-go/ORM"
	"github.com/victorneuret/graphql-go/graph/generated"
	"github.com/victorneuret/graphql-go/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user := ORM.UserByName(input.UserName)

	todo := &model.Todo{
		Text:     input.Text,
		Done:     input.Done,
		UserName: user.Name,
	}

	ORM.CreateTodo(todo)

	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	todo := ORM.UpdateTodo(input)

	return todo, nil
	//iterator := -1
	//for i, v := range r.todos {
	//	if v.ID == input.ID {
	//		iterator = i
	//		break
	//	}
	//}
	//if iterator == -1 {
	//	return nil, errors.New("Todo with ID '" + input.ID + "' does not exist")
	//}
	//
	//found := false
	//for _, v := range r.users {
	//	if v.ID == input.UserID {
	//		found = true
	//		break
	//	}
	//}
	//if found == false {
	//	return nil, errors.New("User " + input.UserID + " not found")
	//}
	//
	//todo := r.todos[iterator]
	//
	//if todo.UserID != input.UserID {
	//	return nil, errors.New("User " + input.UserID + " Can't modify this todo")
	//}
	//
	//todo = &model.Todo{
	//	ID:     todo.ID,
	//	Text:   input.Text,
	//	Done:   input.Done,
	//	UserID: todo.UserID,
	//}
	//r.todos[iterator] = todo
	//return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name: input.Name,
	}
	ORM.CreateUser(user)
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return ORM.GetTodos(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return ORM.GetUsers(), nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
