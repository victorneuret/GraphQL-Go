// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/victorneuret/graphql-go/graph/model"

type Resolver struct{
	articles []*model.Article
	users []*model.User
}
