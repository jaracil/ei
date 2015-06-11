package ei

import (
	"strings"
)

// M selects one key in map[string]interface{} and returns its value as Ei.
func (i Ei) M(k string) Ei {
	switch v := i.v.(type) {
	case error:
		return i
	case map[string]interface{}:
		d, ok := v[k]
		if !ok {
			return N(NewEiErr("key not found"))
		}
		return N(d)
	case M:
		d, ok := v[k]
		if !ok {
			return N(NewEiErr("key not found"))
		}
		return N(d)
	}
	return N(NewEiErr("type is not map[string]interface{}"))
}

// S selects one index in []interface{} and returns its value as Ei.
func (i Ei) S(idx int) Ei {
	switch v := i.v.(type) {
	case error:
		return i
	case []interface{}:
		if idx < 0 || idx > (len(v)-1) {
			return N(NewEiErr("index out of bounds"))
		}
		return N(v[idx])
	case S:
		if idx < 0 || idx > (len(v)-1) {
			return N(NewEiErr("index out of bounds"))
		}
		return N(v[idx])
	}
	return N(NewEiErr("type is not []interface{}"))
}

// Catch returns val if input value is an error.
func (i Ei) Catch(val interface{}) Ei {
	if _, ok := i.v.(error); ok {
		return N(val)
	}
	return i
}

// Clip corrects input value to min if value < min or to max if value > max.
func (i Ei) Clip(min, max interface{}) Ei {
	v, err := i.Float64()
	if err != nil {
		return N(err)
	}
	p1, err := N(min).Float64()
	if err != nil {
		return N(err)
	}
	p2, err := N(max).Float64()
	if err != nil {
		return N(err)
	}
	if v < p1 {
		return N(p1)
	}
	if v > p2 {
		return N(p2)
	}
	return N(v)
}

// Limit retruns an error if val < min or val > max
func (i Ei) Limit(min, max interface{}) Ei {
	v, err := i.Float64()
	if err != nil {
		return N(err)
	}
	p1, err := N(min).Float64()
	if err != nil {
		return N(err)
	}
	p2, err := N(max).Float64()
	if err != nil {
		return N(err)
	}
	if v < p1 {
		return N(NewEiErr("lower limit overflow"))
	}
	if v > p2 {
		return N(NewEiErr("upper limit overflow"))
	}
	return N(v)
}

// Map replaces the input value by the map[value] entry, returns an error if value is not found.
func (i Ei) Map(m M) Ei {
	k, err := i.String()
	if err != nil {
		return N(err)
	}
	v, ok := m[k]
	if !ok {
		return N(NewEiErr("key not found: " + k))
	}
	return N(v)
}

// In searchs the input value in the list, returns an error if value is not found.
func (i Ei) In(sl []string) Ei {
	k, err := i.String()
	if err != nil {
		return N(err)
	}
	for _, ck := range sl {
		if k == ck {
			return i
		}
	}
	return N(NewEiErr("key not found: " + k))
}

// Upper returns input value converted to uppercase.
func (i Ei) Upper() Ei {
	v, err := i.String()
	if err != nil {
		return N(err)
	}
	return N(strings.ToUpper(v))
}

// Lower returns input value converted to lowercase.
func (i Ei) Lower() Ei {
	v, err := i.String()
	if err != nil {
		return N(err)
	}
	return N(strings.ToLower(v))
}
