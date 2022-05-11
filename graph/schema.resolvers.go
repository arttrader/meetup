package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/arttrader/meetup/graph/generated"
	"github.com/arttrader/meetup/graph/model"
)

func (r *meetupResolver) ID(ctx context.Context, obj *model.Meetup) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *meetupResolver) Name(ctx context.Context, obj *model.Meetup) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *meetupResolver) Description(ctx context.Context, obj *model.Meetup) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
