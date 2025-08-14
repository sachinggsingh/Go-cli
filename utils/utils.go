package utils

import (
	"fmt"
	"strings"

	model "github.com/sachinggsingh/go-cli/model"
)

func HandlePeople(people []model.Person, skill, show, filterCity, filterState, name string) {
	found := false
	for _, p := range people {
		// city
		if !strings.EqualFold(filterCity, "all") && !strings.EqualFold(p.Address.City, filterCity) {
			continue
		}
		// state
		if !strings.EqualFold(filterState, "all") && !strings.EqualFold(p.Address.State, filterState) {
			continue
		}
		// skill
		if !strings.EqualFold(skill, "all") {
			matched := false
			for _, s := range p.Skills {
				if strings.EqualFold(s, skill) {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}
		// name
		if !strings.EqualFold(name, "all") && !strings.Contains(strings.ToLower(p.Name), strings.ToLower(name)) {
			continue
		}
		// If we reach here, p matches all filters
		found = true
		printPerson(p, show)
	}
	if !found {
		fmt.Println("No data")
	}
}

func printPerson(p model.Person, show string) {
	switch strings.ToLower(show) {
	case "name":
		fmt.Println(p.Name)
	case "address":
		fmt.Println(p.Address)
	case "skills":
		fmt.Printf("%s: %v\n", p.Name, p.Skills)
	case "city":
		fmt.Printf("%s: %s\n", p.Name, p.Address.City)
	default:
		fmt.Printf("Name: %s\nCity: %s\nSkills: %v\n\n", p.Name, p.Address.City, p.Skills)
	}
}

func HandlePackage(pkg model.PackageJSON, dependency string, printVersion bool) {
	if strings.EqualFold(dependency, "all") {
		if printVersion {
			fmt.Println("Package:", pkg.Name)
			fmt.Println("Version:", pkg.Version)
		}
		fmt.Println("Dependencies:")
		for dep, ver := range pkg.Dependencies {
			fmt.Printf("  %s: %s\n", dep, ver)
		}
		fmt.Println("\nDev Dependencies:")
		for dep, ver := range pkg.DevDeps {
			fmt.Printf("  %s: %s\n", dep, ver)
		}
		if len(pkg.Scripts) > 0 {
			fmt.Println("\nScripts:")
			for s, cmd := range pkg.Scripts {
				fmt.Printf("  %s: %s\n", s, cmd)
			}
		}
		return
	}

	for dep, ver := range pkg.Dependencies {
		if strings.EqualFold(dep, dependency) {
			if printVersion {
				fmt.Printf("Package: %s\nVersion: %s\n", pkg.Name, pkg.Version)
			}
			fmt.Printf("Dependency: %s %s\n", dep, ver)
			return
		}
	}
	for dep, ver := range pkg.DevDeps {
		if strings.EqualFold(dep, dependency) {
			if printVersion {
				fmt.Printf("Package: %s\nVersion: %s\n", pkg.Name, pkg.Version)
			}
			fmt.Printf("Dev Dependency: %s %s\n", dep, ver)
			return
		}
	}
	fmt.Println("No data")
}
