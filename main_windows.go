package main

/*
#include <stdio.h>
#include <stdlib.h>

// C function to open a file and return a FILE*
static FILE* openFile(const char* filename) {
    return fopen(filename, "r");
}
*/
import "C"
import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var eqgfx syscall.Handle
var t3dParseReadWorldFromStream uintptr

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
	fmt.Println("Success")
}

func run() error {
	var err error
	// check if eqgfx_dx8.dll exists
	_, err = os.Stat("eqgfx_dx8.dll")
	if err != nil {
		return fmt.Errorf("eqgfx_dx8.dll not found: %v", err)
	}

	eqgfx, err = syscall.LoadLibrary("eqgfx_dx8.dll")
	if err != nil {
		return fmt.Errorf("loadlibrary failed: %v", err)
	}

	t3dParseReadWorldFromStream, err = syscall.GetProcAddress(eqgfx, "t3dParseReadWorldFromStream")
	if err != nil {
		return fmt.Errorf("getprocaddress failed: %v", err)
	}

	// Open a file using the C function
	filename := C.CString("path/to/your/file.txt")
	defer C.free(unsafe.Pointer(filename))
	file := C.openFile(filename)
	if file == nil {
		fmt.Println("Failed to open file")
		return
	}
	defer C.fclose(file)

	// Read world from stream
	err := readWorldFromStream(file)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success")
	}

	return nil
}

func readWorldFromStream(file *C.FILE) error {
	// Convert the FILE* to a uintptr
	filePtr := uintptr(unsafe.Pointer(file))

	// Call the syscall with the file pointer as the first argument
	r1, r2, errCode := syscall.SyscallN(t3dParseReadWorldFromStream, filePtr, 0, 0, 0)
	if errCode != 0 {
		return fmt.Errorf("syscall failed: %v", errCode)
	}
	if r1 != 0 {
		return fmt.Errorf("syscall failed: %v", r1)
	}
	if r2 != 0 {
		return fmt.Errorf("syscall failed: %v", r2)
	}

	return nil
}
