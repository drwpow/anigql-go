package models

/* A studio is a dedicated unit of artists that produced a film.
Studios can be long- or short-lived.

A studio may have an alias, or a name it was previously known by.
*/
type Studio struct {
	ID           string
	Aliases      []string `firestore:"aliases"`
	City         string   `firestore:"city"`
	Country      string   `firestore:"country"`
	FoundedDay   int      `firestore:"foundedDay"`
	FoundedMonth int      `firestore:"foundedMonth"`
	FoundedYear  int      `firestore:"foundedYear"`
	Name         string   `firestore:"name"`
	Website      string   `firestore:"website"`
}
