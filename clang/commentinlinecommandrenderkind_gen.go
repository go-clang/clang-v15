package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"
import "fmt"

// CommentInlineCommandRenderKind the most appropriate rendering mode for an inline command, chosen on command semantics in Doxygen.
type CommentInlineCommandRenderKind uint32

const (
	// CommentInlineCommandRenderKind_Normal command argument should be rendered in a normal font.
	CommentInlineCommandRenderKind_Normal CommentInlineCommandRenderKind = C.CXCommentInlineCommandRenderKind_Normal
	// CommentInlineCommandRenderKind_Bold command argument should be rendered in a bold font.
	CommentInlineCommandRenderKind_Bold = C.CXCommentInlineCommandRenderKind_Bold
	// CommentInlineCommandRenderKind_Monospaced command argument should be rendered in a monospaced font.
	CommentInlineCommandRenderKind_Monospaced = C.CXCommentInlineCommandRenderKind_Monospaced
	// CommentInlineCommandRenderKind_Emphasized command argument should be rendered emphasized (typically italic font).
	CommentInlineCommandRenderKind_Emphasized = C.CXCommentInlineCommandRenderKind_Emphasized
	// CommentInlineCommandRenderKind_Anchor command argument should not be rendered (since it only defines an anchor).
	CommentInlineCommandRenderKind_Anchor = C.CXCommentInlineCommandRenderKind_Anchor
)

func (cicrk CommentInlineCommandRenderKind) Spelling() string {
	switch cicrk {
	case CommentInlineCommandRenderKind_Normal:
		return "CommentInlineCommandRenderKind=Normal"
	case CommentInlineCommandRenderKind_Bold:
		return "CommentInlineCommandRenderKind=Bold"
	case CommentInlineCommandRenderKind_Monospaced:
		return "CommentInlineCommandRenderKind=Monospaced"
	case CommentInlineCommandRenderKind_Emphasized:
		return "CommentInlineCommandRenderKind=Emphasized"
	case CommentInlineCommandRenderKind_Anchor:
		return "CommentInlineCommandRenderKind=Anchor"
	}

	return fmt.Sprintf("CommentInlineCommandRenderKind unknown %d", int(cicrk))
}

func (cicrk CommentInlineCommandRenderKind) String() string {
	return cicrk.Spelling()
}
