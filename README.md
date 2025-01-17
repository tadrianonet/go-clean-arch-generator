# Generate Go Project with Clean Architecture


This is a Go code generator designed to create projects following the principles of Clean Architecture. The generator provides a well-structured project setup with predefined templates for entities, use cases, repositories, and handlers, making it easy to get started with Clean Architecture in Go.

With a simple command, you can generate a new Go project that adheres to best practices and includes essential components like OpenTelemetry configuration, pre-commit hooks for code quality checks, and example requests for testing API endpoints.

Ideal for developers who want to quickly start a project with Clean Architecture without worrying about boilerplate setup.

 ## Execubles

The last release:

* [Windows](./executables/clean-arch-generator.exe) 
* [Mac](./executables/clean-arch-generator) 
 

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
 - **OpenTelemetry**: Cofigured with Opentelemetry

 ## Pre-Commit Hook

 The generated project includes a pre-commit hook to ensure code quality before committing changes. It performs the following checks:

 - **Code formatting**: Ensures the code is properly formatted using `gofmt`.
 - **Linting**: Runs `golangci-lint` to catch common issues.
 - **Tests**: Executes all tests in the project using `go test`.

 To use the pre-commit hook, make sure `golangci-lint` is installed:
 ```bash
 go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
 ```

  ## OpenTelemetry

  To test Opentelemetry with `zipkin` localhost you can follow this steps:

* Create a container:

 ```bash
docker run -d --name zipkin -p 9411:9411 openzipkin/zipkin 
 ```

 * Install packages

 ```bash
go get  go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
go get 	go.opentelemetry.io/otel/sdk/resource
go get  go.opentelemetry.io/otel/semconv/v1.4.0
go get  go.opentelemetry.io/otel/exporters/zipkin

```

 * Update webwebapp with this code:

```bash

	func Start() {
	container := dependencies.Setup()

	// Cria um router para usar com o middleware otelhttp
	mux := http.NewServeMux()

	err := container.Invoke(func(userHandler *handlers.UserHandler) {
		mux.HandleFunc("/users", userHandler.CreateUser)
		mux.HandleFunc("/users/get", userHandler.GetUserByID)
	})

	if err != nil {
		log.Fatalf("Erro ao resolver dependências: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := infra.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	handler := otelhttp.NewHandler(mux, "http-server")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
```

* Update otel.go -> newTraceProvider with this code:

```bash

    func newTraceProvider(ctx context.Context) (*trace.TracerProvider, error) {
	zipkinEndpoint := "http://localhost:9411/api/v2/spans" // Substitua se usar um endpoint diferente

	zipkinExporter, err := zipkin.New(zipkinEndpoint)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("bal"), // Nome do serviço
		),
		resource.WithFromEnv(), // Adiciona atributos de ambiente, se disponíveis
	)
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(zipkinExporter,
			trace.WithBatchTimeout(time.Second)), // Ajuste o tempo limite conforme necessário
		trace.WithResource(res),
	)

	return traceProvider, nil
}

```
* Update SetupOTelSDK with this code:

```bash
tracerProvider, err := newTraceProvider(ctx)
```



 ## Project Structure

 The generated project includes the following folders and files:

 - **`cmd/`**: Contains the `main.go` file to run the application.
 - **`internal/`**: Contains the core application logic, divided into:
   - **`entities/`**: Business entities and models.
   - **`infra/`**: Infraestructure like otel or database.
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
