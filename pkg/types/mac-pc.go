package types

import (
	"bridge/pkg/base"
	"fmt"
)

type MacPC struct {
	scanner base.Scanner
}

func (pc *MacPC) Scan() {
	pc.scanner.ScanFile()
	fmt.Println("Scanning Mac PC")
}

func (pc *MacPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
