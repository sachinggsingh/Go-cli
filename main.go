package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	models "github.com/sachinggsingh/go-cli/model"
	utils "github.com/sachinggsingh/go-cli/utils"
	progressbar "github.com/schollz/progressbar/v3"
)

func main() {
	path := flag.String("path", "", "Path of the file to read")
	skill := flag.String("skill", "all", "Show only person having this skill")
	show := flag.String("show", "all", "What to show: all | name | address | skills | city")
	filterCity := flag.String("city", "all", "Show only person from this city")
	filterState := flag.String("state", "all", "Show only person from this state")
	name := flag.String("name", "all", "Show only person having this name")
	dependency := flag.String("dependency", "all", "Show only person having this dependency")
	printVersion := flag.Bool("version", false, "If set, also print package.json version")

	flag.Parse()

	if *path == "" {
		flag.PrintDefaults()
		return
	}

	// Show progress bar
	bar := progressbar.Default(10)
	for i := range 10 {
		bar.Add(i)
		time.Sleep(10 * time.Millisecond)
	}

	// Simulate progress
	func() {
		for i := 0; i <= 10; i++ {
			bar.Set(i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Read file while progress bar is showing
	file, err := os.ReadFile(*path)
	if err != nil {
		fmt.Println("\nError reading file:", err)
		return
	}

	// Try decoding as people first
	var people []models.Person
	if err := json.Unmarshal(file, &people); err == nil && len(people) > 0 {
		utils.HandlePeople(people, *skill, *show, *filterCity, *filterState, *name)
		return
	}

	// Try decoding as package.json
	var pkg models.PackageJSON
	if err := json.Unmarshal(file, &pkg); err == nil && len(pkg.Dependencies) > 0 {
		utils.HandlePackage(pkg, *dependency, *printVersion)
		return
	}

	fmt.Println("Unknown JSON format.")
}
