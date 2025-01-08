package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type Project struct {
	Name string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: clean-arch-generator <project-name>")
		return
	}

	projectName := os.Args[1]
	project := Project{Name: projectName}

	// Create folder structure
	createFolderStructure(projectName)

	// Initialize Git repository
	if err := initGitRepository(projectName); err != nil {
		fmt.Printf("Error initializing Git repository: %v\n", err)
		os.Exit(1)
	}

	// Generate base files

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

	generateFile(filepath.Join(projectName, "internal", "delivery", "web.app.go"), "web.app.tmpl", project)
	generateFile(filepath.Join(projectName, "internal", "delivery", "dependencies", "dependencies.go"), "dependencies.tmpl", project)

	// Generate pre-commit hook
	preCommitPath := filepath.Join(projectName, ".git", "hooks", "pre-commit")
	generateFile(preCommitPath, "pre-commit.tmpl", project)

	// Make the pre-commit hook executable
	os.Chmod(preCommitPath, 0755)

	fmt.Printf("Project '%s' created successfully!\n", projectName)
}

func createFolderStructure(projectName string) {
	folders := []string{
		filepath.Join(projectName, "cmd"),
		filepath.Join(projectName, "internal", "entities"),
		filepath.Join(projectName, "internal", "usecases"),
		filepath.Join(projectName, "internal", "interfaces", "repositories"),
		filepath.Join(projectName, "internal", "interfaces", "handlers"),
		filepath.Join(projectName, "internal", "repositories"),
		filepath.Join(projectName, "internal", "delivery", "dependencies"),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			fmt.Printf("Erro ao criar pasta %s: %v\n", folder, err)
			os.Exit(1)
		}
	}
}

func initGitRepository(projectName string) error {
	cmd := exec.Command("git", "init", projectName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize Git repository: %v", err)
	}
	return nil
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
