package main

import "bridge/pkg/types"

var (
	hpScanner    = &types.ScannerHP{}
	epsonScanner = &types.ScannerEpson{}
	linuxPC      = &types.LinuxPC{}
	macPC        = &types.MacPC{}
	windowsPC    = &types.WindowsPC{}
)

func main() {
	linuxPC.AddScanner(hpScanner)
	linuxPC.Scan()

	macPC.AddScanner(epsonScanner)
	macPC.Scan()

	windowsPC.AddScanner(hpScanner)
	windowsPC.Scan()
}
