package nmapreceiver

// func convert(i any) (_ plog.LogRecord, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			err = fmt.Errorf("%v", r)
// 		}
// 	}()

// 	if i == nil {
// 		return plog, nil
// 	}
// 	v := reflect.ValueOf(i)
// 	res, _ := convertValue(v)
// 	return res, nil
// }

// func convertValue(v reflect.Value) (pcommon.Value, bool) {
// 	t := v.Type()
// 	k := t.Kind()
// 	switch k {
// 	case reflect.Bool:
// 		b := v.Bool()
// 		return log.BoolValue(b), b
// 	case reflect.String:
// 		str := v.String()
// 		return log.StringValue(str), str != ""
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		i := v.Int()
// 		return log.Int64Value(i), i != 0
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		i := v.Uint()
// 		return convertUintValue(i), i != 0
// 	case reflect.Float32, reflect.Float64:
// 		f := v.Float()
// 		return log.Float64Value(f), f != 0.0
// 	case reflect.Struct:
// 		if t == reflect.TypeOf(nmap.Timestamp{}) {
// 			nano := time.Time(v.Interface().(nmap.Timestamp)).UnixNano()
// 			return log.Int64Value(nano), nano != 0
// 		}

// 		kvs := make([]log.KeyValue, 0, v.NumField())
// 		for i := 0; i < v.NumField(); i++ {
// 			tag := t.Field(i).Tag.Get("xml")
// 			if tag == "" {
// 				continue
// 			}
// 			key, rest, _ := strings.Cut(tag, ",")
// 			val, ok := convertValue(v.Field(i))
// 			if strings.Contains(rest, "omitempty") && !ok {
// 				continue
// 			}

// 			kvs = append(kvs, log.KeyValue{
// 				Key:   key,
// 				Value: val,
// 			})
// 		}
// 		if len(kvs) == 0 {
// 			return log.Value{}, false
// 		}
// 		return log.MapValue(kvs...), true
// 	case reflect.Slice, reflect.Array:
// 		items := make([]log.Value, 0, v.Len())
// 		for i := 0; i < v.Len(); i++ {
// 			val, _ := convertValue(v.Index(i))
// 			items = append(items, val)
// 		}
// 		return log.SliceValue(items...), len(items) != 0
// 	case reflect.Ptr, reflect.Interface:
// 		if v.IsNil() {
// 			return log.Value{}, false
// 		}
// 		return convertValue(v.Elem())
// 	}

// 	panic(fmt.Sprintf("unhandled: (%s) %+v", t, v.Interface()))
// }