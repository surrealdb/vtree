// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vtree

import (
	"github.com/abcum/tlist"
)

// Item represents a collection of versions and values, stored
// in order of version number.
type Item struct {
	pntr *tlist.List
}

func newItem() *Item {
	return &Item{
		pntr: tlist.New(),
	}
}

// Put inserts a value with the specified version number. It
// returns the previous value, or nil if it does not exist.
func (i *Item) Put(ver int64, val []byte) []byte {

	var v *tlist.Item
	var o []byte

	switch ver {
	default:
		v = i.pntr.Get(ver, tlist.Upto)
	case 0:
		v = i.pntr.Max()
	}

	if v != nil {
		o = v.Val()
	}

	if v != nil && v.Ver() == ver {
		v.Set(val)
	} else {
		i.pntr.Put(ver, val)
	}

	return o

}

// Get selects a value with the specified version number, or
// the nearest latest value prior to the specified version.
// If '0' is specified for the version, then the latest item
// will be returned.
func (i *Item) Get(ver int64) []byte {

	var v *tlist.Item

	switch ver {
	default:
		v = i.pntr.Get(ver, tlist.Upto)
	case 0:
		v = i.pntr.Max()
	}

	if v != nil {
		return v.Val()
	}

	return nil

}

// Del deletes a value with the specified version number, or
// the nearest latest value prior to the specified version.
func (i *Item) Del(ver int64) []byte {

	if v := i.pntr.Del(ver, tlist.Upto); v != nil {
		return v.Del().Val()
	}

	return nil

}

// Min returns the value of the minium version in the list.
func (i *Item) Min() []byte {

	if v := i.pntr.Min(); v != nil {
		return v.Val()
	}

	return nil

}

// Max returns the value of the maximum version in the list.
func (i *Item) Max() []byte {

	if v := i.pntr.Max(); v != nil {
		return v.Val()
	}

	return nil

}

// Seek searches for a value prior to the specified version
// number, and returns its version and value. If math.MinInt64
// is specified for the version, then the first item will be
// returned, and if math.MaxInt64 is used then the latest item
// will be returned.
func (i *Item) Seek(ver int64) (int64, []byte) {

	var v *tlist.Item

	switch ver {
	default:
		v = i.pntr.Get(ver, tlist.Upto)
	case 0:
		v = i.pntr.Max()
	}

	if v != nil {
		return v.Ver(), v.Val()
	}

	return -1, nil

}

// Walk iterates through all of the versions and values in the
// list, in order of version, starting at the first version.
func (i *Item) Walk(fn func(ver int64, val []byte) (exit bool)) {

	i.pntr.Walk(func(i *tlist.Item) bool {
		return fn(i.Ver(), i.Val())
	})

}
