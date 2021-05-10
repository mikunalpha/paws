// Package paws provides some helper functions for gopher's developing experience.
package paws

import "time"

// String returns a new pointer of given s string.
func String(s string) *string { return &s }

// MustString returns a string value of pointer p, if p is not nil.
// s is returned, if  p is nil.
func MustString(p *string, s string) string {
	if p != nil {
		return *p
	}
	return s
}

// Z2NString returns a new pointer of given s string.
// If s is zero value, it returns nil.
func Z2NString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// Bool returns a new pointer of given b bool.
func Bool(b bool) *bool { return &b }

// MustBool returns a bool value of pointer p, if p is not nil.
// b is returned, if  p is nil.
func MustBool(p *bool, b bool) bool {
	if p != nil {
		return *p
	}
	return b
}

// Z2NBool returns a new pointer of given b bool.
// If b is zero value, it returns nil.
func Z2NBool(b bool) *bool {
	if b == false {
		return nil
	}
	return &b
}

// Int returns a new pointer of given i int.
func Int(i int) *int { return &i }

// MustInt returns a int value of pointer p, if p is not nil.
// i is returned, if  p is nil.
func MustInt(p *int, i int) int {
	if p != nil {
		return *p
	}
	return i
}

// Z2NInt returns a new pointer of given i int.
// If i is zero value, it returns nil.
func Z2NInt(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

// Int8 returns a new pointer of given i int8.
func Int8(i int8) *int8 { return &i }

// MustInt8 returns int value of pointer p, if p is not nil.
// i is returned, if  p is nil.
func MustInt8(p *int8, i int8) int8 {
	if p != nil {
		return *p
	}
	return i
}

// Z2NInt8 returns a new pointer of given i int8.
// If i is zero value, it returns nil.
func Z2NInt8(i int8) *int8 {
	if i == 0 {
		return nil
	}
	return &i
}

// Int16 returns a new pointer of given i int16.
func Int16(i int16) *int16 { return &i }

// MustInt16 returns a int value of pointer p, if p is not nil.
// i is returned, if  p is nil.
func MustInt16(p *int16, i int16) int16 {
	if p != nil {
		return *p
	}
	return i
}

// Z2NInt16 returns a new pointer of given i int16.
// If i is zero value, it returns nil.
func Z2NInt16(i int16) *int16 {
	if i == 0 {
		return nil
	}
	return &i
}

// Int32 returns a new pointer of given i int32.
func Int32(i int32) *int32 { return &i }

// MustInt32 returns a int value of pointer p, if p is not nil.
// i is returned, if  p is nil.
func MustInt32(p *int32, i int32) int32 {
	if p != nil {
		return *p
	}
	return i
}

// Z2NInt32 returns a new pointer of given i int32.
// If i is zero value, it returns nil.
func Z2NInt32(i int32) *int32 {
	if i == 0 {
		return nil
	}
	return &i
}

// Int64 returns a new pointer of given i int64.
func Int64(i int64) *int64 { return &i }

// MustInt64 returns a int value of pointer p, if p is not nil.
// i is returned, if  p is nil.
func MustInt64(p *int64, i int64) int64 {
	if p != nil {
		return *p
	}
	return i
}

// Z2NInt64 returns a new pointer of given i int64.
// If u is zero value, it returns nil.
func Z2NInt64(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}

// Uint returns a new pointer of given u uint.
func Uint(u uint) *uint { return &u }

// MustUint returns a uint64 value of pointer p, if p is not nil.
// u is returned, if  p is nil.
func MustUint(p *uint, u uint) uint {
	if p != nil {
		return *p
	}
	return u
}

// Z2NUint returns a new pointer of given u uint.
// If u is zero value, it returns nil.
func Z2NUint(u uint) *uint {
	if u == 0 {
		return nil
	}
	return &u
}

// Uint8 returns a new pointer of given u uint8.
func Uint8(u uint8) *uint8 { return &u }

// MustUint8 returns a uint8 value of pointer p, if p is not nil.
// u is returned, if  p is nil.
func MustUint8(p *uint8, u uint8) uint8 {
	if p != nil {
		return *p
	}
	return u
}

// Z2NUint8 returns a new pointer of given u uint8.
// If u is zero value, it returns nil.
func Z2NUint8(u uint8) *uint8 {
	if u == 0 {
		return nil
	}
	return &u
}

// Uint16 returns a new pointer of given u uint16.
func Uint16(u uint16) *uint16 { return &u }

// MustUint16 returns a uint16 value of pointer p, if p is not nil.
// u is returned, if  p is nil.
func MustUint16(p *uint16, u uint16) uint16 {
	if p != nil {
		return *p
	}
	return u
}

// Z2NUint16 returns a new pointer of given u uint16.
// If u is zero value, it returns nil.
func Z2NUint16(u uint16) *uint16 {
	if u == 0 {
		return nil
	}
	return &u
}

// Uint32 returns a new pointer of given u uint32.
func Uint32(u uint32) *uint32 { return &u }

// MustUint32 returns a uint32 value of pointer p, if p is not nil.
// u is returned, if  p is nil.
func MustUint32(p *uint32, u uint32) uint32 {
	if p != nil {
		return *p
	}
	return u
}

// Z2NUint32 returns a new pointer of given u uint32.
// If u is zero value, it returns nil.
func Z2NUint32(u uint32) *uint32 {
	if u == 0 {
		return nil
	}
	return &u
}

// Uint64 returns a new pointer of given u uint64.
func Uint64(u uint64) *uint64 { return &u }

// MustUint64 returns a uint64 value of pointer p, if p is not nil.
// u is returned, if  p is nil.
func MustUint64(p *uint64, u uint64) uint64 {
	if p != nil {
		return *p
	}
	return u
}

// Z2NUint64 returns a new pointer of given u uint64.
// If u is zero value, it returns nil.
func Z2NUint64(u uint64) *uint64 {
	if u == 0 {
		return nil
	}
	return &u
}

// Float32 returns a new pointer of given f float32.
func Float32(f float32) *float32 { return &f }

// MustFloat32 returns a float32 value of pointer p, if p is not nil.
// f is returned, if  p is nil.
func MustFloat32(p *float32, f float32) float32 {
	if p != nil {
		return *p
	}
	return f
}

// Z2NFloat32 returns a new pointer of given f float32.
// If f is zero value, it returns nil.
func Z2NFloat32(f float32) *float32 {
	if f == 0 {
		return nil
	}
	return &f
}

// Float64 returns a new pointer of given i float64.
func Float64(f float64) *float64 { return &f }

// MustFloat64 returns a float64 value of pointer p, if p is not nil.
// f is returned, if  p is nil.
func MustFloat64(p *float64, f float64) float64 {
	if p != nil {
		return *p
	}
	return f
}

// Z2NFloat64 returns a new pointer of given f float64.
// If f is zero value, it returns nil.
func Z2NFloat64(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

// Time returns a new pointer of given i time.Time.
func Time(t time.Time) *time.Time { return &t }

// MustTime returns a time.Time value of pointer p, if p is not nil.
// t is returned, if  p is nil.
func MustTime(p *time.Time, t time.Time) time.Time {
	if p != nil {
		return *p
	}
	return t
}

// Z2NTime returns a new pointer of given t time.Time.
// If t is zero value, it returns nil.
func Z2NTime(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}
