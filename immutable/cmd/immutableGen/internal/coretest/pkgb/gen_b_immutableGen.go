// Code generated by immutableGen. DO NOT EDIT.

package pkgb

//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

//
// PkgB is an immutable type and has the following template:
//
// 	struct {
// 		Postcode string
// 	}
//
type PkgB struct {
	field_Postcode string

	mutable bool
	__tmpl  *_Imm_PkgB
}

var _ immutable.Immutable = new(PkgB)
var _ = new(PkgB).__tmpl

func (s *PkgB) AsMutable() *PkgB {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *PkgB) AsImmutable(v *PkgB) *PkgB {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *PkgB) Mutable() bool {
	return s.mutable
}

func (s *PkgB) WithMutable(f func(si *PkgB)) *PkgB {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *PkgB) WithImmutable(f func(si *PkgB)) *PkgB {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}

func (s *PkgB) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	return true
}
func (s *PkgB) Postcode() string {
	return s.field_Postcode
}

// SetPostcode is the setter for Postcode()
func (s *PkgB) SetPostcode(n string) *PkgB {
	if s.mutable {
		s.field_Postcode = n
		return s
	}

	res := *s
	res.field_Postcode = n
	return &res
}