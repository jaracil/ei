package ei

import (
	"strconv"
	"time"
)

// convenience shorthand for []interface{} type.
type S []interface{}

// convenience shorthand for map[string]interface{} type.
type M map[string]interface{}

// EiErr specific Ei conversion error
type EiErr struct {
	err string
}

// NewEiErr creates new EiErr.
func NewEiErr(err string) error {
	return &EiErr{err: err}
}

// IsEiErr returns true if err type is EiErr.
func IsEiErr(err error) bool {
	_, ok := err.(*EiErr)
	return ok
}

// Error returns EiErr error string.
func (e *EiErr) Error() string {
	return e.err
}

// Ei (Empty Interface struct). Used to wrap interface{} and provide methods.
type Ei struct {
	v interface{}
}

// N wraps a suppored type value (including interface{}) into Ei type.
//
// Supported types are:
//   int8, int16, int32, int64, int
//   uin8, byte, uint16, uint32, uint64, uint
//   float32, float64
//   string, []byte
//   time.Time
//   map[string]interface{} or shorthand ei.M
//   []interface{} or shorthand ei.S
//   interface{} pointing to one of above types
func N(i interface{}) Ei {
	return Ei{v: i}
}

// F inserts a transformation function to Ei evaluation chain.
func (i Ei) F(f func(Ei, ...interface{}) Ei, p ...interface{}) Ei {
	return f(i, p...)
}

// RawZ returns the content of Ei as interface{} or nil if Ei points
// to EiErr
func (i Ei) RawZ() interface{} {
	r, _ := i.Raw()
	return r
}

// Raw returns the content of Ei as interface{} or error if Ei points
// to EiErr
func (i Ei) Raw() (interface{}, error) {
	switch v := i.v.(type) {
	case error:
		return nil, v
	default:
		return i.v, nil
	}
}

// Int64Z converts Ei to int64, returns zero value on error.
func (i Ei) Int64Z() int64 {
	r, _ := i.Int64()
	return r
}

// Int64 converts Ei to int64.
func (i Ei) Int64() (int64, error) {
	switch v := i.v.(type) {
	case error:
		return 0, v
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if n, err := strconv.ParseInt(v, 0, 64); err == nil {
			return int64(n), nil
		}
	case time.Time:
		return int64(v.Unix()), nil
	}
	return 0, NewEiErr("type conversion error")
}

// Uint64Z converts Ei to uint64, returns zero value on error.
func (i Ei) Uint64Z() uint64 {
	r, _ := i.Uint64()
	return r
}

// Uint64 converts Ei to uint64.
func (i Ei) Uint64() (uint64, error) {
	switch v := i.v.(type) {
	case error:
		return 0, v
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case int:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case float32:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if n, err := strconv.ParseUint(v, 0, 64); err == nil {
			return uint64(n), nil
		}
	case time.Time:
		return uint64(v.Unix()), nil
	}
	return 0, NewEiErr("type conversion error")
}

// Float64Z converts Ei to float64, returns zero value on error
func (i Ei) Float64Z() float64 {
	r, _ := i.Float64()
	return r
}

// Float64 converts Ei to float64.
func (i Ei) Float64() (float64, error) {
	switch v := i.v.(type) {
	case error:
		return 0, v
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return float64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if n, err := strconv.ParseFloat(v, 64); err == nil {
			return float64(n), nil
		}
	case time.Time:
		return float64(v.Unix()), nil

	}
	return 0, NewEiErr("type conversion error")
}

// StringZ converts Ei to string, returns empty string on error.
func (i Ei) StringZ() string {
	r, _ := i.String()
	return r
}

// String converts Ei to string.
func (i Ei) String() (string, error) {
	switch v := i.v.(type) {
	case error:
		return "", v
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(int64(v), 10), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64), nil
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(v), nil
	case string:
		return v, nil
	case time.Time:
		return v.Format(time.RFC3339Nano), nil
	}
	return "", NewEiErr("type conversion error")
}

// TimeZ converts Ei to time.Time, returns time.Time{} value on error.
func (i Ei) TimeZ() time.Time {
	r, _ := i.Time()
	return r
}

// Time converts Ei to time.Time.
func (i Ei) Time() (time.Time, error) {
	switch v := i.v.(type) {
	case error:
		return time.Time{}, v
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case int:
		return time.Unix(int64(v), 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(int64(v), 0), nil
	case float32:
		return time.Unix(int64(v), 0), nil
	case float64:
		return time.Unix(int64(v), 0), nil
	case string:
		formats := []string{
			time.ANSIC,
			time.RFC3339,
			time.RFC822,
			time.RFC822Z,
			time.RFC850,
			time.RFC1123,
			time.RFC1123Z,
			time.UnixDate,
			"2006-01-02 15:04:05", // UTC
			"2006-01-02",          // UTC
		}

		for _, f := range formats {
			t, _ := time.Parse(f, v)
			if !t.IsZero() {
				return t, nil
			}
		}
	case time.Time:
		return v, nil
	}
	return time.Time{}, NewEiErr("type conversion error")
}

// BytesZ converts Ei to []byte, returns nil on error.
func (i Ei) BytesZ() []byte {
	r, _ := i.Bytes()
	return r
}

// Bytes converts Ei to []byte.
func (i Ei) Bytes() ([]byte, error) {
	switch v := i.v.(type) {
	case error:
		return nil, v
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	}
	return nil, NewEiErr("type conversion error")
}

// Float32Z converts Ei to float32, returns zero value on error.
func (i Ei) Float32Z() float32 {
	r, _ := i.Float32()
	return r
}

// Float32 converts Ei to float32.
func (i Ei) Float32() (float32, error) {
	r, e := i.Float64()
	return float32(r), e
}

// BoolZ converts Ei to bool, returns false on error.
func (i Ei) BoolZ() bool {
	r, _ := i.Bool()
	return r
}

// Bool converts Ei to bool.
func (i Ei) Bool() (bool, error) {
	r, e := i.Uint64()
	return r != 0, e
}

// IntZ converts Ei to int, returns zero value on error.
func (i Ei) IntZ() int {
	r, _ := i.Int()
	return r
}

// Int converts Ei to int.
func (i Ei) Int() (int, error) {
	r, e := i.Int64()
	return int(r), e
}

// Int32Z converts Ei to int32, returns zero value on error.
func (i Ei) Int32Z() int32 {
	r, _ := i.Int32()
	return r
}

// Int32 converts Ei to int32.
func (i Ei) Int32() (int32, error) {
	r, e := i.Int64()
	return int32(r), e
}

// Int16Z converts Ei to int16, returns zero value on error.
func (i Ei) Int16Z() int16 {
	r, _ := i.Int16()
	return r
}

// Int16 converts Ei to int16.
func (i Ei) Int16() (int16, error) {
	r, e := i.Int64()
	return int16(r), e
}

// Int8Z converts Ei to int8, returns zero value on error.
func (i Ei) Int8Z() int8 {
	r, _ := i.Int8()
	return r
}

// Int8 converts Ei to int8.
func (i Ei) Int8() (int8, error) {
	r, e := i.Int64()
	return int8(r), e
}

// UintZ converts Ei to uint, returns zero value on error.
func (i Ei) UintZ() uint {
	r, _ := i.Uint()
	return r
}

// Uint converts Ei to uint.
func (i Ei) Uint() (uint, error) {
	r, e := i.Uint64()
	return uint(r), e
}

// Uint32Z converts Ei to uint32, returns zero value on error.
func (i Ei) Uint32Z() uint32 {
	r, _ := i.Uint32()
	return r
}

// Uint32 converts Ei to uint32.
func (i Ei) Uint32() (uint32, error) {
	r, e := i.Uint64()
	return uint32(r), e
}

// Uint16Z converts Ei to uint16, returns zero value on error.
func (i Ei) Uint16Z() uint16 {
	r, _ := i.Uint16()
	return r
}

// Uint16 converts Ei to uint16.
func (i Ei) Uint16() (uint16, error) {
	r, e := i.Uint64()
	return uint16(r), e
}

// Uint8Z converts Ei to uint8, returns zero value on error.
func (i Ei) Uint8Z() uint8 {
	r, _ := i.Uint8()
	return r
}

// Uint8 converts Ei to uint8.
func (i Ei) Uint8() (uint8, error) {
	r, e := i.Uint64()
	return uint8(r), e
}

// ByteZ converts Ei to byte, returns zero value on error.
func (i Ei) ByteZ() byte {
	r, _ := i.Uint8()
	return r
}

// Byte converts Ei to byte.
func (i Ei) Byte() (byte, error) {
	return i.Uint8()
}

// Len returns the length of map[string]interface{} or []interface{}
// stored in Ei.
func (i Ei) Len() (int, error) {
	switch v := i.v.(type) {
	case error:
		return 0, v
	case map[string]interface{}:
		return len(v), nil
	case M:
		return len(v), nil
	case []interface{}:
		return len(v), nil
	case S:
		return len(v), nil
	default:
		return 0, NewEiErr("type don't support Len()")
	}
}

// HasKey returns true if map[string]interface{} stored in Ei has the key k.
func (i Ei) HasKey(k string) (bool, error) {
	switch v := i.v.(type) {
	case error:
		return false, v
	case map[string]interface{}:
		_, ok := v[k]
		return ok, nil
	case M:
		_, ok := v[k]
		return ok, nil
	default:
		return false, NewEiErr("type don't support HasKey()")
	}
}
