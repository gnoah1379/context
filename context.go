package context

import (
	"context"
	"time"
)

type Context interface {
	context.Context
	String(key interface{}) string
	Bool(key interface{}) bool
	Bytes(key interface{}) []byte

	Int(key interface{}) int
	Int8(key interface{}) int8
	Int16(key interface{}) int16
	Int32(key interface{}) int32
	Int64(key interface{}) int64

	Uint(key interface{}) uint
	Uint8(key interface{}) uint8
	Uint16(key interface{}) uint16
	Uint32(key interface{}) uint32
	Uint64(key interface{}) uint64

	Float32(key interface{}) float32
	Float64(key interface{}) float64

	Time(key interface{}) time.Time
	Duration(key interface{}) time.Duration

	StringSlice(key interface{}) []string
	IntSlice(key interface{}) []int

	Map(key interface{}) map[interface{}]interface{}
	StringMap(key interface{}) map[string]interface{}
	StringMapString(key interface{}) map[string]string
}
