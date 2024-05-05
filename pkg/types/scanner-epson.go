package types

import "fmt"

type ScannerEpson struct {
}

func (e *ScannerEpson) ScanFile() {
	fmt.Println("Scanner Epson is scanning file")
}
