package types

import (
	"fmt"
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
