package execution

import (
	"reflect"
	"unsafe"

	"github.com/chewxy/gorgonia/tensor/internal/storage"
	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func (e E) Map(t reflect.Type, fn interface{}, a *storage.Header, incr bool) (err error) {
	as := isScalar(a)
	switch t {
	case Bool:
		var f0 func(bool) bool
		var f1 func(bool) (bool, error)

		switch f := fn.(type) {
		case func(bool) bool:
			f0 = f
		case func(bool) (bool, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Bools()
		if incr {
			return errors.Errorf("Cannot perform increment on t of %v", t)
		}
		switch {
		case as && f0 != nil:
			at[0] = f0(at[0])
		case as && f0 == nil:
			at[0], err = f1(at[0])
		case !as && f0 == nil:
			err = MapErrB(f1, at)
		default:
			MapB(f0, at)
		}
	case Int:
		var f0 func(int) int
		var f1 func(int) (int, error)

		switch f := fn.(type) {
		case func(int) int:
			f0 = f
		case func(int) (int, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Ints()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp int
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrI(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrI(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrI(f1, at)
		default:
			MapI(f0, at)
		}
	case Int8:
		var f0 func(int8) int8
		var f1 func(int8) (int8, error)

		switch f := fn.(type) {
		case func(int8) int8:
			f0 = f
		case func(int8) (int8, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Int8s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp int8
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrI8(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrI8(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrI8(f1, at)
		default:
			MapI8(f0, at)
		}
	case Int16:
		var f0 func(int16) int16
		var f1 func(int16) (int16, error)

		switch f := fn.(type) {
		case func(int16) int16:
			f0 = f
		case func(int16) (int16, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Int16s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp int16
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrI16(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrI16(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrI16(f1, at)
		default:
			MapI16(f0, at)
		}
	case Int32:
		var f0 func(int32) int32
		var f1 func(int32) (int32, error)

		switch f := fn.(type) {
		case func(int32) int32:
			f0 = f
		case func(int32) (int32, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Int32s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp int32
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrI32(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrI32(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrI32(f1, at)
		default:
			MapI32(f0, at)
		}
	case Int64:
		var f0 func(int64) int64
		var f1 func(int64) (int64, error)

		switch f := fn.(type) {
		case func(int64) int64:
			f0 = f
		case func(int64) (int64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Int64s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp int64
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrI64(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrI64(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrI64(f1, at)
		default:
			MapI64(f0, at)
		}
	case Uint:
		var f0 func(uint) uint
		var f1 func(uint) (uint, error)

		switch f := fn.(type) {
		case func(uint) uint:
			f0 = f
		case func(uint) (uint, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Uints()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp uint
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrU(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrU(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrU(f1, at)
		default:
			MapU(f0, at)
		}
	case Uint8:
		var f0 func(uint8) uint8
		var f1 func(uint8) (uint8, error)

		switch f := fn.(type) {
		case func(uint8) uint8:
			f0 = f
		case func(uint8) (uint8, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Uint8s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp uint8
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrU8(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrU8(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrU8(f1, at)
		default:
			MapU8(f0, at)
		}
	case Uint16:
		var f0 func(uint16) uint16
		var f1 func(uint16) (uint16, error)

		switch f := fn.(type) {
		case func(uint16) uint16:
			f0 = f
		case func(uint16) (uint16, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Uint16s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp uint16
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrU16(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrU16(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrU16(f1, at)
		default:
			MapU16(f0, at)
		}
	case Uint32:
		var f0 func(uint32) uint32
		var f1 func(uint32) (uint32, error)

		switch f := fn.(type) {
		case func(uint32) uint32:
			f0 = f
		case func(uint32) (uint32, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Uint32s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp uint32
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrU32(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrU32(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrU32(f1, at)
		default:
			MapU32(f0, at)
		}
	case Uint64:
		var f0 func(uint64) uint64
		var f1 func(uint64) (uint64, error)

		switch f := fn.(type) {
		case func(uint64) uint64:
			f0 = f
		case func(uint64) (uint64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Uint64s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp uint64
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrU64(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrU64(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrU64(f1, at)
		default:
			MapU64(f0, at)
		}
	case Uintptr:
		var f0 func(uintptr) uintptr
		var f1 func(uintptr) (uintptr, error)

		switch f := fn.(type) {
		case func(uintptr) uintptr:
			f0 = f
		case func(uintptr) (uintptr, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Uintptrs()
		if incr {
			return errors.Errorf("Cannot perform increment on t of %v", t)
		}
		switch {
		case as && f0 != nil:
			at[0] = f0(at[0])
		case as && f0 == nil:
			at[0], err = f1(at[0])
		case !as && f0 == nil:
			err = MapErrUintptr(f1, at)
		default:
			MapUintptr(f0, at)
		}
	case Float32:
		var f0 func(float32) float32
		var f1 func(float32) (float32, error)

		switch f := fn.(type) {
		case func(float32) float32:
			f0 = f
		case func(float32) (float32, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Float32s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp float32
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrF32(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrF32(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrF32(f1, at)
		default:
			MapF32(f0, at)
		}
	case Float64:
		var f0 func(float64) float64
		var f1 func(float64) (float64, error)

		switch f := fn.(type) {
		case func(float64) float64:
			f0 = f
		case func(float64) (float64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Float64s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp float64
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrF64(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrF64(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrF64(f1, at)
		default:
			MapF64(f0, at)
		}
	case Complex64:
		var f0 func(complex64) complex64
		var f1 func(complex64) (complex64, error)

		switch f := fn.(type) {
		case func(complex64) complex64:
			f0 = f
		case func(complex64) (complex64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Complex64s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp complex64
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrC64(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrC64(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrC64(f1, at)
		default:
			MapC64(f0, at)
		}
	case Complex128:
		var f0 func(complex128) complex128
		var f1 func(complex128) (complex128, error)

		switch f := fn.(type) {
		case func(complex128) complex128:
			f0 = f
		case func(complex128) (complex128, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Complex128s()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp complex128
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrC128(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrC128(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrC128(f1, at)
		default:
			MapC128(f0, at)
		}
	case String:
		var f0 func(string) string
		var f1 func(string) (string, error)

		switch f := fn.(type) {
		case func(string) string:
			f0 = f
		case func(string) (string, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.Strings()
		switch {
		case as && incr && f0 != nil:
			at[0] += f0(at[0])
		case as && incr && f0 == nil:
			var tmp string
			if tmp, err = f1(at[0]); err != nil {
				return
			}
			at[0] += tmp
		case as && !incr && f0 != nil:
			at[0], err = f1(at[0])
		case as && !incr && f0 == nil:
			at[0], err = f1(at[0])
		case !as && incr && f0 != nil:
			MapIncrStr(f0, at)
		case !as && incr && f0 == nil:
			err = MapIncrErrStr(f1, at)
		case !as && !incr && f0 == nil:
			err = MapErrStr(f1, at)
		default:
			MapStr(f0, at)
		}
	case UnsafePointer:
		var f0 func(unsafe.Pointer) unsafe.Pointer
		var f1 func(unsafe.Pointer) (unsafe.Pointer, error)

		switch f := fn.(type) {
		case func(unsafe.Pointer) unsafe.Pointer:
			f0 = f
		case func(unsafe.Pointer) (unsafe.Pointer, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		at := a.UnsafePointers()
		if incr {
			return errors.Errorf("Cannot perform increment on t of %v", t)
		}
		switch {
		case as && f0 != nil:
			at[0] = f0(at[0])
		case as && f0 == nil:
			at[0], err = f1(at[0])
		case !as && f0 == nil:
			err = MapErrUnsafePointer(f1, at)
		default:
			MapUnsafePointer(f0, at)
		}
	default:
		return errors.Errorf("Cannot map t of %v", t)

	}

	return
}

func (e E) MapIter(t reflect.Type, fn interface{}, a *storage.Header, incr bool, ait Iterator) (err error) {
	switch t {
	case Bool:
		at := a.Bools()
		var f0 func(bool) bool
		var f1 func(bool) (bool, error)

		switch f := fn.(type) {
		case func(bool) bool:
			f0 = f
		case func(bool) (bool, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		if incr {
			return errors.Errorf("Cannot perform increment on t of %v", t)
		}
		switch {
		case f0 == nil:
			err = MapIterErrB(f1, at, ait)
		default:
			MapIterB(f0, at, ait)
		}
	case Int:
		at := a.Ints()
		var f0 func(int) int
		var f1 func(int) (int, error)

		switch f := fn.(type) {
		case func(int) int:
			f0 = f
		case func(int) (int, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrI(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrI(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrI(f1, at, ait)
		default:
			MapIterI(f0, at, ait)
		}
	case Int8:
		at := a.Int8s()
		var f0 func(int8) int8
		var f1 func(int8) (int8, error)

		switch f := fn.(type) {
		case func(int8) int8:
			f0 = f
		case func(int8) (int8, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrI8(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrI8(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrI8(f1, at, ait)
		default:
			MapIterI8(f0, at, ait)
		}
	case Int16:
		at := a.Int16s()
		var f0 func(int16) int16
		var f1 func(int16) (int16, error)

		switch f := fn.(type) {
		case func(int16) int16:
			f0 = f
		case func(int16) (int16, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrI16(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrI16(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrI16(f1, at, ait)
		default:
			MapIterI16(f0, at, ait)
		}
	case Int32:
		at := a.Int32s()
		var f0 func(int32) int32
		var f1 func(int32) (int32, error)

		switch f := fn.(type) {
		case func(int32) int32:
			f0 = f
		case func(int32) (int32, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrI32(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrI32(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrI32(f1, at, ait)
		default:
			MapIterI32(f0, at, ait)
		}
	case Int64:
		at := a.Int64s()
		var f0 func(int64) int64
		var f1 func(int64) (int64, error)

		switch f := fn.(type) {
		case func(int64) int64:
			f0 = f
		case func(int64) (int64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrI64(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrI64(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrI64(f1, at, ait)
		default:
			MapIterI64(f0, at, ait)
		}
	case Uint:
		at := a.Uints()
		var f0 func(uint) uint
		var f1 func(uint) (uint, error)

		switch f := fn.(type) {
		case func(uint) uint:
			f0 = f
		case func(uint) (uint, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrU(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrU(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrU(f1, at, ait)
		default:
			MapIterU(f0, at, ait)
		}
	case Uint8:
		at := a.Uint8s()
		var f0 func(uint8) uint8
		var f1 func(uint8) (uint8, error)

		switch f := fn.(type) {
		case func(uint8) uint8:
			f0 = f
		case func(uint8) (uint8, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrU8(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrU8(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrU8(f1, at, ait)
		default:
			MapIterU8(f0, at, ait)
		}
	case Uint16:
		at := a.Uint16s()
		var f0 func(uint16) uint16
		var f1 func(uint16) (uint16, error)

		switch f := fn.(type) {
		case func(uint16) uint16:
			f0 = f
		case func(uint16) (uint16, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrU16(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrU16(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrU16(f1, at, ait)
		default:
			MapIterU16(f0, at, ait)
		}
	case Uint32:
		at := a.Uint32s()
		var f0 func(uint32) uint32
		var f1 func(uint32) (uint32, error)

		switch f := fn.(type) {
		case func(uint32) uint32:
			f0 = f
		case func(uint32) (uint32, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrU32(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrU32(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrU32(f1, at, ait)
		default:
			MapIterU32(f0, at, ait)
		}
	case Uint64:
		at := a.Uint64s()
		var f0 func(uint64) uint64
		var f1 func(uint64) (uint64, error)

		switch f := fn.(type) {
		case func(uint64) uint64:
			f0 = f
		case func(uint64) (uint64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrU64(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrU64(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrU64(f1, at, ait)
		default:
			MapIterU64(f0, at, ait)
		}
	case Uintptr:
		at := a.Uintptrs()
		var f0 func(uintptr) uintptr
		var f1 func(uintptr) (uintptr, error)

		switch f := fn.(type) {
		case func(uintptr) uintptr:
			f0 = f
		case func(uintptr) (uintptr, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		if incr {
			return errors.Errorf("Cannot perform increment on t of %v", t)
		}
		switch {
		case f0 == nil:
			err = MapIterErrUintptr(f1, at, ait)
		default:
			MapIterUintptr(f0, at, ait)
		}
	case Float32:
		at := a.Float32s()
		var f0 func(float32) float32
		var f1 func(float32) (float32, error)

		switch f := fn.(type) {
		case func(float32) float32:
			f0 = f
		case func(float32) (float32, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrF32(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrF32(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrF32(f1, at, ait)
		default:
			MapIterF32(f0, at, ait)
		}
	case Float64:
		at := a.Float64s()
		var f0 func(float64) float64
		var f1 func(float64) (float64, error)

		switch f := fn.(type) {
		case func(float64) float64:
			f0 = f
		case func(float64) (float64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrF64(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrF64(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrF64(f1, at, ait)
		default:
			MapIterF64(f0, at, ait)
		}
	case Complex64:
		at := a.Complex64s()
		var f0 func(complex64) complex64
		var f1 func(complex64) (complex64, error)

		switch f := fn.(type) {
		case func(complex64) complex64:
			f0 = f
		case func(complex64) (complex64, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrC64(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrC64(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrC64(f1, at, ait)
		default:
			MapIterC64(f0, at, ait)
		}
	case Complex128:
		at := a.Complex128s()
		var f0 func(complex128) complex128
		var f1 func(complex128) (complex128, error)

		switch f := fn.(type) {
		case func(complex128) complex128:
			f0 = f
		case func(complex128) (complex128, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrC128(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrC128(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrC128(f1, at, ait)
		default:
			MapIterC128(f0, at, ait)
		}
	case String:
		at := a.Strings()
		var f0 func(string) string
		var f1 func(string) (string, error)

		switch f := fn.(type) {
		case func(string) string:
			f0 = f
		case func(string) (string, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		switch {
		case incr && f0 != nil:
			MapIterIncrStr(f0, at, ait)
		case incr && f0 == nil:
			err = MapIterIncrErrStr(f1, at, ait)
		case !incr && f0 == nil:
			err = MapIterErrStr(f1, at, ait)
		default:
			MapIterStr(f0, at, ait)
		}
	case UnsafePointer:
		at := a.UnsafePointers()
		var f0 func(unsafe.Pointer) unsafe.Pointer
		var f1 func(unsafe.Pointer) (unsafe.Pointer, error)

		switch f := fn.(type) {
		case func(unsafe.Pointer) unsafe.Pointer:
			f0 = f
		case func(unsafe.Pointer) (unsafe.Pointer, error):
			f1 = f
		default:
			return errors.Errorf("Cannot map fn of %T to array", fn)
		}

		if incr {
			return errors.Errorf("Cannot perform increment on t of %v", t)
		}
		switch {
		case f0 == nil:
			err = MapIterErrUnsafePointer(f1, at, ait)
		default:
			MapIterUnsafePointer(f0, at, ait)
		}
	default:
		return errors.Errorf("Cannot map t of %v", t)
	}

	return
}
