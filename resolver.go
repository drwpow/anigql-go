package ani_gql

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/dangodev/ani-gql/internal/models"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type Resolver struct {
	films *models.FilmConnection
}

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

/* Films (and resolvers) */
type filmResolver struct{ *Resolver }

func parseFilm(data *firestore.DocumentSnapshot) *models.Film {
	film := &models.Film{}
	dataErr := data.DataTo(film)
	if dataErr != nil {
		panic(dataErr)
	}
	return film
}

func (r *queryResolver) Films(ctx context.Context, first int, after *string) (*FilmConnection, error) {
	client, _ := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	defer client.Close()

	query := client.Collection("films").Limit(first)
	if first > 50 {
		query = query.Limit(50)
	}
	if after != nil {
		query = query.StartAfter(after)
	}
	films, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	edges := make([]*FilmEdge, len(films))
	for i, data := range films {
		film := parseFilm(data)
		edges[i] = &FilmEdge{Node: film, Cursor: film.ID}
	}
	pageInfo := &PageInfo{
		EndCursor:   edges[len(edges)-1].Cursor,
		HasNextPage: false,
		StartCursor: edges[0].Cursor,
	}
	return &FilmConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}, nil
}

/* Release */
type releaseResolver struct{ *Resolver }

func (r *releaseResolver) Notes(ctx context.Context, obj *models.Release) (*string, error) {
	return nil, nil
}

/* Studio */
type studioResolver struct{ *Resolver }

func (r *studioResolver) FoundedYear(ctx context.Context, obj *models.Studio) (int, error) {
	return 1960, nil
}

/*
func (r *filmResolver) Artists(ctx context.Context, obj *models.Film) ([]*PersonEdge, error) {
	panic("not implemented")
}
func (r *filmResolver) Composer(ctx context.Context, obj *models.Film) (*models.Person, error) {
	panic("not implemented")
}
func (r *filmResolver) Director(ctx context.Context, obj *models.Film) (*models.Person, error) {
	panic("not implemented")
}
func (r *filmResolver) FrameSequences(ctx context.Context, obj *models.Film) ([]*FrameSequenceEdge, error) {
	panic("not implemented")
}
func (r *filmResolver) KeyframeArtists(ctx context.Context, obj *models.Film) ([]*PersonEdge, error) {
	panic("not implemented")
}
func (r *filmResolver) Releases(ctx context.Context, obj *models.Film) ([]*ReleaseEdge, error) {
	panic("not implemented")
}
func (r *filmResolver) Studio(ctx context.Context, obj *models.Film) (*models.Studio, error) {
	panic("not implemented")
}
func (r *filmResolver) Writers(ctx context.Context, obj *models.Film) ([]*PersonEdge, error) {
	panic("not implemented")
}

func (r *frameResolver) Artists(ctx context.Context, obj *models.Frame) ([]*PersonEdge, error) {
	panic("not implemented")
}
func (r *frameResolver) Order(ctx context.Context, obj *models.Frame) (*int, error) {
	panic("not implemented")
}

func (r *frameSequenceResolver) Artists(ctx context.Context, obj *models.FrameSequence) ([]*PersonEdge, error) {
	panic("not implemented")
}
func (r *frameSequenceResolver) Film(ctx context.Context, obj *models.FrameSequence) (*models.Film, error) {
	panic("not implemented")
}
func (r *frameSequenceResolver) Frames(ctx context.Context, obj *models.FrameSequence) ([]*models.Frame, error) {
	panic("not implemented")
}
func (r *frameSequenceResolver) KeyframeArtists(ctx context.Context, obj *models.FrameSequence) ([]*PersonEdge, error) {
	panic("not implemented")
}

func (r *personResolver) Films(ctx context.Context, obj *models.Person) ([]*FilmEdge, error) {
	panic("not implemented")
}

func (r *releaseResolver) Film(ctx context.Context, obj *models.Release) (*models.Film, error) {
	panic("not implemented")
}
func (r *releaseResolver) Notes(ctx context.Context, obj *models.Release) (*string, error) {
	panic("not implemented")
}


func (r *studioResolver) City(ctx context.Context, obj *models.Studio) (*string, error) {
	panic("not implemented")
}
func (r *studioResolver) Country(ctx context.Context, obj *models.Studio) (string, error) {
	panic("not implemented")
}
func (r *studioResolver) Films(ctx context.Context, obj *models.Studio) ([]*FilmEdge, error) {
	panic("not implemented")
}
func (r *studioResolver) FoundedYear(ctx context.Context, obj *models.Studio) (*string, error) {
	panic("not implemented")
}
func (r *studioResolver) Founders(ctx context.Context, obj *models.Studio) ([]*PersonEdge, error) {
	panic("not implemented")
}
*/
