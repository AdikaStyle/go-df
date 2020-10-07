package types

type MissingType struct{}

func (m MissingType) String() string {
	return "null"
}

func (m MissingType) Equals(other Type) bool {
	if other == nil {
		return true
	} else {
		return false
	}
}

func (m MissingType) Compare(other Type) TypeComparision {
	if other == nil {
		return Equals
	} else {
		return RightIsBigger
	}
}

func (m MissingType) Cast(toPtr interface{}) {
	toPtr = nil
}

func (m MissingType) Kind() TypeKind {
	return Missing
}
