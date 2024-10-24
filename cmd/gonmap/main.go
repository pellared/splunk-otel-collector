package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Ullaakut/nmap/v3"
	"go.opentelemetry.io/otel/log"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	lp, err := newLoggerProvider(ctx)
	if err != nil {
		return fmt.Errorf("unable to otel logger provider: %w", err)
	}
	defer lp.Shutdown(context.Background())

	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithCustomArguments(os.Args[1:]...),
	)
	if err != nil {
		return fmt.Errorf("unable to create nmap scanner: %w", err)
	}

	var res *nmap.Run

	if os.Getenv("FAKE") != "" {
		if err := xml.Unmarshal([]byte(example), &res); err != nil {
			return err
		}
	} else {
		var warnings *[]string
		var err error
		res, warnings, err = scanner.Run()
		if len(*warnings) > 0 {
			fmt.Println("run finished with warnings:", *warnings) // Warnings are non-critical errors from nmap.
		}
		if err != nil {
			return fmt.Errorf("unable to run nmap scan: %w", err)
		}

	}

	v, err := convert(res)
	if err != nil {
		return fmt.Errorf("unable to covert result to log body: %w", err)
	}
	fmt.Println(v)

	var r log.Record
	r.AddAttributes(log.String("event.name", "nmap.run"))
	r.SetBody(v)
	lp.Logger("gonmap").Emit(context.Background(), r)
	return nil
}

func convert(i any) (_ log.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	if i == nil {
		return log.Value{}, nil
	}
	v := reflect.ValueOf(i)
	res, _ := convertValue(v)
	return res, nil
}

func convertValue(v reflect.Value) (log.Value, bool) {
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Bool:
		b := v.Bool()
		return log.BoolValue(b), b
	case reflect.String:
		str := v.String()
		return log.StringValue(str), str != ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		return log.Int64Value(i), i != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i := v.Uint()
		return convertUintValue(i), i != 0
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		return log.Float64Value(f), f != 0.0
	case reflect.Struct:
		if t == reflect.TypeOf(nmap.Timestamp{}) {
			nano := time.Time(v.Interface().(nmap.Timestamp)).UnixNano()
			return log.Int64Value(nano), nano != 0
		}

		kvs := make([]log.KeyValue, 0, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			tag := t.Field(i).Tag.Get("xml")
			if tag == "" {
				continue
			}
			key, rest, _ := strings.Cut(tag, ",")
			val, ok := convertValue(v.Field(i))
			if strings.Contains(rest, "omitempty") && !ok {
				continue
			}

			kvs = append(kvs, log.KeyValue{
				Key:   key,
				Value: val,
			})
		}
		if len(kvs) == 0 {
			return log.Value{}, false
		}
		return log.MapValue(kvs...), true
	case reflect.Slice, reflect.Array:
		items := make([]log.Value, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			val, _ := convertValue(v.Index(i))
			items = append(items, val)
		}
		return log.SliceValue(items...), len(items) != 0
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return log.Value{}, false
		}
		return convertValue(v.Elem())
	}

	panic(fmt.Sprintf("unhandled: (%s) %+v", t, v.Interface()))
}

// convertUintValue converts a uint64 to a log.Value.
// If the value is too large to fit in an int64, it is converted to a string.
func convertUintValue(v uint64) log.Value {
	if v > math.MaxInt64 {
		return log.StringValue(strconv.FormatUint(v, 10))
	}
	return log.Int64Value(int64(v))
}
