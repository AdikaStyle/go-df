package types

type MissingValue struct{}

var Missing MissingValue

func (m MissingValue) String() string {
	return "null"
}

func (m MissingValue) Equals(other TypedValue) bool {
	if other == nil {
		return true
	} else {
		return false
	}
}

func (m MissingValue) Compare(other TypedValue) TypeComparision {
	if other == nil {
		return Equals
	} else {
		return RightIsBigger
	}
}

func (m MissingValue) Cast(toPtr interface{}) {
	toPtr = nil
}

func (m MissingValue) Kind() TypeKind {
	return KindMissing
}

func (m MissingValue) Ptr() TypedValue {
	return nil
}

func (b MissingValue) NativeType() interface{} {
	return nil
}

func (m MissingValue) Precedence() int {
	return -1
}
