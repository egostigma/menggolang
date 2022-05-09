package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"main/graph/generated"
	"main/graph/model"
	"main/repository"
	"math/rand"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	todoRepository := repository.NewTodoRepository()
	todoRepository.Save(todo)
	return todo, nil
}

func (r *mutationResolver) UpdateStatusTodo(ctx context.Context, id string) (*model.Message, error) {
	todo, err := r.ServiceTodo.GetTodo(id)

	if err != nil {
		return nil, err
	}

	todo.Done = true

	err = r.ServiceTodo.UpdateTodo(todo)

	if err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Success",
	}, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   input.ID,
		Text: input.Text,
		Done: input.Done,
	}
	err := r.ServiceTodo.UpdateTodo(todo)

	if err != nil {
		return nil, err
	} else {
		return todo, nil
	}
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Message, error) {
	if err := r.ServiceTodo.DeleteTodo(id); err != nil {
		return nil, err
	} else {
		return &model.Message{
			Message: "Delete success",
		}, nil
	}
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todoRepository := repository.NewTodoRepository()
	return todoRepository.FindAll(), nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	todoRepository := repository.NewTodoRepository()
	return todoRepository.FindOne(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
