package gmap

import (
	"hash/maphash"
	"strconv"
	"unsafe"
)

type hmap[K comparable, V any] struct {
	count   int
	buckets []bucket[K, V]
	hasher  func(key K) uint64
}

type pair[K comparable, V any] struct {
	key   K
	value V
}

type bucket[K comparable, V any] struct {
	pairs []pair[K, V]
}

func New[K comparable, V any](numberOfElements int) *hmap[K, V] {
	seed := maphash.MakeSeed()

	numberOfBuckets := numberOfElements // TODO optimize

	buckets := make([]bucket[K, V], numberOfBuckets)

	hasher := func(key K) uint64 {
		s := keyToString(key)

		up := unsafe.Pointer(&s)
		sb := *(*[]byte)(up)

		hash := maphash.Bytes(seed, sb)

		return hash
	}

	return &hmap[K, V]{
		buckets: buckets,
		hasher:  hasher,
	}
}

// keyToString converts any of comparable types to string
func keyToString(key interface{}) string {
	switch v := key.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case *int:
		return strconv.FormatInt(int64(*v), 10)
	case *int8:
		return strconv.FormatInt(int64(*v), 10)
	case *int16:
		return strconv.FormatInt(int64(*v), 10)
	case *int32:
		return strconv.FormatInt(int64(*v), 10)
	case *int64:
		return strconv.FormatInt(*v, 10)
	case *uint:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint8:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint16:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint32:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint64:
		return strconv.FormatUint(*v, 10)
	case *float32:
		return strconv.FormatFloat(float64(*v), 'f', -1, 32)
	case *float64:
		return strconv.FormatFloat(*v, 'f', -1, 64)
	case *string:
		return *v
	case *bool:
		return strconv.FormatBool(*v)
	default:
		return "Unsupported type"
	}
}

func (m *hmap[K, V]) Assign(key K, value V) {
	bucketNumber := m.getBucketNumber(key)

	m.count++

	pair := pair[K, V]{
		key:   key,
		value: value,
	}

	m.buckets[bucketNumber].pairs = append(m.buckets[bucketNumber].pairs, pair)
}

func (m *hmap[K, V]) getBucketNumber(key K) uint64 {
	hash := m.hasher(key)

	bucketNumber := hash % uint64(len(m.buckets)) // TODO можно сделать в двоичном виде

	return bucketNumber
}

func (m *hmap[K, V]) Access1(key K) (value V) {
	value, _ = m.Access2(key)
	return value
}

func (m *hmap[K, V]) Access2(key K) (value V, ok bool) {
	bucketNumber := m.getBucketNumber(key)

	bucketPairs := m.buckets[bucketNumber].pairs

	for _, pair := range bucketPairs {
		if pair.key == key {
			return pair.value, true
		}
	}

	return value, false
}

func (m *hmap[K, V]) Len() int {
	return m.count
}

func (m *hmap[K, V]) LoadFactor() float32 {
	return float32(m.count) / float32(len(m.buckets))
}
