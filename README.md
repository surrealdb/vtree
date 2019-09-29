# vtree

Vtree is an in-memory multi-version radix / patricia tree package for Golang.

[![](https://img.shields.io/circleci/token/1ad2b78f6bd4c9c6acb0a9c7c85f0771767bebaa/project/abcum/vtree/master.svg?style=flat-square)](https://circleci.com/gh/abcum/vtree) [![](https://img.shields.io/badge/status-alpha-ff00bb.svg?style=flat-square)](https://github.com/abcum/vtree) [![](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abcum/vtree) [![](https://goreportcard.com/badge/github.com/abcum/vtree?style=flat-square)](https://goreportcard.com/report/github.com/abcum/vtree) [![](https://img.shields.io/badge/license-Apache_License_2.0-00bfff.svg?style=flat-square)](https://github.com/abcum/vtree) 

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
go get github.com/abcum/vtree
```
