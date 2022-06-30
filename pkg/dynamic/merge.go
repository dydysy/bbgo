package dynamic

import "reflect"

// MergeStructValues merges the field value from the source struct to the dest struct.
// Only fields with the same type and the same name will be updated.
func MergeStructValues(dst, src interface{}) {
	if dst == nil {
		return
	}

	rtA := reflect.TypeOf(dst)
	srcStructType := reflect.TypeOf(src)

	rtA = rtA.Elem()
	srcStructType = srcStructType.Elem()

	for i := 0; i < rtA.NumField(); i++ {
		fieldType := rtA.Field(i)
		fieldName := fieldType.Name

		if !fieldType.IsExported() {
			continue
		}

		// if there is a field with the same name
		if fieldSrcType, ok := srcStructType.FieldByName(fieldName); ok {
			// ensure that the type is the same
			if fieldSrcType.Type == fieldType.Type {
				srcValue := reflect.ValueOf(src).Elem().FieldByName(fieldName)
				dstValue := reflect.ValueOf(dst).Elem().FieldByName(fieldName)
				if (fieldType.Type.Kind() == reflect.Ptr && dstValue.IsNil()) || dstValue.IsZero() {
					dstValue.Set(srcValue)
				}
			}
		}
	}
}
