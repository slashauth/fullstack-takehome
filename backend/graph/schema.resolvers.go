package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/getdebrief/fullstack-takehome/graph/generated"
	"github.com/getdebrief/fullstack-takehome/graph/model"
	"github.com/getdebrief/fullstack-takehome/notif"
	"github.com/sirupsen/logrus"
)

func (r *mutationResolver) AddSymbolToWatchList(ctx context.Context, id string) (*model.Symbol, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveSymbolFromWatchList(ctx context.Context, id string) (*model.Symbol, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WatchedSymbols(ctx context.Context) ([]*model.Symbol, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Symbol(ctx context.Context, id string) (*model.Symbol, error) {
	return &model.Symbol{
		ID:   "AAPL",
		Name: "AAPL",
		Sessions: []*model.TradingSession{
			{
				Time:  time.Now().Add(time.Minute * -1),
				Open:  10000,
				High:  10050,
				Low:   9990,
				Close: 10010,
			},
			{
				Time:  time.Now().Add(time.Minute * -2),
				Open:  9991,
				High:  10050,
				Low:   9990,
				Close: 10000,
			},
		},
	}, nil
}

func (r *subscriptionResolver) PriceUpdatesFromSymbol(ctx context.Context, id string) (<-chan *model.PriceUpdate, error) {
	events := make(chan *model.PriceUpdate, 1)

	err := notif.AddSubscriber(id, events)
	if err != nil {
		return nil, err
	}
	go func() {
		<-ctx.Done()
		close(events)

		err := notif.RemoveSubscriber(id, events)
		if err != nil {
			logrus.WithError(err).Errorf("Failed to remove subscriber")
		}
	}()

	return events, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
