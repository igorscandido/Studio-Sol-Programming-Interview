package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/igorscandido/ssromanos/graph/generated"
	"github.com/igorscandido/ssromanos/graph/model"
	"github.com/igorscandido/ssromanos/providers"
)

func (r *mutationResolver) Search(ctx context.Context, text string) (*model.Response, error) {

	resposta := providers.SearchHighestRoman(text)

	return &resposta, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
