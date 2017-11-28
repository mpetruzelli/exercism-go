package house

import (
	"strings"
)

const testVersion = 1

var phrases = []string{
	" the house that Jack built.",
	" the malt\nthat lay in",
	" the rat\nthat ate",
	" the cat\nthat killed",
	" the dog\nthat worried",
	" the cow with the crumpled horn\nthat tossed",
	" the maiden all forlorn\nthat milked",
	" the man all tattered and torn\nthat kissed",
	" the priest all shaven and shorn\nthat married",
	" the rooster that crowed in the morn\nthat woke",
	" the farmer sowing his corn\nthat kept",
	" the horse and the hound and the horn\nthat belonged to",
}

// Song returns the whole stupid song.
func Song() string {
	verses := []string{}
	for i := 1; i <= len(phrases); i++ {
		verses = append(verses, Verse(i))
	}
	retval := strings.Join(verses, "\n\n")
	return retval
}

// Verse returns one of the verses.
func Verse(v int) string {
	verse := "This is"
	for i := v - 1; i >= 0; i-- {
		verse += phrases[i]
	}
	return verse
}
