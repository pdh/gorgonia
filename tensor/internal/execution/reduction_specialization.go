package execution

import (
	"reflect"

	"github.com/chewxy/gorgonia/tensor/internal/storage"
	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func MonotonicSum(t reflect.Type, a *storage.Header) (retVal interface{}, err error) {
	switch t {
	case Int:
		retVal = SumI(a.Ints())
		return
	case Int8:
		retVal = SumI8(a.Int8s())
		return
	case Int16:
		retVal = SumI16(a.Int16s())
		return
	case Int32:
		retVal = SumI32(a.Int32s())
		return
	case Int64:
		retVal = SumI64(a.Int64s())
		return
	case Uint:
		retVal = SumU(a.Uints())
		return
	case Uint8:
		retVal = SumU8(a.Uint8s())
		return
	case Uint16:
		retVal = SumU16(a.Uint16s())
		return
	case Uint32:
		retVal = SumU32(a.Uint32s())
		return
	case Uint64:
		retVal = SumU64(a.Uint64s())
		return
	case Float32:
		retVal = SumF32(a.Float32s())
		return
	case Float64:
		retVal = SumF64(a.Float64s())
		return
	case Complex64:
		retVal = SumC64(a.Complex64s())
		return
	case Complex128:
		retVal = SumC128(a.Complex128s())
		return
	default:
		err = errors.Errorf("Cannot perform Sum on %v", t)
		return
	}
}

func SumMethods(t reflect.Type) (firstFn, lasFn, defaultFn interface{}, err error) {
	switch t {
	case Int:
		return VecAddI, SumI, AddI, nil
	case Int8:
		return VecAddI8, SumI8, AddI8, nil
	case Int16:
		return VecAddI16, SumI16, AddI16, nil
	case Int32:
		return VecAddI32, SumI32, AddI32, nil
	case Int64:
		return VecAddI64, SumI64, AddI64, nil
	case Uint:
		return VecAddU, SumU, AddU, nil
	case Uint8:
		return VecAddU8, SumU8, AddU8, nil
	case Uint16:
		return VecAddU16, SumU16, AddU16, nil
	case Uint32:
		return VecAddU32, SumU32, AddU32, nil
	case Uint64:
		return VecAddU64, SumU64, AddU64, nil
	case Float32:
		return VecAddF32, SumF32, AddF32, nil
	case Float64:
		return VecAddF64, SumF64, AddF64, nil
	case Complex64:
		return VecAddC64, SumC64, AddC64, nil
	case Complex128:
		return VecAddC128, SumC128, AddC128, nil
	default:
		return nil, nil, nil, errors.Errorf("No methods found for Sum for %v", t)
	}
}

func MonotonicMax(t reflect.Type, a *storage.Header) (retVal interface{}, err error) {
	switch t {
	case Int:
		retVal = SliceMaxI(a.Ints())
		return
	case Int8:
		retVal = SliceMaxI8(a.Int8s())
		return
	case Int16:
		retVal = SliceMaxI16(a.Int16s())
		return
	case Int32:
		retVal = SliceMaxI32(a.Int32s())
		return
	case Int64:
		retVal = SliceMaxI64(a.Int64s())
		return
	case Uint:
		retVal = SliceMaxU(a.Uints())
		return
	case Uint8:
		retVal = SliceMaxU8(a.Uint8s())
		return
	case Uint16:
		retVal = SliceMaxU16(a.Uint16s())
		return
	case Uint32:
		retVal = SliceMaxU32(a.Uint32s())
		return
	case Uint64:
		retVal = SliceMaxU64(a.Uint64s())
		return
	case Float32:
		retVal = SliceMaxF32(a.Float32s())
		return
	case Float64:
		retVal = SliceMaxF64(a.Float64s())
		return
	default:
		err = errors.Errorf("Cannot perform Max on %v", t)
		return
	}
}

func MaxMethods(t reflect.Type) (firstFn, lasFn, defaultFn interface{}, err error) {
	switch t {
	case Int:
		return VecMaxI, SliceMaxI, MaxI, nil
	case Int8:
		return VecMaxI8, SliceMaxI8, MaxI8, nil
	case Int16:
		return VecMaxI16, SliceMaxI16, MaxI16, nil
	case Int32:
		return VecMaxI32, SliceMaxI32, MaxI32, nil
	case Int64:
		return VecMaxI64, SliceMaxI64, MaxI64, nil
	case Uint:
		return VecMaxU, SliceMaxU, MaxU, nil
	case Uint8:
		return VecMaxU8, SliceMaxU8, MaxU8, nil
	case Uint16:
		return VecMaxU16, SliceMaxU16, MaxU16, nil
	case Uint32:
		return VecMaxU32, SliceMaxU32, MaxU32, nil
	case Uint64:
		return VecMaxU64, SliceMaxU64, MaxU64, nil
	case Float32:
		return VecMaxF32, SliceMaxF32, MaxF32, nil
	case Float64:
		return VecMaxF64, SliceMaxF64, MaxF64, nil
	default:
		return nil, nil, nil, errors.Errorf("No methods found for Max for %v", t)
	}
}

func MonotonicMin(t reflect.Type, a *storage.Header) (retVal interface{}, err error) {
	switch t {
	case Int:
		retVal = SliceMinI(a.Ints())
		return
	case Int8:
		retVal = SliceMinI8(a.Int8s())
		return
	case Int16:
		retVal = SliceMinI16(a.Int16s())
		return
	case Int32:
		retVal = SliceMinI32(a.Int32s())
		return
	case Int64:
		retVal = SliceMinI64(a.Int64s())
		return
	case Uint:
		retVal = SliceMinU(a.Uints())
		return
	case Uint8:
		retVal = SliceMinU8(a.Uint8s())
		return
	case Uint16:
		retVal = SliceMinU16(a.Uint16s())
		return
	case Uint32:
		retVal = SliceMinU32(a.Uint32s())
		return
	case Uint64:
		retVal = SliceMinU64(a.Uint64s())
		return
	case Float32:
		retVal = SliceMinF32(a.Float32s())
		return
	case Float64:
		retVal = SliceMinF64(a.Float64s())
		return
	default:
		err = errors.Errorf("Cannot perform Min on %v", t)
		return
	}
}

func MinMethods(t reflect.Type) (firstFn, lasFn, defaultFn interface{}, err error) {
	switch t {
	case Int:
		return VecMinI, SliceMinI, MinI, nil
	case Int8:
		return VecMinI8, SliceMinI8, MinI8, nil
	case Int16:
		return VecMinI16, SliceMinI16, MinI16, nil
	case Int32:
		return VecMinI32, SliceMinI32, MinI32, nil
	case Int64:
		return VecMinI64, SliceMinI64, MinI64, nil
	case Uint:
		return VecMinU, SliceMinU, MinU, nil
	case Uint8:
		return VecMinU8, SliceMinU8, MinU8, nil
	case Uint16:
		return VecMinU16, SliceMinU16, MinU16, nil
	case Uint32:
		return VecMinU32, SliceMinU32, MinU32, nil
	case Uint64:
		return VecMinU64, SliceMinU64, MinU64, nil
	case Float32:
		return VecMinF32, SliceMinF32, MinF32, nil
	case Float64:
		return VecMinF64, SliceMinF64, MinF64, nil
	default:
		return nil, nil, nil, errors.Errorf("No methods found for Min for %v", t)
	}
}
