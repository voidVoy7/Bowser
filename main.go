package main

import (
	"fmt"
	"log"
	"os"

	modules "github.com/Bowser/modules/homebrew"
	"github.com/hairyhenderson/go-which"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bowser <command>")
		fmt.Println("Commands:")
		fmt.Println("  update")
		fmt.Println("  upgrade")
		fmt.Println("  install")
		fmt.Println("  search")
		fmt.Println("  remove")
		os.Exit(1)
	}

	pm, ok := PackageManagerCheck()
	if !ok {
		log.Fatal("no supported package manager found")
	}

	switch pm {

	case "brew":
		switch os.Args[1] {
		case "update":
			modules.BrewUpdateCommand(os.Args[2:])
		case "upgrade":
			modules.BrewUpgradeCommmand(os.Args[2:])
		case "install":
			modules.BrewInstallCommand(os.Args[2:])
		case "search":
			modules.BrewSearchCommand(os.Args[2:])
		case "remove":
			modules.BrewRemoveCommand(os.Args[2:])
		default:
			fmt.Println("Unknown command:", os.Args[1])
			fmt.Println("Commands:")
			fmt.Println("  update")
			fmt.Println("  upgrade")
			fmt.Println("  install")
			fmt.Println("  search")
			fmt.Println("  remove")
			os.Exit(1)
		}

	case "zypper":
		fmt.Println("PLACEHOLDER TEXT")
	}

}

func PackageManagerCheck() (string, bool) {
	supported := []string{"brew", "zypper"}

	for _, candidate := range supported {
		if which.Found(candidate) {
			return candidate, true
		}
	}

	return "", false
}
