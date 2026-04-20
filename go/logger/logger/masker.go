package logger

import (
	"encoding"
	"encoding/json"
	"reflect"
	"strings"
)

func normalizeMaskFields(fieldMap map[string]any) map[string]any {
	if len(fieldMap) == 0 {
		return nil
	}

	normalized := make(map[string]any, len(fieldMap))
	for key, value := range fieldMap {
		normalized[strings.ToLower(strings.TrimSpace(key))] = value
	}

	return normalized
}

func (l *Logger) maskArgs(args []any) []any {
	if !l.maskingEnabled || len(l.maskFields) == 0 || len(args) == 0 {
		return args
	}

	masked := make([]any, len(args))
	for index, arg := range args {
		masked[index] = l.maskValue(arg)
	}

	return masked
}

func (l *Logger) maskValue(value any) any {
	if !l.maskingEnabled || len(l.maskFields) == 0 {
		return value
	}

	return l.maskReflectValue(reflect.ValueOf(value))
}

func (l *Logger) maskReflectValue(value reflect.Value) any {
	if !value.IsValid() {
		return nil
	}

	if replacement, ok := l.maskForType(value.Type()); ok {
		return replacement
	}

	for value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}

		value = value.Elem()
		if replacement, ok := l.maskForType(value.Type()); ok {
			return replacement
		}
	}

	if preserveStructuredValue(value.Type()) {
		return value.Interface()
	}

	switch value.Kind() {
	case reflect.Struct:
		return l.maskStruct(value)
	case reflect.Map:
		return l.maskMap(value)
	case reflect.Slice, reflect.Array:
		masked := make([]any, value.Len())
		for index := 0; index < value.Len(); index++ {
			masked[index] = l.maskReflectValue(value.Index(index))
		}
		return masked
	default:
		return value.Interface()
	}
}

func (l *Logger) maskStruct(value reflect.Value) map[string]any {
	masked := make(map[string]any)
	valueType := value.Type()

	for index := 0; index < value.NumField(); index++ {
		field := valueType.Field(index)
		if !field.IsExported() {
			continue
		}

		fieldName, include := serializedFieldName(field)
		if !include {
			continue
		}

		fieldValue := value.Field(index)
		if replacement, ok := l.maskForField(field.Name, fieldName); ok {
			masked[fieldName] = replacement
			continue
		}

		masked[fieldName] = l.maskReflectValue(fieldValue)
	}

	return masked
}

func (l *Logger) maskMap(value reflect.Value) any {
	if value.IsNil() {
		return nil
	}

	if value.Type().Key().Kind() != reflect.String {
		return value.Interface()
	}

	masked := make(map[string]any, value.Len())
	iter := value.MapRange()
	for iter.Next() {
		key := iter.Key().String()
		if replacement, ok := l.maskForField(key); ok {
			masked[key] = replacement
			continue
		}

		masked[key] = l.maskReflectValue(iter.Value())
	}

	return masked
}

func (l *Logger) maskForType(valueType reflect.Type) (any, bool) {
	return l.lookupMask(valueType.Name())
}

func (l *Logger) maskForField(names ...string) (any, bool) {
	for _, name := range names {
		if replacement, ok := l.lookupMask(name); ok {
			return replacement, true
		}
	}

	return nil, false
}

func (l *Logger) lookupMask(name string) (any, bool) {
	if name == "" {
		return nil, false
	}

	replacement, ok := l.maskFields[strings.ToLower(strings.TrimSpace(name))]
	if !ok {
		return nil, false
	}

	if replacement == nil {
		return defaultMaskValue, true
	}

	return replacement, true
}

func serializedFieldName(field reflect.StructField) (string, bool) {
	jsonTag := field.Tag.Get("json")
	if jsonTag == "-" {
		return "", false
	}

	if jsonTag == "" {
		return field.Name, true
	}

	name := strings.Split(jsonTag, ",")[0]
	if name == "" {
		return field.Name, true
	}

	return name, true
}

func preserveStructuredValue(valueType reflect.Type) bool {
	jsonMarshalerType := reflect.TypeFor[json.Marshaler]()
	textMarshalerType := reflect.TypeFor[encoding.TextMarshaler]()

	if valueType.Implements(jsonMarshalerType) || reflect.PointerTo(valueType).Implements(jsonMarshalerType) {
		return true
	}

	if valueType.Implements(textMarshalerType) || reflect.PointerTo(valueType).Implements(textMarshalerType) {
		return true
	}

	return false
}
