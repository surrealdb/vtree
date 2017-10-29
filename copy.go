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
	"bytes"
)

// Copy is a copy of a tree which can be used to apply changes to
// the radix tree. All changes are applied atomically and a new tree
// is returned when committed. A Copy is not thread safe.
type Copy struct {
	size int
	root *Node
}

// Size is used to return the total number of elements in the tree.
func (c *Copy) Size() int {
	return c.size
}

// Root returns the root of the radix tree within this tree copy.
func (c *Copy) Root() *Node {
	return c.root
}

// Tree returns a new tree with the changes committed in memory.
func (c *Copy) Tree() *Tree {
	return &Tree{c.size, c.root}
}

// Cursor returns a new cursor for iterating through the radix tree.
func (c *Copy) Cursor(ver int64) *Cursor {
	return &Cursor{tree: c, ver: ver}
}

// Get is used to retrieve a specific key, returning the current value.
	if lst := c.root.get(key); lst != nil {
		return lst.(*list).Get(ver)
func (c *Copy) Get(ver int64, key []byte) []byte {
	}
	return nil
}

// Del is used to delete a given key, returning the previous value.
func (c *Copy) Del(ver int64, key []byte) []byte {

	if ver > 0 {
		if lst := c.root.get(key); lst != nil {
			return lst.(*list).Del(ver)
		}
		return nil
	}

	root, leaf, old := c.del(nil, c.root, ver, key)
	if root != nil {
		c.root = root
	}
	if leaf != nil {
		c.size--
	}
	return old
}

// Put is used to insert a specific key, returning the previous value.
func (c *Copy) Put(ver int64, key, val []byte) []byte {
	root, leaf, old := c.put(nil, c.root, ver, key, key, val)
	if root != nil {
		c.root = root
	}
	if leaf == nil {
		c.size++
	}
	return old
}

// ---------------------------------------------------------------------------

func prefix(a, b []byte) (i int) {
	for i = 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return
}

func concat(a, b []byte) (c []byte) {
	c = make([]byte, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return
}

func (c *Copy) del(p, n *Node, t int64, s []byte) (*Node, *leaf, []byte) {

	if len(s) == 0 {

		if !n.isLeaf() {
			return nil, nil, nil
		}

		d := n.dup()

		// Remove the leaf node
		d.leaf = nil

		// Check if the node should be merged
		if n != c.root && len(d.edges) == 1 {
			d.mergeChild()
		}

		if t == 0 {

		}

		// Return the found node and leaf node
		return d, n.leaf, n.leaf.val.Get(t)

	}

	// Look for an edge
	l := s[0]
	i, e := n.getSub(l)
	if e == nil || !bytes.HasPrefix(s, e.prefix) {
		return nil, nil, nil
	}

	// Consume the search prefix
	s = s[len(e.prefix):]

	node, leaf, old := c.del(n, e, t, s)
	if node == nil {
		return nil, nil, nil
	}

	// Copy this node
	d := n.dup()

	// Delete the edge if the node has no edges
	if node.leaf == nil && len(node.edges) == 0 {
		d.delSub(l)
		if n != c.root && len(d.edges) == 1 && !d.isLeaf() {
			d.mergeChild()
		}
	} else {
		d.edges[i] = node
	}

	return d, leaf, old

}

func (c *Copy) put(p, n *Node, t int64, s, k, v []byte) (*Node, *leaf, []byte) {

	if len(s) == 0 {

		d := n.dup()

		// Create the leaf if necessary
		if !n.isLeaf() {
			d.leaf = &leaf{key: k, val: newList()}
		}

		// Get the old value
		o := n.leaf.val.Get(t)

		// Update the leaf value
		d.leaf.val.Put(t, v)

		// Return the new node and leaf node
		return d, n.leaf, o

	}

	// Look for the edge
	i, e := n.getSub(s[0])

	// No edge, create one
	if e == nil {
		e := &Node{
			leaf: &leaf{
				key: k,
				val: newList(),
			},
			prefix: s,
		}
		e.leaf.val.Put(t, v)
		d := n.dup()
		d.addSub(e)
		return d, nil, nil
	}

	// Determine longest prefix of the search key on match
	cl := prefix(s, e.prefix)

	if cl == len(e.prefix) {
		s = s[cl:]
		node, leaf, old := c.put(n, e, t, s, k, v)
		if node != nil {
			nc := n.dup()
			nc.edges[i] = node
			return nc, leaf, old
		}
		return nil, leaf, old
	}

	// Split the node
	nc := n.dup()
	splitNode := &Node{
		prefix: s[:cl],
	}
	nc.repSub(splitNode)

	// Restore the existing child node
	modChild := e.dup()
	splitNode.addSub(modChild)
	modChild.prefix = modChild.prefix[cl:]

	// Create a new leaf node
	leaf := &leaf{
		key: k,
		val: newList(),
	}
	leaf.val.Put(t, v)

	// If the new key is a subset, add to to this node
	s = s[cl:]
	if len(s) == 0 {
		splitNode.leaf = leaf
		return nc, nil, nil
	}

	// Create a new edge for the node
	splitNode.addSub(&Node{
		leaf:   leaf,
		prefix: s,
	})

	return nc, nil, nil

}
