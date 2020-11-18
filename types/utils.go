package types

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func MatchKind(value interface{}) TypeKind {
	if value == nil {
		return KindMissing
	}

	switch tp := value.(type) {
	case bool:
		return KindBoolean
	case uint8, uint16, uint32, uint64, uint, int8, int16, int32, int64, int:
		return KindInteger
	case float32, float64:
		return KindDecimal
	case string:
		return KindString
	case time.Time:
		return KindDatetime
	default:
		panic(fmt.Sprintf("type: %v isn't supported", tp))
	}
}

func Convert(value interface{}, kind TypeKind) TypedValue {
	v := reflect.ValueOf(value)
	switch tp := value.(type) {
	case bool:
		return Boolean(v.Bool())
	case uint8, uint16, uint32, uint64, uint, int8, int16, int32, int64, int:
		return Integer(v.Int())
	case float32, float64:
		return Decimal(v.Float())
	case string:
		return String(v.String())
	case time.Time:
		return Datetime(value.(time.Time))
	default:
		panic(fmt.Sprintf("type: %v isn't supported", tp))
	}
}

func AutoCast(from TypedValue, to TypeKind) TypedValue {

	if from == nil {
		return Missing
	}

	switch to {
	case KindBoolean:
		var out bool
		from.Cast(&out)
		return Boolean(out)
	case KindInteger:
		var out int64
		from.Cast(&out)
		return Integer(out)
	case KindDecimal:
		var out float64
		from.Cast(&out)
		return Decimal(out)
	case KindDatetime:
		var out time.Time
		from.Cast(&out)
		return Datetime(out)
	case KindString:
		var out string
		from.Cast(&out)
		return String(out)
	case KindMissing:
		return Missing
	default:
		log.Fatalf("unable to auto cast type from: %v to kind: %s", from, to)
		return Missing
	}
}
