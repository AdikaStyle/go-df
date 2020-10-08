package aggs

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
)

type Aggregation func(column backend.Column) types.TypedValue
