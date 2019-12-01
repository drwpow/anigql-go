package models

/* A frame sequence is a collection of frames that typically make up a “scene.”
 */
type FrameSequence struct {
	Notes       string `firestore:"notes"`
	Order       int    `firestore:"order"`
	SakugaBooru string `firestore:"sakugaBooru"`
	Video       string `firestore:"video"`
}

/* A frame is a static work of art that when played quickly with the other
frames, make an animated picture. But you probably knew that. A frame must
belong to a FrameSequence (even if we only have data for 1 frame in that
sequence).

A frame may be any speed, and framerate for a film shouldn’t be assumed.
*/
type Frame struct {
	Keyframe bool   `firestore:"keyframe"`
	Notes    string `firestore:"notes"`
	Order    int    `firestore:"order"`
}
