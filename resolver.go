package anigql

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/dangodev/anigql/internal/db"
	"github.com/dangodev/anigql/internal/models"
	"github.com/dangodev/anigql/internal/utils"
	"github.com/imdario/mergo"
)

type Resolver struct {
	addDirectorToFilm *models.Film
	addIntlTitle      *models.Film
	addPerson         *models.Person
	film              *models.Film
	films             []models.Film
	people            []models.Person
	person            *models.Person
}

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

func (r *Resolver) Film() FilmResolver {
	return &filmResolver{r}
}

func (r *Resolver) Frame() FrameResolver {
	return &frameResolver{r}
}
func (r *Resolver) FrameSequence() FrameSequenceResolver {
	return &frameSequenceResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Release() ReleaseResolver {
	return &releaseResolver{r}
}

type frameResolver struct{ *Resolver }

func (r *frameResolver) Image(ctx context.Context, obj *models.Frame) (string, error) {
	panic("not implemented")
}

type frameSequenceResolver struct{ *Resolver }

func (r *frameSequenceResolver) ID(ctx context.Context, obj *models.FrameSequence) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddFilm(ctx context.Context, id string, directorIDs []*string, releaseYear *int, title string) (*models.Film, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	directors := []models.Person{}
	if directorIDs != nil {
		for _, directorID := range directorIDs {
			var director *models.Person
			query, _ := client.Collection("people").Doc(*directorID).Get(ctx)
			query.DataTo(&director)
			if director != nil {
				director.ID = *directorID
				directors = append(directors, *director)
			}
		}
	}

	film := &models.Film{
		Directors:   directors,
		ReleaseYear: releaseYear,
		Title:       title,
	}

	_, mutationErr := client.Collection("films").Doc(id).Set(ctx, film)
	if mutationErr != nil {
		return nil, mutationErr
	}
	film.ID = id

	return film, nil
}

func (r *mutationResolver) AddPerson(ctx context.Context, id string, birthDay *int, birthMonth *int, birthYear *int, deathDay *int, deathMonth *int, deathYear *int, description *string, kanji *string, location *string, name string, surname string, website *string) (*models.Person, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	person := &models.Person{
		BirthDay:    birthDay,
		BirthMonth:  birthMonth,
		BirthYear:   birthYear,
		DeathDay:    deathDay,
		DeathMonth:  deathMonth,
		DeathYear:   deathYear,
		Description: description,
		Kanji:       kanji,
		Location:    location,
		Name:        name,
		Surname:     surname,
		Website:     website,
	}

	_, mutationErr := client.Collection("people").Doc(id).Set(ctx, person)
	if mutationErr != nil {
		return nil, mutationErr
	}
	person.ID = id

	return person, nil
}

func (r *mutationResolver) AddIntlTitle(ctx context.Context, filmID string, title Country) (*models.Film, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	var film *models.Film
	filmQuery, filmErr := client.Collection("films").Doc(filmID).Get(ctx)
	if filmErr != nil {
		return nil, filmErr
	}
	filmQuery.DataTo(&film)
	film.ID = filmID
	mergo.Merge(&film.TitleIntl, title)

	_, mutationErr := client.Collection("films").Doc(filmID).Set(ctx, &models.Film{TitleIntl: film.TitleIntl}, firestore.Merge([]string{"titleIntl"}))
	if mutationErr != nil {
		return nil, mutationErr
	}

	return film, nil
}

func (r *mutationResolver) AddDirectorToFilm(ctx context.Context, personID string, filmID string) (*models.Film, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	// query both
	var person *models.Person
	personQuery, personErr := client.Collection("people").Doc(personID).Get(ctx)
	if personErr != nil {
		return nil, personErr
	}
	personQuery.DataTo(&person)
	person.ID = personID

	var film *models.Film
	filmQuery, filmErr := client.Collection("films").Doc(filmID).Get(ctx)
	if filmErr != nil {
		return nil, filmErr
	}
	filmQuery.DataTo(&film)
	film.ID = filmID

	// set
	directors := []models.Person{*person}
	if len(film.Directors) > 0 {
		alreadyAdded := false
		directors = make([]models.Person, len(film.Directors))
		for i, director := range film.Directors {
			if director.ID == personID {
				alreadyAdded = true
			}
			directors[i] = director
		}
		if alreadyAdded == false {
			directors = append(directors, *person)
		}
	}
	film.Directors = directors
	_, mutationErr := client.Collection("films").Doc(filmID).Set(ctx, film, firestore.Merge([]string{"directors"}))
	if mutationErr != nil {
		return nil, mutationErr
	}
	return film, nil
}

type filmResolver struct{ *Resolver }

func (r *queryResolver) Film(ctx context.Context, id string) (*models.Film, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	var film *models.Film
	query, queryErr := client.Collection("film").Doc(id).Get(ctx)
	if queryErr != nil {
		return nil, queryErr
	}
	query.DataTo(&film)
	film.ID = id

	return film, nil
}

func (r *filmResolver) TitleIntl(ctx context.Context, obj *models.Film) (*TitleIntl, error) {
	return &TitleIntl{
		Au: obj.TitleIntl.Au,
		Br: obj.TitleIntl.Br,
		Ca: obj.TitleIntl.Ca,
		Cn: obj.TitleIntl.Cn,
		Es: obj.TitleIntl.Es,
		Fr: obj.TitleIntl.Fr,
		Gb: obj.TitleIntl.Gb,
		It: obj.TitleIntl.It,
		Jp: obj.TitleIntl.Jp,
		Mx: obj.TitleIntl.Mx,
		Nz: obj.TitleIntl.Nz,
		Tr: obj.TitleIntl.Tr,
		Ua: obj.TitleIntl.Ua,
		Us: obj.TitleIntl.Us,
	}, nil
}

func (r *queryResolver) Films(ctx context.Context, orderBy []*FilmOrderFields, yearStart *int, yearEnd *int) ([]*models.Film, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	// order/filtering
	root := client.Collection("films")
	var query firestore.Query
	if len(orderBy) > 0 {
		for _, order := range orderBy {
			query = root.OrderBy(utils.OrderByName(string(*order)), utils.OrderByDirection(string(*order)))
		}
	} else {
		query = root.OrderBy("releaseYear", firestore.Asc)
	}
	if yearStart != nil && *yearStart > 0 {
		query = root.Where("releaseYear", ">=", yearStart)
	}
	if yearEnd != nil && *yearEnd > 0 {
		query = root.Where("releaseYear", "<=", yearEnd)
	}
	data, queryErr := query.Documents(ctx).GetAll()
	if queryErr != nil {
		return nil, queryErr
	}

	var films []*models.Film
	for _, data := range data {
		var film *models.Film
		parseErr := data.DataTo(&film)
		if parseErr != nil {
			return nil, parseErr
		}
		films = append(films, film)
	}

	return films, nil
}

func (r *queryResolver) People(ctx context.Context, orderBy []*PeopleOrderFields) ([]*models.Person, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	query, queryErr := client.Collection("people").Documents(ctx).GetAll()
	if queryErr != nil {
		return nil, queryErr
	}

	var people []*models.Person
	for _, data := range query {
		var person *models.Person
		parseErr := data.DataTo(&person)
		if parseErr != nil {
			return nil, parseErr
		}
		people = append(people, person)
	}

	return people, nil
}

func (r *queryResolver) Person(ctx context.Context, id string) (*models.Person, error) {
	client, _ := db.Client(ctx)
	defer client.Close()

	var person *models.Person
	query, queryErr := client.Collection("people").Doc(id).Get(ctx)
	if queryErr != nil {
		return nil, queryErr
	}
	query.DataTo(&person)
	person.ID = id

	return person, nil
}

type releaseResolver struct{ *Resolver }

func (r *releaseResolver) ID(ctx context.Context, obj *models.Release) (string, error) {
	panic("not implemented")
}
