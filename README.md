# NoFun

------------------

[![tag](https://img.shields.io/github/v/tag/kirilldd2/go-no-fun?style=flat-square)](https://github.com/kirilldd2/go-no-fun)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-lightgrey?style=flat-square)
[![GoDoc](https://pkg.go.dev/badge/github.com/kirilldd2/go-no-fun.svg)](https://pkg.go.dev/github.com/kirilldd2/go-no-fun)
![tests](https://img.shields.io/github/workflow/status/kirilldd2/go-no-fun/Tests?label=tests&logo=github&style=flat-square)
[![Go report](https://goreportcard.com/badge/github.com/kirilldd2/go-no-fun?style=flat-square)](https://goreportcard.com/report/github.com/kirilldd2/go-no-fun)
[![codecov](https://img.shields.io/codecov/c/github/kirilldd2/go-no-fun?style=flat-square&token=jNRLNzybbM)](https://codecov.io/gh/kirilldd2/go-no-fun)
[![License](https://img.shields.io/github/license/kirilldd2/go-no-fun?style=flat-square)](./LICENSE)


😎 `kirilldd2/go-no-fun`/NoFun (Not Only FUNctional) provides utility functions inspired by functional-style programming, 
python standard library and "[samber/lo](https://github.com/samber/lo)" Go package. Based on generics. Supports Go 1.18+.

*Designed with simplicity in mind. In this project I will try to implement simple and useful API. 
Signatures like __DoX2 -> DoX9__ is a big no-no*.

**‼ Disclaimer:** In some cases this kind of approach + generics gives big overhead, 
compared to the same thing written via for loop. So think twice before using this where performance is important. 

------------------
## Checklist

* [x] [Map](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Map)
* [x] [Reduce](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Reduce)
* [x] [Filter](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Filter)
* [x] [Min](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Min)/[Max](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Max)
* [x] [Reverse](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Reverse)/[Reversed](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Reversed)
* [x] [Equal](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Equal)
* [x] [Index](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Index)/[IndexAB](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#IndexAB)
* [x] [Any](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#Any)/[All](https://pkg.go.dev/github.com/kirilldd2/go-no-fun#All)
* [ ] ExtendMap/ExtendSlice/ExtendedMap/ExtendedSlice
* [ ] Range
* [ ] Sum
* [ ] Set + methods
* [ ] Fast versions
