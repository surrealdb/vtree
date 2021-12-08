# vtree

Vtree is an in-memory multi-version radix / patricia tree package for Golang.

[![](https://img.shields.io/badge/status-1.0.0-ff00bb.svg?style=flat-square)](https://github.com/surrealdb/vtree) [![](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/surrealdb/vtree) [![](https://goreportcard.com/badge/github.com/surrealdb/vtree?style=flat-square)](https://goreportcard.com/report/github.com/surrealdb/vtree) [![](https://img.shields.io/badge/license-Apache_License_2.0-00bfff.svg?style=flat-square)](https://github.com/surrealdb/vtree) 

#### Features

- Immutable radix tree
- Copy-on-write radix tree
- Rich transaction support
- Versioned key-value items
- Select key-value items since a specific version
- Insert, and delete key-value items with a specific version
- Iterate through all versions of every key-value item

#### Installation

```bash
go get github.com/surrealdb/vtree
```
