// Package cqr provides a common query representation for keyword and Boolean queries in go.
package cqr

import (
	"fmt"
	"strings"
	"sort"
)

// CommonQueryRepresentation is the parent type for all subtypes.
type CommonQueryRepresentation interface {
	String() string
	StringPretty() string
	GetOption(string) interface{}
	SetOption(string, interface{}) CommonQueryRepresentation
}

// Keyword is a single query expression.
type Keyword struct {
	QueryString string                 `json:"query"`
	Fields      []string               `json:"fields"`
	Options     map[string]interface{} `json:"options"`
}

// BooleanQuery is a nested set of queries, containing either more Boolean queries, or keywords.
type BooleanQuery struct {
	Operator string                      `json:"operator"`
	Children []CommonQueryRepresentation `json:"children"`
	Options  map[string]interface{}      `json:"options"`
}

// String computes the string representation of a keyword.
func (k Keyword) String() string {
	s := make([]string, len(k.Options))
	i := 0
	for k, v := range k.Options {
		s[i] = fmt.Sprintf("%v:%v", k, v)
		i++
	}
	sort.Strings(s)
	return fmt.Sprintf("%v %v {%v}", k.QueryString, k.Fields, strings.Join(s, " "))
}

func (k Keyword) StringPretty() string {
	return k.QueryString
}

// String computes the string representation of a Boolean query.
func (b BooleanQuery) String() (s string) {
	s += fmt.Sprintf(" ( %v[%v]", b.Operator, b.Options)
	for _, child := range b.Children {
		s += fmt.Sprintf(" %v", child.String())
	}
	s += ") "
	return strings.TrimSpace(s)
}

// String computes the string representation of a Boolean query.
func (b BooleanQuery) StringPretty() (s string) {
	return b.Operator
}

// SetOption sets an optional parameter on the keyword.
func (k Keyword) SetOption(key string, value interface{}) CommonQueryRepresentation {
	k.Options[key] = value
	return k
}

// SetOption sets an optional parameter on the Boolean query.
func (b BooleanQuery) SetOption(key string, value interface{}) CommonQueryRepresentation {
	b.Options[key] = value
	return b
}

// GetOption gets an optional parameter of the keyword.
func (k Keyword) GetOption(key string) interface{} {
	return k.Options[key]
}

// GetOption gets an optional parameter of the Boolean Query.
func (b BooleanQuery) GetOption(key string) interface{} {
	return b.Options[key]
}

// NewKeyword constructs a new keyword.
func NewKeyword(queryString string, fields ...string) Keyword {
	return Keyword{
		QueryString: queryString,
		Fields:      fields,
		Options:     map[string]interface{}{},
	}
}

// NewBooleanQuery constructs a new Boolean query.
func NewBooleanQuery(operator string, children []CommonQueryRepresentation) BooleanQuery {
	return BooleanQuery{
		Operator: operator,
		Children: children,
		Options:  map[string]interface{}{},
	}
}

func IsBoolean(query CommonQueryRepresentation) bool {
	if _, ok := query.(BooleanQuery); ok {
		return true
	}
	return false
}

func CopyKeyword(query Keyword) Keyword {
	fields := make([]string, len(query.Fields))
	copy(fields, query.Fields)
	nq := NewKeyword(query.QueryString, fields...)
	for k, v := range query.Options {
		nq.Options[k] = v
	}
	return nq
}
