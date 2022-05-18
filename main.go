package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
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

func PrintModule(mod *debug.Module, indent string) {
	fmt.Printf("%s%s\n", indent, mod.Path)
	fmt.Printf("%sVersion : %s\n", indent, mod.Version)
	fmt.Printf("%sSum     : %s\n", indent, mod.Sum)
	if mod.Replace != nil {
		fmt.Printf("%sReplace:\n", indent)
		PrintModule(mod.Replace, indent+"  ")
	}
}
