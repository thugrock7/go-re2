//go:build !re2_wasm

package cre2

/*
#cgo CXXFLAGS: -std=c++17
// Only use pkg-config if RE2_USE_PKGCONFIG build tag is set
// By default, assumes cre2 symbols available via dlopen
#cgo !re2_use_pkgconfig LDFLAGS:
#cgo re2_use_pkgconfig pkg-config: re2
*/
import "C"
