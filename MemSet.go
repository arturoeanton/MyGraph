package main

// Operation is
// MemSet this is struct base of consice set
type MemSet struct {
	mem map[uint]uint64
}

// NewMemSet is contructor
func NewMemSet() MemSet {
	return MemSet{mem: make(map[uint]uint64)}
}

// Size return lengh of MemSet
func (concurrentSet MemSet) Size() uint {
	var count uint
	count = 0
	for k1 := range concurrentSet.mem {
		for i := 0; i < 64; i++ {
			var num uint64 = 1
			if i != 0 {
				num = uint64(2) << uint64(i-1)
			}
			if (num & concurrentSet.mem[uint(k1)]) == num {
				count++
			}
		}
	}
	return count
}

// Add is method for add element
func (concurrentSet MemSet) Add(value uint) bool {
	var row uint
	row = value / 64
	concurrentSet.mem[row] = concurrentSet.mem[row] | bitPosition(value)
	return true
}

// Remove  is method for remove element
func (concurrentSet MemSet) Remove(value uint) bool {
	var row uint
	row = value / 64
	concurrentSet.mem[row] = concurrentSet.mem[row] & ^(bitPosition(value))
	if concurrentSet.mem[row] == 0 {
		delete(concurrentSet.mem, row)
	}
	return true
}

// Contains validate is exist element in set
func (concurrentSet MemSet) Contains(value uint) bool {
	var row uint
	row = value / 64
	var flag uint64
	flag = bitPosition(value)
	return (concurrentSet.mem[row] & flag) == flag
}

// Intersection calulate this operation
func (concurrentSet MemSet) Intersection(set2 MemSet) MemSet {
	result := MemSet{mem: make(map[uint]uint64)}
	for k1 := range concurrentSet.mem {
		result.mem[k1] = concurrentSet.mem[k1] & set2.mem[k1]
	}
	return result
}

// Union calulate this operation
func (concurrentSet MemSet) Union(set2 MemSet) MemSet {
	result := MemSet{mem: make(map[uint]uint64)}

	for k1 := range concurrentSet.mem {
		result.mem[k1] = concurrentSet.mem[k1] | set2.mem[k1]
	}

	for k1 := range set2.mem {
		result.mem[k1] = concurrentSet.mem[k1] | set2.mem[k1]
	}
	return result
}

// SymDiff  calulate Symetric difference  operation
func (concurrentSet MemSet) SymDiff(set2 MemSet) MemSet {
	result := MemSet{mem: make(map[uint]uint64)}
	for k1 := range concurrentSet.mem {
		result.mem[k1] = concurrentSet.mem[k1] ^ set2.mem[k1]
	}
	return result
}

// Diff  calulate difference  operation
func (concurrentSet MemSet) Diff(set2 MemSet) MemSet {
	result := MemSet{mem: make(map[uint]uint64)}
	for k1 := range concurrentSet.mem {
		result.mem[k1] = concurrentSet.mem[k1] & ^set2.mem[k1]
	}
	return result
}

// Array  return array with element of set
func (concurrentSet MemSet) Array() []uint {
	var result []uint
	for k1, value := range concurrentSet.mem {
		for i := 0; i < 64; i++ {
			var num uint64 = 1
			if i != 0 {
				num = uint64(2) << uint64(i-1)
			}
			if (num & value) == num {
				result = append(result, uint(k1*64)+uint(i))
			}
		}
	}
	return result
}

func bitPosition(n uint) uint64 {
	var col uint
	col = n % 64
	if col != 0 {
		return uint64(2) << (col - 1)
	}
	return 1
}