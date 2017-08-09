package execution

import "unsafe"

/*
GENERATED FILE. DO NOT EDIT
*/

func GtI(a []int, b []int, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtI8(a []int8, b []int8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtI16(a []int16, b []int16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtI32(a []int32, b []int32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtI64(a []int64, b []int64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtU(a []uint, b []uint, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtU8(a []uint8, b []uint8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtU16(a []uint16, b []uint16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtU32(a []uint32, b []uint32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtU64(a []uint64, b []uint64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtF32(a []float32, b []float32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtF64(a []float64, b []float64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GtStr(a []string, b []string, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] > b[i]
	}
}

func GteI(a []int, b []int, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteI8(a []int8, b []int8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteI16(a []int16, b []int16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteI32(a []int32, b []int32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteI64(a []int64, b []int64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteU(a []uint, b []uint, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteU8(a []uint8, b []uint8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteU16(a []uint16, b []uint16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteU32(a []uint32, b []uint32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteU64(a []uint64, b []uint64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteF32(a []float32, b []float32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteF64(a []float64, b []float64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func GteStr(a []string, b []string, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] >= b[i]
	}
}

func LtI(a []int, b []int, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtI8(a []int8, b []int8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtI16(a []int16, b []int16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtI32(a []int32, b []int32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtI64(a []int64, b []int64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtU(a []uint, b []uint, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtU8(a []uint8, b []uint8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtU16(a []uint16, b []uint16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtU32(a []uint32, b []uint32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtU64(a []uint64, b []uint64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtF32(a []float32, b []float32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtF64(a []float64, b []float64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LtStr(a []string, b []string, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] < b[i]
	}
}

func LteI(a []int, b []int, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteI8(a []int8, b []int8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteI16(a []int16, b []int16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteI32(a []int32, b []int32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteI64(a []int64, b []int64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteU(a []uint, b []uint, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteU8(a []uint8, b []uint8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteU16(a []uint16, b []uint16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteU32(a []uint32, b []uint32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteU64(a []uint64, b []uint64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteF32(a []float32, b []float32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteF64(a []float64, b []float64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func LteStr(a []string, b []string, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] <= b[i]
	}
}

func EqB(a []bool, b []bool, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqI(a []int, b []int, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqI8(a []int8, b []int8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqI16(a []int16, b []int16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqI32(a []int32, b []int32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqI64(a []int64, b []int64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqU(a []uint, b []uint, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqU8(a []uint8, b []uint8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqU16(a []uint16, b []uint16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqU32(a []uint32, b []uint32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqU64(a []uint64, b []uint64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqUintptr(a []uintptr, b []uintptr, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqF32(a []float32, b []float32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqF64(a []float64, b []float64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqC64(a []complex64, b []complex64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqC128(a []complex128, b []complex128, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqStr(a []string, b []string, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func EqUnsafePointer(a []unsafe.Pointer, b []unsafe.Pointer, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] == b[i]
	}
}

func NeB(a []bool, b []bool, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeI(a []int, b []int, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeI8(a []int8, b []int8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeI16(a []int16, b []int16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeI32(a []int32, b []int32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeI64(a []int64, b []int64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeU(a []uint, b []uint, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeU8(a []uint8, b []uint8, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeU16(a []uint16, b []uint16, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeU32(a []uint32, b []uint32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeU64(a []uint64, b []uint64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeUintptr(a []uintptr, b []uintptr, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeF32(a []float32, b []float32, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeF64(a []float64, b []float64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeC64(a []complex64, b []complex64, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeC128(a []complex128, b []complex128, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeStr(a []string, b []string, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func NeUnsafePointer(a []unsafe.Pointer, b []unsafe.Pointer, retVal []bool) {
	a = a[:]
	b = b[:len(a)]
	retVal = retVal[:len(a)]
	for i := range retVal {
		retVal[i] = a[i] != b[i]
	}
}

func GtSameI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameF32(a []float32, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameF64(a []float64, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameStr(a []string, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] > b[i] {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func GteSameI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameF32(a []float32, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameF64(a []float64, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameStr(a []string, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] >= b[i] {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func LtSameI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameF32(a []float32, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameF64(a []float64, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameStr(a []string, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] < b[i] {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func LteSameI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameF32(a []float32, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameF64(a []float64, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameStr(a []string, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] <= b[i] {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func EqSameB(a []bool, b []bool) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = true
		} else {
			a[i] = false
		}
	}
}

func EqSameI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameUintptr(a []uintptr, b []uintptr) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameF32(a []float32, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameF64(a []float64, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameC64(a []complex64, b []complex64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameC128(a []complex128, b []complex128) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameStr(a []string, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] == b[i] {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func NeSameB(a []bool, b []bool) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = true
		} else {
			a[i] = false
		}
	}
}

func NeSameI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameUintptr(a []uintptr, b []uintptr) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameF32(a []float32, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameF64(a []float64, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameC64(a []complex64, b []complex64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameC128(a []complex128, b []complex128) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameStr(a []string, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func GtIterI(a []int, b []int, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterI8(a []int8, b []int8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterI16(a []int16, b []int16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterI32(a []int32, b []int32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterI64(a []int64, b []int64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterU(a []uint, b []uint, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterU8(a []uint8, b []uint8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterU16(a []uint16, b []uint16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterU32(a []uint32, b []uint32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterU64(a []uint64, b []uint64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterF32(a []float32, b []float32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterF64(a []float64, b []float64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GtIterStr(a []string, b []string, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] > b[j]
		}
	}
	return
}

func GteIterI(a []int, b []int, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterI8(a []int8, b []int8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterI16(a []int16, b []int16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterI32(a []int32, b []int32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterI64(a []int64, b []int64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterU(a []uint, b []uint, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterU8(a []uint8, b []uint8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterU16(a []uint16, b []uint16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterU32(a []uint32, b []uint32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterU64(a []uint64, b []uint64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterF32(a []float32, b []float32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterF64(a []float64, b []float64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func GteIterStr(a []string, b []string, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] >= b[j]
		}
	}
	return
}

func LtIterI(a []int, b []int, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterI8(a []int8, b []int8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterI16(a []int16, b []int16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterI32(a []int32, b []int32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterI64(a []int64, b []int64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterU(a []uint, b []uint, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterU8(a []uint8, b []uint8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterU16(a []uint16, b []uint16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterU32(a []uint32, b []uint32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterU64(a []uint64, b []uint64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterF32(a []float32, b []float32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterF64(a []float64, b []float64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LtIterStr(a []string, b []string, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] < b[j]
		}
	}
	return
}

func LteIterI(a []int, b []int, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterI8(a []int8, b []int8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterI16(a []int16, b []int16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterI32(a []int32, b []int32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterI64(a []int64, b []int64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterU(a []uint, b []uint, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterU8(a []uint8, b []uint8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterU16(a []uint16, b []uint16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterU32(a []uint32, b []uint32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterU64(a []uint64, b []uint64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterF32(a []float32, b []float32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterF64(a []float64, b []float64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func LteIterStr(a []string, b []string, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] <= b[j]
		}
	}
	return
}

func EqIterB(a []bool, b []bool, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterI(a []int, b []int, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterI8(a []int8, b []int8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterI16(a []int16, b []int16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterI32(a []int32, b []int32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterI64(a []int64, b []int64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterU(a []uint, b []uint, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterU8(a []uint8, b []uint8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterU16(a []uint16, b []uint16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterU32(a []uint32, b []uint32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterU64(a []uint64, b []uint64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterUintptr(a []uintptr, b []uintptr, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterF32(a []float32, b []float32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterF64(a []float64, b []float64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterC64(a []complex64, b []complex64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterC128(a []complex128, b []complex128, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterStr(a []string, b []string, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func EqIterUnsafePointer(a []unsafe.Pointer, b []unsafe.Pointer, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] == b[j]
		}
	}
	return
}

func NeIterB(a []bool, b []bool, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterI(a []int, b []int, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterI8(a []int8, b []int8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterI16(a []int16, b []int16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterI32(a []int32, b []int32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterI64(a []int64, b []int64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterU(a []uint, b []uint, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterU8(a []uint8, b []uint8, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterU16(a []uint16, b []uint16, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterU32(a []uint32, b []uint32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterU64(a []uint64, b []uint64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterUintptr(a []uintptr, b []uintptr, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterF32(a []float32, b []float32, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterF64(a []float64, b []float64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterC64(a []complex64, b []complex64, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterC128(a []complex128, b []complex128, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterStr(a []string, b []string, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func NeIterUnsafePointer(a []unsafe.Pointer, b []unsafe.Pointer, retVal []bool, ait Iterator, bit Iterator, rit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			retVal[k] = a[i] != b[j]
		}
	}
	return
}

func GtSameIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterStr(a []string, b []string, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] > b[j] {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func GteSameIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterStr(a []string, b []string, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] >= b[j] {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func LtSameIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterStr(a []string, b []string, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] < b[j] {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func LteSameIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterStr(a []string, b []string, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] <= b[j] {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func EqSameIterB(a []bool, b []bool, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = true
			} else {
				a[i] = false
			}
		}
	}
	return
}

func EqSameIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterUintptr(a []uintptr, b []uintptr, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterC64(a []complex64, b []complex64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterC128(a []complex128, b []complex128, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterStr(a []string, b []string, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] == b[j] {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func NeSameIterB(a []bool, b []bool, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = true
			} else {
				a[i] = false
			}
		}
	}
	return
}

func NeSameIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterUintptr(a []uintptr, b []uintptr, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterC64(a []complex64, b []complex64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterC128(a []complex128, b []complex128, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterStr(a []string, b []string, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if a[i] != b[j] {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}
