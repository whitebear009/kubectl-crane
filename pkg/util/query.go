package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type Field string
type Value string

const (
	FieldName      = "name"
	FieldNamespace = "namespace"
)

// Query represents api search terms
type Query struct {
	Filters map[Field]Value

	LabelSelector string
}

type Filter struct {
	Field Field
	Value Value
}

func NewQuery() *Query {
	return &Query{
		Filters: map[Field]Value{},
	}
}

func ObjectMetaFilter(item metav1.ObjectMeta, filter Filter) bool {
	switch filter.Field {
	case FieldName:
		return strings.Contains(item.Name, string(filter.Value))
	case FieldNamespace:
		return strings.Compare(item.Namespace, string(filter.Value)) == 0
	default:
		return false
	}
}
