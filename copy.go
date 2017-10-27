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
	"github.com/abcum/ptree"
)

// Copy is a copy of a tree which can be used to apply changes to
// the radix tree. All changes are applied atomically and a new tree
// is returned when committed. A Copy is not thread safe.
type Copy struct {
	pntr *ptree.Copy
}

// Size is used to return the total number of elements in the tree.
func (c *Copy) Size() int {
	return c.pntr.Size()
}

// Root returns the root of the radix tree within this tree copy.
func (c *Copy) Root() *Node {
	return &Node{pntr: c.pntr.Root()}
}

// Tree returns a new tree with the changes committed in memory.
func (c *Copy) Tree() *Tree {
	return &Tree{pntr: c.pntr.Tree()}
}

// Cursor returns a new cursor for iterating through the radix tree.
func (c *Copy) Cursor(ver int64) *Cursor {
	return &Cursor{pntr: c.pntr.Cursor(), ver: ver}
}

// Put is used to insert a specific key, returning the previous value.
func (c *Copy) Put(ver int64, key []byte, val interface{}) interface{} {

	if lst := c.pntr.Get(key); lst != nil {
		return lst.(*list).Put(ver, val)
	}

	lst := newList()
	lst.Put(ver, val)
	c.pntr.Put(key, lst)

	return nil

}

// Get is used to retrieve a specific key, returning the current value.
func (c *Copy) Get(ver int64, key []byte) interface{} {

	switch ver {

	default:

		if lst := c.pntr.Get(key); lst != nil {
			return lst.(*list).Get(ver)
		}

	case 0:

		if lst := c.pntr.Get(key); lst != nil {
			return lst.(*list).Max()
		}

	}

	return nil

}

// Del is used to delete a given key, returning the previous value.
func (c *Copy) Del(ver int64, key []byte) interface{} {

	switch ver {

	default:

		if lst := c.pntr.Get(key); lst != nil {
			return lst.(*list).Del(ver)
		}

	case 0:

		if lst := c.pntr.Del(key); lst != nil {
			return lst.(*list).Max()
		}

	}

	return nil

}
