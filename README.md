# AniArchive GQL

Animation database, served in GraphQL via [gqlgen][gqlgen].

## Setup

Copy `.env.example` to `.env` and fill in the values. Then run:

```bash
go run server/server.go
```

## Making edits

To update the generated code, run:

```bash
go run github.com/99designs/gqlgen -v
```

[gqlgen]: https://gqlgen.com/
