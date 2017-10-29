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

// List represents a collection of versions and values, stored
// in order of version number.
type list struct {
	pntr *tlist.List
}

func newList() *list {
	return &list{
		pntr: tlist.New(),
	}
}

// Put inserts a value with the specified version number. It
// returns the previous value, or nil if it does not exist.
func (i *Item) Put(ver int64, val []byte) []byte {

	var i *tlist.Item
	var o interface{}

	switch ver {
	default:
		i = l.pntr.Get(ver, tlist.Upto)
	case 0:
		i = l.pntr.Max()
	}

	if i != nil {
		o = i.Val()
	}

	if i != nil && i.Ver() == ver {
		i.Set(val)
	} else {
		l.pntr.Put(ver, val)
	}

	return o

}

// Get inserts a value with the specified version number, or
// the nearest latest value prior to the specified version.
// If math.MinInt64 is specified for the version, then the
// first item will be returned, and if math.MaxInt64 is used
// then the latest item will be returned.
func (i *Item) Get(ver int64) []byte {

	var i *tlist.Item

	switch ver {
	default:
		i = l.pntr.Get(ver, tlist.Upto)
	case 0:
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
func (i *Item) Del(ver int64) []byte {

	if i := l.pntr.Del(ver, tlist.Upto); i != nil {
		return i.Del().Val()
	}

	return nil

}

// Min returns the value of the minium version in the list.
func (i *Item) Min() []byte {

	if i := l.pntr.Min(); i != nil {
		return i.Val()
	}

	return nil

}

// Max returns the value of the maximum version in the list.
func (i *Item) Max() []byte {

	if i := l.pntr.Max(); i != nil {
		return i.Val()
	}

	return nil

}
