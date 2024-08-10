package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sys/windows"
)

var totalLoadTime time.Duration

func init() {
	start := time.Now()

	// List of DLLs to load
	dlls := []string{
		"C:\\Program Files\\Java\\jdk-22\\bin\\api-ms-win-core-file-l1-1-0.dll",
		"C:\\Users\\Sandeep Garg\\Desktop\\Garima\\golang\\telnets-master\\project\\src\\LoadDLL\\Qt6.dll",
		"D:\\QT\\bin\\QtWidgets.dll",
		"D:\\QT\\6.7.0\\mingw_64\\bin\\Qt6Core.dll",
	}

	for _, dllPath := range dlls {
		// Load the DLL
		dll, err := windows.LoadDLL(dllPath)
		if err != nil {
			log.Fatalf("Error loading DLL %s: %v\n", dllPath, err)
		}
		defer dll.Release()
	}

	totalLoadTime = time.Since(start)
}

func main() {
	fmt.Printf("Total time taken to load all DLLs: %v\n", totalLoadTime)
}
