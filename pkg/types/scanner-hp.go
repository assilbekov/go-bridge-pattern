package types

import "fmt"

type ScannerHP struct {
}

func (e *ScannerHP) ScanFile() {
	fmt.Println("Scanner Epson is scanning file")
}
