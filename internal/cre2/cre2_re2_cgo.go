//go:build !re2_wasm && ((linux && (amd64 || arm64)) || (windows && amd64))

package cre2

/*
#cgo CXXFLAGS: -std=c++17
// Note: This file provides C declarations only.
// LDFLAGS must be provided by the importing package to link against RE2.
#include "cre2.h"
*/
import "C"
