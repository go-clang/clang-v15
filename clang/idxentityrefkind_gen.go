package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// IdxEntityRefKind data for IndexerCallbacks#indexEntityReference.
//
// This may be deprecated in a future version as this duplicates
// the CXSymbolRole_Implicit bit in CXSymbolRole.
type IdxEntityRefKind uint32

const (
	// IdxEntityRef_Direct the entity is referenced directly in user's code.
	IdxEntityRef_Direct IdxEntityRefKind = C.CXIdxEntityRef_Direct
	// IdxEntityRef_Implicit an implicit reference, e.g. a reference of an Objective-C method via the dot syntax.
	IdxEntityRef_Implicit = C.CXIdxEntityRef_Implicit
)

func (ierk IdxEntityRefKind) Spelling() string {
	switch ierk {
	case IdxEntityRef_Direct:
		return "IdxEntityRef=Direct"
	case IdxEntityRef_Implicit:
		return "IdxEntityRef=Implicit"
	}

	return fmt.Sprintf("IdxEntityRefKind unknown %d", int(ierk))
}

func (ierk IdxEntityRefKind) String() string {
	return ierk.Spelling()
}
