package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
)

func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "show version")
	flag.Parse()

	if version {
		fmt.Printf("buildinfo-test: %s\n", Version())
		return
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("failed to read build info")
		return
	}

	fmt.Printf("GoVersion : %s\n", info.GoVersion)
	fmt.Printf("Path      : %s\n", info.Path)
	fmt.Printf("Main      :\n")
	PrintModule(&info.Main, "  ")
	fmt.Printf("Deps      : \n")
	for i := range info.Deps {
		PrintModule(info.Deps[i], "  ")
	}
	fmt.Printf("Settings  : \n")
	for _, s := range info.Settings {
		fmt.Printf("%-20s: %s\n", s.Key, s.Value)
	}
}

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unkown"
	}

	return fmt.Sprintf("%s %s/%s", info.Main.Version, runtime.GOOS, runtime.GOARCH)
}

func PrintModule(mod *debug.Module, indent string) {
	fmt.Printf("%s%s\n", indent, mod.Path)
	fmt.Printf("%sVersion : %s\n", indent, mod.Version)
	fmt.Printf("%sSum     : %s\n", indent, mod.Sum)
	if mod.Replace != nil {
		fmt.Printf("%sReplace:\n", indent)
		PrintModule(mod.Replace, indent+"  ")
	}
}
