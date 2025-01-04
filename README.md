 # Generate project with Clean Architecture in Go

 This is a code generator project designed to create Go projects following Clean Architecture.

 ## How to Use

 1. **Build the generator**:
    ```bash
    go build -o clean-arch-generator
    ```

 2. **Generate a new project**:
    ```bash
    ./clean-arch-generator <project-name>
    ```

 3. **Navigate to the generated project**:
    ```bash
    cd <project-name>
    ```

 4. **Install dependencies**:
    ```bash
    go mod tidy
    ```

 5. **Run the project**:
    ```bash
    go run cmd/main.go
    ```

 ## Features

 - **Clean Architecture**: Generates a project structure following Clean Architecture principles.
 - **Ready-to-use templates**: Includes templates for entities, use cases, repositories, and handlers.
 - **Pre-configured `.gitignore`**: Ignores unnecessary files like binaries and IDE configurations.
 - **Example requests**: Provides a `requests.http` file for testing API endpoints.
 - **Pre-commit hook**: Automatically sets up a pre-commit hook to ensure code quality.

 ## Pre-Commit Hook

 The generated project includes a pre-commit hook to ensure code quality before committing changes. It performs the following checks:

 - **Code formatting**: Ensures the code is properly formatted using `gofmt`.
 - **Linting**: Runs `golangci-lint` to catch common issues.
 - **Tests**: Executes all tests in the project using `go test`.

 To use the pre-commit hook, make sure `golangci-lint` is installed:
 ```bash
 go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
 ```

 ## Project Structure

 The generated project includes the following folders and files:

 - **`cmd/`**: Contains the `main.go` file to run the application.
 - **`internal/`**: Contains the core application logic, divided into:
   - **`entities/`**: Business entities and models.
   - **`usecases/`**: Business logic and use cases.
   - **`interfaces/`**: Handlers and repositories interfaces.
   - **`repositories/`**: Implementation of data persistence.
 - **`requests.http`**: Example HTTP requests for testing the API.
 - **`.gitignore`**: Pre-configured to ignore unnecessary files.
 - **`README.md`**: Project documentation (this file).
 - **`go.mod`**: Go module configuration.

 ## How to Contribute

 1. Fork the repository.
 2. Create a new branch for your feature (`git checkout -b feature/new-feature`).
 3. Commit your changes (`git commit -m 'Add new feature'`).
 4. Push to the branch (`git push origin feature/new-feature`).
 5. Open a Pull Request.

 ## License

 This project is licensed under the [MIT License](LICENSE).