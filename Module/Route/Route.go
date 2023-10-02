package Route

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type route struct {
	Name     string   `yaml:"name"`
	Pattern  string   `yaml:"pattern"`
	Function string   `yaml:"function"`
	Methods  []string `yaml:"methods"`
	Redirect string   `yaml:"redirect"`
}

func InitRoute(routeDirectory string) error {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, routeDirectory)

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		yamlPath := filepath.Join(path, file.Name())
		yamlFile, err := os.ReadFile(yamlPath)
		if err != nil {
			return err
		}

		routes, err := getRoute(yamlFile)
		for _, routeI := range *routes {
			routeI.invokeRoute()
		}
	}

	return nil
}

func getRoute(file []byte) (*map[string]route, error) {

	data := make(map[string]route)
	err := yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *route) invokeRoute() {
	pattern := r.Pattern
	function := r.Function
	methods := r.Methods
	redirect := r.Redirect

	if len(methods) == 0 {
		methods = append(methods, "GET")
	}

	fmt.Println(methods)

	invokeHandler(pattern, function, methods, redirect)
}

func haveMethod(methods []string, method string) bool {
	for _, met := range methods {
		if method == met {
			return true
		}
	}
	return false
}
