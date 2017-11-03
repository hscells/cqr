package cqr

import (
	"encoding/json"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	c := NewBooleanQuery("or", []CommonQueryRepresentation{
		NewKeyword("abc*", "abstract", "title").SetOption("exploded", true).SetOption("truncated", true),
		NewKeyword("def", "abstract", "title").SetOption("exploded", false),
	}).SetOption("slop", 5)

	b, _ := json.Marshal(c)
	log.Println(string(b))
}
