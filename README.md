BasicList: A sorted Linked List implement in Go
==================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/basiclist/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/basiclist?status.svg)](https://godoc.org/github.com/kkdai/basiclist)  [![Build Status](https://travis-ci.org/kkdai/basiclist.svg?branch=master)](https://travis-ci.org/kkdai/basiclist)


What is Basic List
---------------

It is not a new data structure, it is a sorted linked list implement in Golang. I implement this for study linked list part for [Skip List](https://en.wikipedia.org/wiki/Skip_list).


Install
---------------
`go get github.com/kkdai/basiclist`


Usage
---------------

    // New a BasicList
    bList := NewBasicList()
    //Insert a search key 3, value is string3 (value could be any `interface{}`)
    bList.Insert(3, "string3")
    bList.Insert(4, "string4")
    bList.Insert(2, "string2")
    //Display all linked list.
    bList.DisplayAll()
    //head->[key:0][val:<nil>]->[key:2][val:string2]->[key:3][val:string3]->[key:4][val:string4]->nil
    //Remove from list
    bList.Remove(3)    
    bList.DisplayAll()
    //head->[key:0][val:<nil>]->[key:2][val:string2]->[key:4][val:string4]->nil

### Inspired By:

- [大数据日知录：架构与算法](http://product.dangdang.com/23561651.html)


License
---------------

This package is licensed under MIT license. See LICENSE for details.


