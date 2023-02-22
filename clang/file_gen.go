package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "time"

// File a particular source file that is part of a translation unit.
type File struct {
	c C.CXFile
}

// GetFileName retrieve the complete file and path name of the given file.
func (f File) Name() string {
	o := cxstring{C.clang_getFileName(f.c)}
	defer o.Dispose()

	return o.String()
}

// GetFileTime retrieve the last modification time of the given file.
func (f File) Time() time.Time {
	return time.Unix(int64(C.clang_getFileTime(f.c)), 0)
}

// GetFileUniqueID retrieve the unique ID for the given file.
//
// Parameter file the file to get the ID for.
// Parameter outID stores the returned CXFileUniqueID.
// Returns If there was a failure getting the unique ID, returns non-zero,
// otherwise returns 0.
func (f File) UniqueID() (FileUniqueID, int32) {
	var outID FileUniqueID

	o := int32(C.clang_getFileUniqueID(f.c, &outID.c))

	return outID, o
}

// IsEqual returns non-zero if the file1 and file2 point to the same file, or they are both NULL.
func (f File) IsEqual(file2 File) bool {
	o := C.clang_File_isEqual(f.c, file2.c)

	return o != C.int(0)
}

// TryGetRealPathName returns the real path name of file.
//
// An empty string may be returned. Use clang_getFileName() in that case.
func (f File) TryGetRealPathName() string {
	o := cxstring{C.clang_File_tryGetRealPathName(f.c)}
	defer o.Dispose()

	return o.String()
}
