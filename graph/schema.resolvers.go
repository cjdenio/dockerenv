package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cjdenio/dockerenv/graph/generated"
	"github.com/cjdenio/dockerenv/graph/model"
	"github.com/cjdenio/dockerenv/pkg/images"
)

func (r *queryResolver) Images(ctx context.Context) ([]*model.Image, error) {
	return images.LoadedImages.Images, nil
}

func (r *queryResolver) Image(ctx context.Context, name string) (*model.Image, error) {
	image, err := images.LoadedImages.FindOne(name)
	return image, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
