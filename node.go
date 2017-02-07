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

// Node represents an immutable node in the radix tree which
// can be either an edge node or a leaf node.
type Node struct {
	pntr *ptree.Node
}

// Min returns the key and value of the minimum item in the
// subtree of the current node.
func (n *Node) Min() ([]byte, *List) {
	if key, val := n.pntr.Min(); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}

// Max returns the key and value of the maximum item in the
// subtree of the current node.
func (n *Node) Max() ([]byte, *List) {
	if key, val := n.pntr.Max(); key != nil {
		return key, val.(*List)
	}
	return nil, nil
}

// Path is used to recurse over the tree only visiting nodes
// which are above this node in the tree.
func (n *Node) Path(p []byte, f Walker) {
	n.pntr.Path(p, func(key []byte, val interface{}) (exit bool) {
		return f(key, val.(*List))
	})
}

// Subs is used to recurse over the tree only visiting nodes
// which are directly under this node in the tree.
func (n *Node) Subs(p []byte, f Walker) {
	n.pntr.Subs(p, func(key []byte, val interface{}) (exit bool) {
		return f(key, val.(*List))
	})
}

// Walk is used to recurse over the tree only visiting nodes
// which are under this node in the tree.
func (n *Node) Walk(p []byte, f Walker) {
	n.pntr.Walk(p, func(key []byte, val interface{}) (exit bool) {
		return f(key, val.(*List))
	})
}
