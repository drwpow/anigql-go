package models

/**
 * A person is usually an artist, but can generally be anyone associated with a
 * film.
 */
type Person struct {
	ID          string `firestore:"id"`
	BirthDay    int    `firestore:"birthDay"`
	BirthMonth  int    `firestore:"birthMonth"`
	BirthYear   int    `firestore:"birthYear"`
	DeathDay    int    `firestore:"deathDay"`
	DeathMonth  int    `firestore:"deathMonth"`
	DeathYear   int    `firestore:"deathYear"`
	Description string `firestore:"description"`
	Kanji       string `firestore:"kanji"`
	Location    string `firestore:"location"`
	Name        string `firestore:"name"`
	Surname     string `firestore:"surname"`
	Website     string `firestore:"website"`
}
