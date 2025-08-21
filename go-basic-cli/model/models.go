package model

type Person struct {
	Name    string   `json:"name"`
	Skills  []string `json:"skills"`
	Address struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
}

type PackageJSON struct {
	Name         string            `json:"name"`
	Private      bool              `json:"private"`
	Version      string            `json:"version"`
	Type         string            `json:"type"`
	Scripts      map[string]string `json:"scripts"`
	Dependencies map[string]string `json:"dependencies"`
	DevDeps      map[string]string `json:"devDependencies"`
}
