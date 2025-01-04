package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Project struct {
	Name string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: clean-arch-generator <nome-do-projeto>")
		return
	}

	projectName := os.Args[1]
	project := Project{Name: projectName}

	createFolderStructure(projectName)

	generateFile(filepath.Join(projectName, "cmd", "main.go"), "cmd_main.tmpl", project)
	generateFile(filepath.Join(projectName, "internal", "entities", "user.go"), "entities_user.tmpl", project)
	generateFile(filepath.Join(projectName, "internal", "usecases", "user_usecase.go"), "usecases_user_usecase.tmpl", project)
	generateFile(filepath.Join(projectName, "internal", "interfaces", "repositories", "user_repository.go"), "interfaces_repositories_user_repository.tmpl", project)
	generateFile(filepath.Join(projectName, "internal", "repositories", "user_repository_impl.go"), "repositories_user_repository_impl.tmpl", project)
	generateFile(filepath.Join(projectName, "internal", "interfaces", "handlers", "user_handler.go"), "interfaces_handlers_user_handler.tmpl", project)
	generateFile(filepath.Join(projectName, "go.mod"), "go_mod.tmpl", project)
	generateFile(filepath.Join(projectName, "requests.http"), "requests_http.tmpl", project)
	generateFile(filepath.Join(projectName, ".gitignore"), "gitignore.tmpl", project)
	generateFile(filepath.Join(projectName, "README.md"), "readme_md.tmpl", project)

	fmt.Printf("Projeto '%s' criado com sucesso!\n", projectName)
}

func createFolderStructure(projectName string) {
	folders := []string{
		filepath.Join(projectName, "cmd"),
		filepath.Join(projectName, "internal", "entities"),
		filepath.Join(projectName, "internal", "usecases"),
		filepath.Join(projectName, "internal", "interfaces", "repositories"),
		filepath.Join(projectName, "internal", "interfaces", "handlers"),
		filepath.Join(projectName, "internal", "repositories"),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			fmt.Printf("Erro ao criar pasta %s: %v\n", folder, err)
			os.Exit(1)
		}
	}
}

func generateFile(filePath, templateFile string, data Project) {
	tmpl, err := template.New(templateFile).ParseFiles(filepath.Join("templates", templateFile))
	if err != nil {
		fmt.Printf("Erro ao carregar template %s: %v\n", templateFile, err)
		os.Exit(1)
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Printf("Erro ao gerar arquivo %s: %v\n", filePath, err)
		os.Exit(1)
	}
}
