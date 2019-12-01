package utils

import (
	"regexp"
	"strings"

	"cloud.google.com/go/firestore"
)

func OrderByDirection(name string) firestore.Direction {
	if strings.HasSuffix(name, "_DESC") {
		return firestore.Desc
	}
	return firestore.Asc
}

func OrderByName(name string) string {
	re := regexp.MustCompile(`(_ASC|_DESC)$`)
	return re.ReplaceAllString(name, "")
}
