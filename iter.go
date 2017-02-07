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

// Cursor represents an iterator that can traverse over all key-value
// pairs in a tree in sorted order. Cursors can be obtained from a
// transaction and are valid as long as the transaction is open.
// Changing data while traversing with a cursor may cause it to be
// invalidated and return unexpected keys and/or values. You must
// reposition your cursor after mutating data.
type Cursor struct {
	pntr *ptree.Cursor
}

// Del removes the current item under the cursor from the tree. If
// the cursor has not yet been positioned using First, Last, or Seek,
// then no item is deleted and a nil key and value are returned.
func (c *Cursor) Del() ([]byte, *List) {
	if key, val := c.pntr.Del(); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}

// First moves the cursor to the first item in the tree and returns
// its key and value. If the tree is empty then a nil key and value
// are returned.
func (c *Cursor) First() ([]byte, *List) {
	if key, val := c.pntr.First(); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}

// Last moves the cursor to the last item in the tree and returns its
// key and value. If the tree is empty then a nil key and value are
// returned.
func (c *Cursor) Last() ([]byte, *List) {
	if key, val := c.pntr.Last(); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}

// Prev moves the cursor to the previous item in the tree and returns
// its key and value. If the tree is empty then a nil key and value are
// returned, and if the cursor is at the start of the tree then a nil key
// and value are returned. If the cursor has not yet been positioned
// using First, Last, or Seek, then a nil key and value are returned.
func (c *Cursor) Prev() ([]byte, *List) {
	if key, val := c.pntr.Prev(); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}

// Next moves the cursor to the next item in the tree and returns its
// key and value. If the tree is empty then a nil key and value are
// returned, and if the cursor is at the end of the tree then a nil key
// and value are returned. If the cursor has not yet been positioned
// using First, Last, or Seek, then a nil key and value are returned.
func (c *Cursor) Next() ([]byte, *List) {
	if key, val := c.pntr.Next(); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}

// Seek moves the cursor to a given key in the tree and returns it.
// If the specified key does not exist then the next key in the tree
// is used. If no keys follow, then a nil key and value are returned.
func (c *Cursor) Seek(k []byte) ([]byte, *List) {
	if key, val := c.pntr.Seek(k); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}
