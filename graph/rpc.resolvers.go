package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sync"

	uuid "github.com/satori/go.uuid"
	"github.com/stobbsm/notforoil"
	"github.com/stobbsm/notforoil/graph/generated"
	"github.com/stobbsm/notforoil/graph/model"
)

var results = make(map[uuid.UUID]*model.Result)

func (r *mutationResolver) Call(ctx context.Context, input model.CallInput) (*model.Result, error) {
	wg := new(sync.WaitGroup)
	c := notforoil.NewCommand(wg, input.Cmd, input.Args...)
	results[c.ID] = &model.Result{
		ID:     c.ID.String(),
		Cmd:    c.Cmd,
		Args:   c.Args,
		Stdout: []string{},
		Stderr: []string{},
		Start:  c.Start,
		End:    c.End,
	}
	go func(wg *sync.WaitGroup, c *notforoil.Command) {
	loop:
		for {
			select {
			case out, ok := <-c.Out:
				if out != "" {
					results[c.ID].Stdout = append(results[c.ID].Stdout, out)
				}
				if !ok {
					break loop
				}
			case err := <-c.Err:
				if err != "" {
					results[c.ID].Stderr = append(results[c.ID].Stderr, err)
					break loop
				}
			}
		}
		wg.Wait()
		results[c.ID].Start = c.Start
		results[c.ID].End = c.End
	}(wg, c)
	c.Do(ctx)

	return results[c.ID], nil
}

func (r *queryResolver) Result(ctx context.Context, id string) (*model.Result, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return results[uid], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
