package models

/* Film is difficult. What is a film? If some of a film’s scenes were censored
in a certain country, does that constitute one, or two films? If a director’s
cut was made in addition to the standard theatrical release, is that a
different film or not? What about remasters—how much can change before it’s a
different film (think about the Star Wars remasters)?

The more you think about it, there’s some “fuzziness” around the definition
about what constitutes a work of art. In this context, a “film” is believed
to be the work of art produced by a studio during a limited span of time.

For that reason, “Frames” are considered a core part of the film, but things
like release date and run time are considered a part of a “Release.”
*/
type Film struct {
	ID          string
	Composers   []Person  `firestore:"composers"`
	Directors   []Person  `firestore:"directors"`
	ReleaseYear *int      `firestore:"releaseYear"`
	Title       string    `firestore:"title"`
	TitleIntl   TitleIntl `firestore:"titleIntl"`
	Website     string    `firestore:"website"`
	Writers     []Person  `firestore:"writers"`
}

type TitleIntl struct {
	Au *string `firestore:"au"`
	Br *string `firestore:"br"`
	Ca *string `firestore:"ca"`
	Cn *string `firestore:"cn"`
	Es *string `firestore:"es"`
	Fr *string `firestore:"fr"`
	Gb *string `firestore:"gb"`
	It *string `firestore:"it"`
	Jp *string `firestore:"jp"`
	Mx *string `firestore:"mx"`
	Nz *string `firestore:"nz"`
	Tr *string `firestore:"tr"`
	Ua *string `firestore:"ua"`
	Us string  `firestore:"us"`
}
