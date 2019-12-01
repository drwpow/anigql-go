// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package anigql

import (
	"fmt"
	"io"
	"strconv"
)

type Country struct {
	Au *string `json:"au"`
	Br *string `json:"br"`
	Ca *string `json:"ca"`
	Cn *string `json:"cn"`
	Es *string `json:"es"`
	Fr *string `json:"fr"`
	Gb *string `json:"gb"`
	It *string `json:"it"`
	Jp *string `json:"jp"`
	Mx *string `json:"mx"`
	Nz *string `json:"nz"`
	Tr *string `json:"tr"`
	Ua *string `json:"ua"`
	Us string  `json:"us"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
	StartCursor string `json:"startCursor"`
}

// TitleIntl is a map of 2-letter ISO country codes of all a film’s titles.
type TitleIntl struct {
	Au *string `json:"au"`
	Br *string `json:"br"`
	Ca *string `json:"ca"`
	Cn *string `json:"cn"`
	Es *string `json:"es"`
	Fr *string `json:"fr"`
	Gb *string `json:"gb"`
	It *string `json:"it"`
	Jp *string `json:"jp"`
	Mx *string `json:"mx"`
	Nz *string `json:"nz"`
	Tr *string `json:"tr"`
	Ua *string `json:"ua"`
	Us string  `json:"us"`
}

type Direction string

const (
	DirectionAsc  Direction = "ASC"
	DirectionDesc Direction = "DESC"
)

var AllDirection = []Direction{
	DirectionAsc,
	DirectionDesc,
}

func (e Direction) IsValid() bool {
	switch e {
	case DirectionAsc, DirectionDesc:
		return true
	}
	return false
}

func (e Direction) String() string {
	return string(e)
}

func (e *Direction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Direction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Direction", str)
	}
	return nil
}

func (e Direction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilmOrderFields string

const (
	FilmOrderFieldsReleaseYearAsc  FilmOrderFields = "releaseYear_ASC"
	FilmOrderFieldsReleaseYearDesc FilmOrderFields = "releaseYear_DESC"
	FilmOrderFieldsTitleAsc        FilmOrderFields = "title_ASC"
	FilmOrderFieldsTitleDesc       FilmOrderFields = "title_DESC"
)

var AllFilmOrderFields = []FilmOrderFields{
	FilmOrderFieldsReleaseYearAsc,
	FilmOrderFieldsReleaseYearDesc,
	FilmOrderFieldsTitleAsc,
	FilmOrderFieldsTitleDesc,
}

func (e FilmOrderFields) IsValid() bool {
	switch e {
	case FilmOrderFieldsReleaseYearAsc, FilmOrderFieldsReleaseYearDesc, FilmOrderFieldsTitleAsc, FilmOrderFieldsTitleDesc:
		return true
	}
	return false
}

func (e FilmOrderFields) String() string {
	return string(e)
}

func (e *FilmOrderFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilmOrderFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilmOrderFields", str)
	}
	return nil
}

func (e FilmOrderFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PeopleOrderFields string

const (
	PeopleOrderFieldsBirthDayAsc    PeopleOrderFields = "birthDay_ASC"
	PeopleOrderFieldsBirthDayDesc   PeopleOrderFields = "birthDay_DESC"
	PeopleOrderFieldsBirthMonthAsc  PeopleOrderFields = "birthMonth_ASC"
	PeopleOrderFieldsBirthMonthDesc PeopleOrderFields = "birthMonth_DESC"
	PeopleOrderFieldsBirthYearAsc   PeopleOrderFields = "birthYear_ASC"
	PeopleOrderFieldsBirthYearDesc  PeopleOrderFields = "birthYear_DESC"
	PeopleOrderFieldsDeathDayAsc    PeopleOrderFields = "deathDay_ASC"
	PeopleOrderFieldsDeathDayDesc   PeopleOrderFields = "deathDay_DESC"
	PeopleOrderFieldsDeathMonthAsc  PeopleOrderFields = "deathMonth_ASC"
	PeopleOrderFieldsDeathMonthDesc PeopleOrderFields = "deathMonth_DESC"
	PeopleOrderFieldsDeathYearAsc   PeopleOrderFields = "deathYear_ASC"
	PeopleOrderFieldsDeathYearDesc  PeopleOrderFields = "deathYear_DESC"
	PeopleOrderFieldsNameAsc        PeopleOrderFields = "name_ASC"
	PeopleOrderFieldsNameDesc       PeopleOrderFields = "name_DESC"
	PeopleOrderFieldsSurnameAsc     PeopleOrderFields = "surname_ASC"
	PeopleOrderFieldsSurnameDesc    PeopleOrderFields = "surname_DESC"
)

var AllPeopleOrderFields = []PeopleOrderFields{
	PeopleOrderFieldsBirthDayAsc,
	PeopleOrderFieldsBirthDayDesc,
	PeopleOrderFieldsBirthMonthAsc,
	PeopleOrderFieldsBirthMonthDesc,
	PeopleOrderFieldsBirthYearAsc,
	PeopleOrderFieldsBirthYearDesc,
	PeopleOrderFieldsDeathDayAsc,
	PeopleOrderFieldsDeathDayDesc,
	PeopleOrderFieldsDeathMonthAsc,
	PeopleOrderFieldsDeathMonthDesc,
	PeopleOrderFieldsDeathYearAsc,
	PeopleOrderFieldsDeathYearDesc,
	PeopleOrderFieldsNameAsc,
	PeopleOrderFieldsNameDesc,
	PeopleOrderFieldsSurnameAsc,
	PeopleOrderFieldsSurnameDesc,
}

func (e PeopleOrderFields) IsValid() bool {
	switch e {
	case PeopleOrderFieldsBirthDayAsc, PeopleOrderFieldsBirthDayDesc, PeopleOrderFieldsBirthMonthAsc, PeopleOrderFieldsBirthMonthDesc, PeopleOrderFieldsBirthYearAsc, PeopleOrderFieldsBirthYearDesc, PeopleOrderFieldsDeathDayAsc, PeopleOrderFieldsDeathDayDesc, PeopleOrderFieldsDeathMonthAsc, PeopleOrderFieldsDeathMonthDesc, PeopleOrderFieldsDeathYearAsc, PeopleOrderFieldsDeathYearDesc, PeopleOrderFieldsNameAsc, PeopleOrderFieldsNameDesc, PeopleOrderFieldsSurnameAsc, PeopleOrderFieldsSurnameDesc:
		return true
	}
	return false
}

func (e PeopleOrderFields) String() string {
	return string(e)
}

func (e *PeopleOrderFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PeopleOrderFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PeopleOrderFields", str)
	}
	return nil
}

func (e PeopleOrderFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
