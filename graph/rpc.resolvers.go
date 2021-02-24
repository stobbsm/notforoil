package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/stobbsm/notforoil/graph/generated"
	"github.com/stobbsm/notforoil/graph/model"
)

func (r *mutationResolver) Call(ctx context.Context, input model.CallInput) (*model.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Result(ctx context.Context, id string) (*model.Result, error) {
	var hello, world string = "Hello", "World!"
	var helloworld string = "Hello World!"
	var res = &model.Result{
		ID:     "1",
		Cmd:    "echo",
		Args:   []string{hello, world},
		Stdout: []string{helloworld},
		Stderr: []string{},
		Start:  time.Now(),
		End:    time.Now(),
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
