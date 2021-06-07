package context

import (
	"context"
	"time"
)

var (
	background = &contextWrapper{context.Background()}
	todo       = &contextWrapper{context.TODO()}
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

func FromContext(ctx context.Context) Context {
	return &contextWrapper{ctx}
}

func WithCancel(parent context.Context) (Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	return &contextWrapper{ctx}, cancel
}

func WithDeadline(parent context.Context, deadline time.Time) (Context, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(parent, deadline)
	return &contextWrapper{ctx}, cancel
}

func WithTimeout(parent context.Context, timeout time.Duration) (Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	return &contextWrapper{ctx}, cancel
}

type contextWrapper struct {
	ctx context.Context
}

func (c *contextWrapper) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *contextWrapper) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *contextWrapper) Err() error {
	return c.ctx.Err()
}

func (c *contextWrapper) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c *contextWrapper) String(key interface{}) (value string) {
	value, _ = c.Value(key).(string)
	return
}

func (c *contextWrapper) Bool(key interface{}) (value bool) {
	value, _ = c.Value(key).(bool)
	return
}

func (c *contextWrapper) Bytes(key interface{}) (value []byte) {
	value, _ = c.Value(key).([]byte)
	return
}

func (c *contextWrapper) Int(key interface{}) (value int) {
	value, _ = c.Value(key).(int)
	return
}

func (c *contextWrapper) Int8(key interface{}) (value int8) {
	value, _ = c.Value(key).(int8)
	return
}

func (c *contextWrapper) Int16(key interface{}) (value int16) {
	value, _ = c.Value(key).(int16)
	return
}

func (c *contextWrapper) Int32(key interface{}) (value int32) {
	value, _ = c.Value(key).(int32)
	return
}

func (c *contextWrapper) Int64(key interface{}) (value int64) {
	value, _ = c.Value(key).(int64)
	return
}

func (c *contextWrapper) Uint(key interface{}) (value uint) {
	value, _ = c.Value(key).(uint)
	return
}

func (c *contextWrapper) Uint8(key interface{}) (value uint8) {
	value, _ = c.Value(key).(uint8)
	return
}

func (c *contextWrapper) Uint16(key interface{}) (value uint16) {
	value, _ = c.Value(key).(uint16)
	return
}

func (c *contextWrapper) Uint32(key interface{}) (value uint32) {
	value, _ = c.Value(key).(uint32)
	return
}

func (c *contextWrapper) Uint64(key interface{}) (value uint64) {
	value, _ = c.Value(key).(uint64)
	return
}

func (c *contextWrapper) Float32(key interface{}) (value float32) {
	value, _ = c.Value(key).(float32)
	return
}

func (c *contextWrapper) Float64(key interface{}) (value float64) {
	value, _ = c.Value(key).(float64)
	return
}

func (c *contextWrapper) Time(key interface{}) (value time.Time) {
	value, _ = c.Value(key).(time.Time)
	return
}

func (c *contextWrapper) Duration(key interface{}) (value time.Duration) {
	value, _ = c.Value(key).(time.Duration)
	return
}

func (c *contextWrapper) StringSlice(key interface{}) (value []string) {
	value, _ = c.Value(key).([]string)
	return
}

func (c *contextWrapper) IntSlice(key interface{}) (value []int) {
	value, _ = c.Value(key).([]int)
	return
}

func (c *contextWrapper) Map(key interface{}) (value map[interface{}]interface{}) {
	value, _ = c.Value(key).(map[interface{}]interface{})
	return
}

func (c *contextWrapper) StringMap(key interface{}) (value map[string]interface{}) {
	value, _ = c.Value(key).(map[string]interface{})
	return
}

func (c *contextWrapper) StringMapString(key interface{}) (value map[string]string) {
	value, _ = c.Value(key).(map[string]string)
	return
}
