package context

import (
	"context"
	"time"
)

type KV struct {
	Key   interface{}
	Value interface{}
}

func KeyValue(key interface{}, value interface{}) KV {
	return KV{key, value}
}

func WithValues(parent context.Context, values ...KV) Context {
	mapValues := make(map[interface{}]interface{})
	for _, kv := range values {
		mapValues[kv.Key] = kv.Value
	}
	return &contextValues{parent: parent, values: mapValues}
}

type contextValues struct {
	parent context.Context
	values map[interface{}]interface{}
}

func (c *contextValues) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

func (c *contextValues) Done() <-chan struct{} {
	return c.parent.Done()
}

func (c *contextValues) Err() error {
	return c.parent.Err()
}

func (c *contextValues) Value(key interface{}) interface{} {
	value, ok := c.values[key]
	if !ok {
		return c.parent.Value(key)
	}
	return value
}

func (c *contextValues) String(key interface{}) (value string) {
	value, _ = c.Value(key).(string)
	return
}

func (c *contextValues) Bool(key interface{}) (value bool) {
	value, _ = c.Value(key).(bool)
	return
}

func (c *contextValues) Bytes(key interface{}) (value []byte) {
	value, _ = c.Value(key).([]byte)
	return
}

func (c *contextValues) Int(key interface{}) (value int) {
	value, _ = c.Value(key).(int)
	return
}

func (c *contextValues) Int8(key interface{}) (value int8) {
	value, _ = c.Value(key).(int8)
	return
}

func (c *contextValues) Int16(key interface{}) (value int16) {
	value, _ = c.Value(key).(int16)
	return
}

func (c *contextValues) Int32(key interface{}) (value int32) {
	value, _ = c.Value(key).(int32)
	return
}

func (c *contextValues) Int64(key interface{}) (value int64) {
	value, _ = c.Value(key).(int64)
	return
}

func (c *contextValues) Uint(key interface{}) (value uint) {
	value, _ = c.Value(key).(uint)
	return
}

func (c *contextValues) Uint8(key interface{}) (value uint8) {
	value, _ = c.Value(key).(uint8)
	return
}

func (c *contextValues) Uint16(key interface{}) (value uint16) {
	value, _ = c.Value(key).(uint16)
	return
}

func (c *contextValues) Uint32(key interface{}) (value uint32) {
	value, _ = c.Value(key).(uint32)
	return
}

func (c *contextValues) Uint64(key interface{}) (value uint64) {
	value, _ = c.Value(key).(uint64)
	return
}

func (c *contextValues) Float32(key interface{}) (value float32) {
	value, _ = c.Value(key).(float32)
	return
}

func (c *contextValues) Float64(key interface{}) (value float64) {
	value, _ = c.Value(key).(float64)
	return
}

func (c *contextValues) Time(key interface{}) (value time.Time) {
	value, _ = c.Value(key).(time.Time)
	return
}

func (c *contextValues) Duration(key interface{}) (value time.Duration) {
	value, _ = c.Value(key).(time.Duration)
	return
}

func (c *contextValues) StringSlice(key interface{}) (value []string) {
	value, _ = c.Value(key).([]string)
	return
}

func (c *contextValues) IntSlice(key interface{}) (value []int) {
	value, _ = c.Value(key).([]int)
	return
}

func (c *contextValues) Map(key interface{}) (value map[interface{}]interface{}) {
	value, _ = c.Value(key).(map[interface{}]interface{})
	return
}

func (c *contextValues) StringMap(key interface{}) (value map[string]interface{}) {
	value, _ = c.Value(key).(map[string]interface{})
	return
}

func (c *contextValues) StringMapString(key interface{}) (value map[string]string) {
	value, _ = c.Value(key).(map[string]string)
	return
}
