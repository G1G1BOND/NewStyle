package utils

import "reflect"

func StructToMap(value interface{}) map[string]interface{} {
	v := reflect.Indirect(reflect.ValueOf(value))
	if v.Kind() != reflect.Struct {
		panic("value must be a struct")
	}
	vTyp := v.Type()
	m := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		//匿名字段或非暴露字段
		if !vTyp.Field(i).IsExported() || vTyp.Field(i).Anonymous {
			continue
		}
		fKey := vTyp.Field(i).Name
		var fValue interface{}
		switch field.Kind() {
		case reflect.Struct:
			fValue = StructToMap(field.Interface())
		default:
			fValue = reflect.Indirect(field).Interface()
		}
		m[fKey] = fValue
	}
	return m
}
