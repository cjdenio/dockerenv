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

func (r *queryResolver) Search(ctx context.Context, query string, maxItems *int) ([]*model.Image, error) {
	matched_images := images.LoadedImages.Search(query)

	if len(matched_images) > *maxItems {
		matched_images = matched_images[0:*maxItems]
	}

	return matched_images, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
