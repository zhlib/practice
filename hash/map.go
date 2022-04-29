package hash

import "unsafe"

type DictEntry struct {
	Key  interface{}
	V    interface{}
	next *DictEntry
}

type DictType struct {
	HashFunction  func(key unsafe.Pointer) uint64
	KeyDup        func(privdate unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer
	ValDup        func(privdata unsafe.Pointer, obj unsafe.Pointer) unsafe.Pointer
	KeyCompare    func(privdate unsafe.Pointer, key1, key2 unsafe.Pointer) int
	KeyDestructor func(privdate, key unsafe.Pointer)
	valDestructor func(privdata, obj unsafe.Pointer)
}

type Dictht struct {
	Table    **DictEntry
	Size     int64
	SizeMask int64
	Used     int64
}

type Dict struct {
	Type      *DictType
	PrivData  *interface{}
	Ht        [2]Dictht
	ReHashIDX int64 /* rehashing not in progress if rehashidx == -1 */
	Iterators int64 /* number of iterators currently running */
}

func DictCreate(t *DictType, privDataPtr *interface{}) (Dict, error) {
	var resp Dict
	resp.Type = t
	resp.PrivData = privDataPtr
	resp.ReHashIDX = -1
	resp.Iterators = 0
	return resp, nil
}

func dictSdsHash(key unsafe.Pointer) uint64 {
	return func(key string) uint64 {
		return uint64(MurmurHash2([]byte(key)))
	}(*(*string)(key))
}

func dictSdsKeyCompare(privdata unsafe.Pointer, key1, key2 unsafe.Pointer) int {
	k1, k2 := *(*string)(key1), *(*string)(key2)
	l1, l2 := len(k1), len(k2)
	if l1 != l2 {
		return 0
	}
	if k1 > k2 {
		return 1
	} else if k1 == k2 {
		return 0
	} else {
		return -1
	}
}

func dictRedisObjectDestructor(privdate, val unsafe.Pointer) {

}

func dictSdsDestructor(privdate, val unsafe.Pointer) {

}

var normalDict = DictType{
	HashFunction:  dictSdsHash,
	KeyDup:        nil,
	ValDup:        nil,
	KeyCompare:    dictSdsKeyCompare,
	KeyDestructor: dictSdsDestructor,
	valDestructor: dictRedisObjectDestructor,
}
