// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	"regexp"
	// "strings"

	// "github.com/datacommonsorg/mixer/internal/parser/tmcf"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

const (
	TypeOf = "typeOf"
)

func toBqTable(table string, db string) string {
	return fmt.Sprintf("`%s.%s`", db, table)
}

// QueryOptions contains options for query.
type QueryOptions struct {
	Limit    int
	Db       string
	Prov     bool
	Distinct bool
	Orderby  string
	ASC      bool
}

// Node represents a reference of a graph node in datalog query.
type Node struct {
	// Alias of a node in datalog query, should start with "?".
	Alias string
}

func (n Node) String() string {
	return n.Alias
}

// NewNode creates a new Node instance.
func NewNode(nodeAlias string) Node {
	return Node{nodeAlias}
}

// Table represents a spanner table in through translation process
type Table struct {
	// Spanner table name.
	Name string
	// Id used to distinguish same table when self join.
	ID string
}

func (t Table) String() string {
	return fmt.Sprintf("%s%s", t.Name, t.ID)
}

// Alias gets table's alias used in SQL query.
func (t Table) Alias() string {
	r, _ := regexp.Compile("[.`:-]")
	return r.ReplaceAllString(t.Name, "_") + t.ID
}

// FuncDeps represents the functional deps predicate.
type FuncDeps struct {
}

// Query represents a datalog query statement.
type Query struct {
	// Query predicate is a string of schema.
	Pred string
	// Query subject is a node.
	Sub Node
	// Query object is a node or string.
	Obj interface{}
}

// NewQuery creates a new Query instance.
func NewQuery(pred string, nodeAlias string, obj interface{}) *Query {
	return &Query{Pred: pred, Sub: NewNode(nodeAlias), Obj: obj}
}

// IsTypeOf checks if a query is typeOf statement.
func (q *Query) IsTypeOf() bool {
	return q.Pred == TypeOf
}

// OutArcInfo is used for out arcs pred column.
type OutArcInfo struct {
	Pred   string
	Column string
	IsNode bool
}

// InArcInfo is used for in arcs pred column.
type InArcInfo struct {
	Table  string
	Pred   string
	SubCol string
	ObjCol string
}