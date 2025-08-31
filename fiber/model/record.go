package model

import (
	"encoding/json"
	"iter"
	"maps"
	"reflect"

	"github.com/bincooo/gptsdk/common"
	"github.com/bincooo/gptsdk/stream"
)

type Record[Key comparable, Value any] map[Key]Value

func (rec Record[Key, Value]) Put(k Key, v Value) {
	rec[k] = v
}

// 获取值
func (rec Record[Key, Value]) Get(k Key) Value {
	return rec[k]
}

// 删除值
func (rec Record[Key, Value]) Del(k Key) {
	delete(rec, k)
}

// 元素数量
func (rec Record[Key, Value]) Len() int {
	return len(rec)
}

// keys 迭代器
func (rec Record[Key, Value]) Keys() iter.Seq[Key] {
	return maps.Keys(rec)
}

// values 迭代器
func (rec Record[Key, Value]) Values() iter.Seq[Value] {
	return maps.Values(rec)
}

// 是否包含 key
func (rec Record[Key, Value]) Contains(k Key) bool {
	value := rec[k]
	return stream.NotNil[Value]()(value)
}

// 深克隆
func (rec Record[Key, Value]) Clone() Record[Key, Value] {
	return common.Copy(rec)
}

// 字符串序列化
func (rec Record[Key, Value]) String() string {
	chunk, err := json.Marshal(rec)
	if err != nil {
		panic(err)
	}
	return string(chunk)
}

// 值比较
func (rec Record[Key, Value]) ValueEqual(k Key, v Value) (ok bool) {
	if !rec.Contains(k) {
		return
	}

	return reflect.DeepEqual(v, rec.Get(k))
}

// 值包含
func (rec Record[Key, Value]) ValueEquals(k Key, values ...Value) (ok bool) {
	if !rec.Contains(k) {
		return
	}

	for _, value := range values {
		if rec.ValueEqual(k, value) {
			return true
		}
	}
	return
}

// 获取值
//
//	@param rec Record实例
//	@param k 实例的key值
func JustValue[Key comparable, Value any](rec Record[Key, any], k Key) (value Value) {
	value, _ = GetValue[Key, Value](rec, k)
	return
}

// 获取值
//
//	@param rec Record实例
//	@param k 实例的key值
func GetValue[Key comparable, Value any](rec Record[Key, any], k Key) (value Value, ok bool) {
	return get[Key, Value](rec, k)
}

func get[Key comparable, Value any](rec Record[Key, any], k Key) (value Value, ok bool) {
	if rec == nil || !rec.Contains(k) {
		return
	}

	value, ok = rec.Get(k).(Value)
	if !ok {
		return
	}

	return
}
