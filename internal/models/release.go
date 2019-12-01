package models

/* A release can be anything from an original theatrical release, to a
director’s cut, to a DVD re-master re-release.

When My Neighbor Totoro was released in Hong Kong, it was 4 minutes shorter
than the original Japanese theatrical release. From Simpsons to Star Wars,
fans will tell you that with each new format comes some deviation from the
original work of art. While there’s no judgment on whether some releases are
“better” or “worse,” it still stands that releases constitute different
versions of a film. This model keeps track of that complexity.

It is possible for a film not to have any releases, too. Either that means
the film is in progress, or it was a notable film that was cancelled.
*/
type Release struct {
	Area         string `firestore:"area"`
	Format       string `firestore:"format"`
	ReleaseDay   int    `firestore:"releaseDay"`
	ReleaseMonth int    `firestore:"releaseMonth"`
	ReleaseYear  int    `firestore:"releaseYear"`
	Runtime      int    `firestore:"runtime"`
}
