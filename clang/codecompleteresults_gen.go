package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// CodeCompleteResults contains the results of code-completion.
//
// This data structure contains the results of code completion, as
// produced by clang_codeCompleteAt(). Its contents must be freed by
// clang_disposeCodeCompleteResults.
type CodeCompleteResults struct {
	c *C.CXCodeCompleteResults
}

// GetCompletionNumFixIts retrieve the number of fix-its for the given completion index.
//
// Calling this makes sense only if CXCodeComplete_IncludeCompletionsWithFixIts
// option was set.
//
// # Parameter results The structure keeping all completion results
//
// Parameter completion_index The index of the completion
//
// Return The number of fix-its which must be applied before the completion at
// completion_index can be applied
func (ccr *CodeCompleteResults) CompletionNumFixIts(completionIndex uint32) uint32 {
	return uint32(C.clang_getCompletionNumFixIts(ccr.c, C.uint(completionIndex)))
}

// GetCompletionFixIt fix-its that *must* be applied before inserting the text for the
// corresponding completion.
//
// By default, clang_codeCompleteAt() only returns completions with empty
// fix-its. Extra completions with non-empty fix-its should be explicitly
// requested by setting CXCodeComplete_IncludeCompletionsWithFixIts.
//
// For the clients to be able to compute position of the cursor after applying
// fix-its, the following conditions are guaranteed to hold for
// replacement_range of the stored fix-its:
// - Ranges in the fix-its are guaranteed to never contain the completion
// point (or identifier under completion point, if any) inside them, except
// at the start or at the end of the range.
// - If a fix-it range starts or ends with completion point (or starts or
// ends after the identifier under completion point), it will contain at
// least one character. It allows to unambiguously recompute completion
// point after applying the fix-it.
//
// The intuition is that provided fix-its change code around the identifier we
// complete, but are not allowed to touch the identifier itself or the
// completion point. One example of completions with corrections are the ones
// replacing '.' with '->' and vice versa:
//
// std::unique_ptr<std::vector<int>> vec_ptr;
// In 'vec_ptr.^', one of the completions is 'push_back', it requires
// replacing '.' with '->'.
// In 'vec_ptr->^', one of the completions is 'release', it requires
// replacing '->' with '.'.
//
// # Parameter results The structure keeping all completion results
//
// Parameter completion_index The index of the completion
//
// Parameter fixit_index The index of the fix-it for the completion at
// completion_index
//
// Parameter replacement_range The fix-it range that must be replaced before the
// completion at completion_index can be applied
//
// Returns The fix-it string that must replace the code at replacement_range
// before the completion at completion_index can be applied
func (ccr *CodeCompleteResults) CompletionFixIt(completionIndex uint32, fixitIndex uint32) (SourceRange, string) {
	var replacementRange SourceRange

	o := cxstring{C.clang_getCompletionFixIt(ccr.c, C.uint(completionIndex), C.uint(fixitIndex), &replacementRange.c)}
	defer o.Dispose()

	return replacementRange, o.String()
}

// DisposeCodeCompleteResults free the given set of code-completion results.
func (ccr *CodeCompleteResults) Dispose() {
	C.clang_disposeCodeCompleteResults(ccr.c)
}

// CodeCompleteGetNumDiagnostics determine the number of diagnostics produced prior to the location where code completion was performed.
func (ccr *CodeCompleteResults) NumDiagnostics() uint32 {
	return uint32(C.clang_codeCompleteGetNumDiagnostics(ccr.c))
}

// CodeCompleteGetDiagnostic retrieve a diagnostic associated with the given code completion.
//
// Parameter Results the code completion results to query.
// Parameter Index the zero-based diagnostic number to retrieve.
//
// Returns the requested diagnostic. This diagnostic must be freed
// via a call to clang_disposeDiagnostic().
func (ccr *CodeCompleteResults) Diagnostic(index uint32) Diagnostic {
	return Diagnostic{C.clang_codeCompleteGetDiagnostic(ccr.c, C.uint(index))}
}

// CodeCompleteGetContexts determines what completions are appropriate for the context
// the given code completion.
//
// # Parameter Results the code completion results to query
//
// Returns the kinds of completions that are appropriate for use
// along with the given code completion results.
func (ccr *CodeCompleteResults) Contexts() uint64 {
	return uint64(C.clang_codeCompleteGetContexts(ccr.c))
}

// CodeCompleteGetContainerKind returns the cursor kind for the container for the current code
// completion context. The container is only guaranteed to be set for
// contexts where a container exists (i.e. member accesses or Objective-C
// message sends); if there is not a container, this function will return
// CXCursor_InvalidCode.
//
// # Parameter Results the code completion results to query
//
// Parameter IsIncomplete on return, this value will be false if Clang has complete
// information about the container. If Clang does not have complete
// information, this value will be true.
//
// Returns the container kind, or CXCursor_InvalidCode if there is not a
// container
func (ccr *CodeCompleteResults) ContainerKind() (uint32, CursorKind) {
	var isIncomplete C.uint

	o := CursorKind(C.clang_codeCompleteGetContainerKind(ccr.c, &isIncomplete))

	return uint32(isIncomplete), o
}

// CodeCompleteGetContainerUSR returns the USR for the container for the current code completion
// context. If there is not a container for the current context, this
// function will return the empty string.
//
// # Parameter Results the code completion results to query
//
// Returns the USR for the container
func (ccr *CodeCompleteResults) ContainerUSR() string {
	o := cxstring{C.clang_codeCompleteGetContainerUSR(ccr.c)}
	defer o.Dispose()

	return o.String()
}

// CodeCompleteGetObjCSelector returns the currently-entered selector for an Objective-C message
// send, formatted like "initWithFoo:bar:". Only guaranteed to return a
// non-empty string for CXCompletionContext_ObjCInstanceMessage and
// CXCompletionContext_ObjCClassMessage.
//
// # Parameter Results the code completion results to query
//
// Returns the selector (or partial selector) that has been entered thus far
// for an Objective-C message send.
func (ccr *CodeCompleteResults) Selector() string {
	o := cxstring{C.clang_codeCompleteGetObjCSelector(ccr.c)}
	defer o.Dispose()

	return o.String()
}

// Results the code-completion results.
func (ccr CodeCompleteResults) Results() []CompletionResult {
	var s []CompletionResult
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(ccr.c.NumResults)
	gos_s.Len = int(ccr.c.NumResults)
	gos_s.Data = uintptr(unsafe.Pointer(ccr.c.Results))

	return s
}

// NumResults the number of code-completion results stored in the Results array.
func (ccr CodeCompleteResults) NumResults() uint32 {
	return uint32(ccr.c.NumResults)
}
