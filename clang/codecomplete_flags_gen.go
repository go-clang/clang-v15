package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// CodeComplete_Flags flags that can be passed to clang_codeCompleteAt() to
// modify its behavior.
//
// The enumerators in this enumeration can be bitwise-OR'd together to
// provide multiple options to clang_codeCompleteAt().
type CodeComplete_Flags uint32

const (
	// CodeComplete_IncludeMacros whether to include macros within the set of code completions returned.
	CodeComplete_IncludeMacros CodeComplete_Flags = C.CXCodeComplete_IncludeMacros
	// CodeComplete_IncludeCodePatterns whether to include code patterns for language constructs within the set of code completions, e.g., for loops.
	CodeComplete_IncludeCodePatterns = C.CXCodeComplete_IncludeCodePatterns
	// CodeComplete_IncludeBriefComments whether to include brief documentation within the set of code completions returned.
	CodeComplete_IncludeBriefComments = C.CXCodeComplete_IncludeBriefComments
	// CodeComplete_SkipPreamble whether to speed up completion by omitting top- or namespace-level entities defined in the preamble. There's no guarantee any particular entity is omitted. This may be useful if the headers are indexed externally.
	CodeComplete_SkipPreamble = C.CXCodeComplete_SkipPreamble
	// CodeComplete_IncludeCompletionsWithFixIts whether to include completions with small fix-its, e.g. change '.' to '->' on member access, etc.
	CodeComplete_IncludeCompletionsWithFixIts = C.CXCodeComplete_IncludeCompletionsWithFixIts
)

func (ccf CodeComplete_Flags) Spelling() string {
	switch ccf {
	case CodeComplete_IncludeMacros:
		return "CodeComplete=IncludeMacros"
	case CodeComplete_IncludeCodePatterns:
		return "CodeComplete=IncludeCodePatterns"
	case CodeComplete_IncludeBriefComments:
		return "CodeComplete=IncludeBriefComments"
	case CodeComplete_SkipPreamble:
		return "CodeComplete=SkipPreamble"
	case CodeComplete_IncludeCompletionsWithFixIts:
		return "CodeComplete=IncludeCompletionsWithFixIts"
	}

	return fmt.Sprintf("CodeComplete_Flags unknown %d", int(ccf))
}

func (ccf CodeComplete_Flags) String() string {
	return ccf.Spelling()
}
