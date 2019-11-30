// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package ani_gql

import (
	"github.com/dangodev/ani-gql/internal/models"
)

type FilmConnection struct {
	Edges    []*FilmEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type FilmEdge struct {
	Cursor string       `json:"cursor"`
	Node   *models.Film `json:"node"`
}

type FrameSequenceConnection struct {
	Edges    []*FrameSequenceEdge `json:"edges"`
	PageInfo *PageInfo            `json:"pageInfo"`
}

type FrameSequenceEdge struct {
	Cursor string                `json:"cursor"`
	Node   *models.FrameSequence `json:"node"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
	StartCursor string `json:"startCursor"`
}

type PersonConnection struct {
	Edges    []*PersonEdge `json:"edges"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type PersonEdge struct {
	Cursor string         `json:"cursor"`
	Node   *models.Person `json:"node"`
}

type ReleaseConnection struct {
	Edges    []*ReleaseEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type ReleaseEdge struct {
	Cursor string          `json:"cursor"`
	Node   *models.Release `json:"node"`
}

type StudioConnection struct {
	Edges    []*StudioEdge `json:"edges"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type StudioEdge struct {
	Cursor string         `json:"cursor"`
	Node   *models.Studio `json:"node"`
}