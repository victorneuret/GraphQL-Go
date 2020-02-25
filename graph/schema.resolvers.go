// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"errors"
	"fmt"

	"github.com/victorneuret/graphql-go/graph/generated"
	"github.com/victorneuret/graphql-go/graph/model"
)

func (r *articleResolver) User(ctx context.Context, obj *model.Article) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	found := false
	for _, v := range r.users {
		if v.ID == input.UserID {
			found = true
			break
		}
	}
	if found == false {
		return nil, errors.New("User " + input.UserID + " not found")
	}

	article := &model.Article{
		Title:  input.Title,
		Text:   input.Text,
		ID:     fmt.Sprintf("%d", len(r.articles)),
		UserID: input.UserID,
	}
	r.articles = append(r.articles, article)
	return article, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID: fmt.Sprintf("%d", len(r.users)),
		Name: input.Name,
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	return r.articles, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *Resolver) Article() generated.ArticleResolver   { return &articleResolver{r} }
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type articleResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
