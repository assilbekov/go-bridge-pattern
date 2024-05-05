package types

import (
	"bridge/pkg/base"
	"fmt"
)

type WindowsPC struct {
	scanner base.Scanner
}

func (pc *WindowsPC) Scan() {
	pc.scanner.ScanFile()
	fmt.Println("Scanning Windows PC")
}

func (pc *WindowsPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
