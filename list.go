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
	"math"

	"github.com/abcum/tlist"
)

// List represents a collection of versions and values, stored
// in order of version number.
type List struct {
	pntr *tlist.List
}

func newList() *List {
	return &List{
		pntr: tlist.New(),
	}
}

// Put inserts a value with the specified version number. It
// returns the previous value, or nil if it does not exist.
func (l *List) Put(ver int64, val []byte) []byte {

	var i *tlist.Item
	var o []byte

	switch ver {
	default:
		i = l.pntr.Get(ver, tlist.Upto)
	case math.MinInt64:
		i = l.pntr.Min()
	case math.MaxInt64:
		i = l.pntr.Max()
	}

	if i != nil {
		o = i.Val()
	}

	l.pntr.Put(ver, val)

	if i != nil {
		return o
	}

	return nil

}

// Get inserts a value with the specified version number, or
// the nearest latest value prior to the specified version.
// If math.MinInt64 is specified for the version, then the
// first item will be returned, and if math.MaxInt64 is used
// then the latest item will be returned.
func (l *List) Get(ver int64) []byte {

	var i *tlist.Item

	switch ver {
	default:
		i = l.pntr.Get(ver, tlist.Upto)
	case math.MinInt64:
		i = l.pntr.Min()
	case math.MaxInt64:
		i = l.pntr.Max()
	}

	if i != nil {
		return i.Val()
	}

	return nil

}

// Del deletes a value with the specified version number, or
// the nearest latest value prior to the specified version.
// If math.MinInt64 is specified for the version, then the
// first item will be deleted, and if math.MaxInt64 is used
// then the latest item will be deleted.
func (l *List) Del(ver int64) []byte {

	var i *tlist.Item

	switch ver {
	default:
		i = l.pntr.Del(ver, tlist.Upto)
	case math.MinInt64:
		i = l.pntr.Min()
	case math.MaxInt64:
		i = l.pntr.Max()
	}

	if i != nil {
		return i.Del().Val()
	}

	return nil

}

// Min returns the value of the minium version in the list.
func (l *List) Min() []byte {

	i := l.pntr.Min()

	if i != nil {
		return i.Val()
	}

	return nil

}

// Max returns the value of the maximum version in the list.
func (l *List) Max() []byte {

	i := l.pntr.Max()

	if i != nil {
		return i.Val()
	}

	return nil

}

// Seek searches for a value prior to the specified version
// number, and returns its version and value. If math.MinInt64
// is specified for the version, then the first item will be
// returned, and if math.MaxInt64 is used then the latest item
// will be returned.
func (l *List) Seek(ver int64) (int64, []byte) {

	var i *tlist.Item

	switch ver {
	default:
		i = l.pntr.Get(ver, tlist.Upto)
	case math.MinInt64:
		i = l.pntr.Min()
	case math.MaxInt64:
		i = l.pntr.Max()
	}

	if i != nil {
		return i.Ver(), i.Val()
	}

	return -1, nil

}

// Walk iterates through all of the versions and values in the
// list, in order of version, starting at the first version.
func (l *List) Walk(fn func(ver int64, val []byte) (exit bool)) {

	l.pntr.Walk(func(i *tlist.Item) bool {
		return fn(i.Ver(), i.Val())
	})

}
