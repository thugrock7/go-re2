//go:build !re2_wasm && ((linux && (amd64 || arm64)) || (windows && amd64))

// Package cre2 provides public access to cre2 C API wrappers.
// This is a thin wrapper that re-exports internal/cre2 for external consumption.
package cre2

import (
	"unsafe"

	"github.com/wasilibs/go-re2/internal/cre2"
)

// Re-export all functions from internal/cre2

func New(patternPtr unsafe.Pointer, patternLen int, opts unsafe.Pointer) unsafe.Pointer {
	return cre2.New(patternPtr, patternLen, opts)
}

func Delete(ptr unsafe.Pointer) {
	cre2.Delete(ptr)
}

func ErrorCode(rePtr unsafe.Pointer) int {
	return cre2.ErrorCode(rePtr)
}

func ErrorArg(rePtr unsafe.Pointer) unsafe.Pointer {
	return cre2.ErrorArg(rePtr)
}

func FindAndConsume(rePtr, textPtr, matchPtr unsafe.Pointer, nMatch int) bool {
	return cre2.FindAndConsume(rePtr, textPtr, matchPtr, nMatch)
}

func GlobalReplace(rePtr, textAndTargetPtr, rewritePtr unsafe.Pointer) bool {
	return cre2.GlobalReplace(rePtr, textAndTargetPtr, rewritePtr)
}

func Match(rePtr, textPtr unsafe.Pointer, textLen, startPos, endPos, anchor int, matchArr unsafe.Pointer, nMatch int) bool {
	return cre2.Match(rePtr, textPtr, textLen, startPos, endPos, anchor, matchArr, nMatch)
}

func NamedGroupsIterNew(rePtr unsafe.Pointer) unsafe.Pointer {
	return cre2.NamedGroupsIterNew(rePtr)
}

func NamedGroupsIterNext(iterPtr unsafe.Pointer, namePtr *unsafe.Pointer, indexPtr *int) bool {
	return cre2.NamedGroupsIterNext(iterPtr, namePtr, indexPtr)
}

func NamedGroupsIterDelete(iterPtr unsafe.Pointer) {
	cre2.NamedGroupsIterDelete(iterPtr)
}

func NumCapturingGroups(rePtr unsafe.Pointer) int {
	return cre2.NumCapturingGroups(rePtr)
}

func NewOpt() unsafe.Pointer {
	return cre2.NewOpt()
}

func DeleteOpt(opt unsafe.Pointer) {
	cre2.DeleteOpt(opt)
}

func OptSetLogErrors(opt unsafe.Pointer, flag bool) {
	cre2.OptSetLogErrors(opt, flag)
}

func OptSetLongestMatch(opt unsafe.Pointer, flag bool) {
	cre2.OptSetLongestMatch(opt, flag)
}

func OptSetPosixSyntax(opt unsafe.Pointer, flag bool) {
	cre2.OptSetPosixSyntax(opt, flag)
}

func OptSetCaseSensitive(opt unsafe.Pointer, flag bool) {
	cre2.OptSetCaseSensitive(opt, flag)
}

func OptSetLatin1Encoding(opt unsafe.Pointer) {
	cre2.OptSetLatin1Encoding(opt)
}

func OptSetMaxMem(opt unsafe.Pointer, size int) {
	cre2.OptSetMaxMem(opt, size)
}

func NewSet(opt unsafe.Pointer, anchor int) unsafe.Pointer {
	return cre2.NewSet(opt, anchor)
}

func SetAdd(set, patternPtr unsafe.Pointer, patternLen int) unsafe.Pointer {
	return cre2.SetAdd(set, patternPtr, patternLen)
}

func SetCompile(set unsafe.Pointer) int {
	return cre2.SetCompile(set)
}

func SetMatch(set, textPtr unsafe.Pointer, textLen int, match unsafe.Pointer, nMatch int) int {
	return cre2.SetMatch(set, textPtr, textLen, match, nMatch)
}

func SetDelete(ptr unsafe.Pointer) {
	cre2.SetDelete(ptr)
}

func Malloc(size int) unsafe.Pointer {
	return cre2.Malloc(size)
}

func Free(ptr unsafe.Pointer) {
	cre2.Free(ptr)
}

func CopyCBytes(sPtr unsafe.Pointer, sLen int) []byte {
	return cre2.CopyCBytes(sPtr, sLen)
}

func CopyCString(sPtr unsafe.Pointer) string {
	return cre2.CopyCString(sPtr)
}

func CopyCStringN(sPtr unsafe.Pointer, n int) string {
	return cre2.CopyCStringN(sPtr, n)
}
