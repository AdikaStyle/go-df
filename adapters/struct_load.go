package adapters

import (
	"fmt"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"reflect"
)

func LoadStructs(structsArray interface{}, tag string) dataframe.Dataframe {
	structs := reflect.ValueOf(structsArray)
	headers := analyzeStruct(structs.Index(0).Type(), tag)
	be := backend.New(headers)
	populateStructs(be, structs, tag)
	return dataframe.New(be)
}

func populateStructs(be backend.Backend, structs reflect.Value, tag string) {
	len := structs.Len()
	for i := 0; i < len; i++ {
		val := structs.Index(i)
		be.AppendRow(structToRow(val, tag))
	}
}

func structToRow(valOf reflect.Value, tag string) backend.Row {
	row := make(backend.Row)
	for i := 0; i < valOf.NumField(); i++ {
		key := extractFieldName(valOf.Type().Field(i), tag)
		kind := types.MatchKind(valOf.Field(i).Interface())
		val := types.Convert(valOf.Field(i).Interface(), kind)
		row[key] = val
	}
	return row
}

func analyzeStruct(tpe reflect.Type, tag string) backend.Headers {
	var out backend.Headers
	fieldsCount := tpe.NumField()

	for i := 0; i < fieldsCount; i++ {
		field := tpe.Field(i)
		tagVal := extractFieldName(field, tag)

		if tagVal == "-" {
			continue
		}

		out = append(out, backend.Header{
			Name:    tagVal,
			Kind:    goTypeToKind(field.Type),
			Visible: true,
		})
	}

	return out
}

func extractFieldName(field reflect.StructField, tag string) string {
	tagVal, found := field.Tag.Lookup(tag)
	if !found {
		tagVal = field.Name
	}

	return tagVal
}

func goTypeToKind(tpe reflect.Type) types.TypeKind {
	if tpe.Name() == "Time" {
		return types.KindDatetime
	}

	switch tpe.Kind() {
	case reflect.Bool:
		return types.KindBoolean
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return types.KindInteger
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return types.KindInteger
	case reflect.Float32, reflect.Float64:
		return types.KindDecimal
	case reflect.String:
		return types.KindString
	default:
		panic(fmt.Errorf("failed to convert go type: %d to dataframe type", tpe.Kind()))
	}
}
