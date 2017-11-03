// Package cqr provides a common query representation for keyword and Boolean queries in go.
package cqr

// CommonQueryRepresentation is the parent type for all subtypes.
type CommonQueryRepresentation interface {
	String() string
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
	return k.QueryString
}

// String computes the string representation of a Boolean query.
func (b BooleanQuery) String() string {
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
