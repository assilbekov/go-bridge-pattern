package types

import (
	"bridge/pkg/base"
	"fmt"
)

type LinuxPC struct {
	scanner base.Scanner
}

func (pc *LinuxPC) Scan() {
	pc.scanner.ScanFile()
	fmt.Println("Scanning Linux PC")
}

func (pc *LinuxPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
