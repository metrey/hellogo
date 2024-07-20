Starting a project in Go (version 1.19) involves setting up your development environment, creating a new project directory, and writing some initial code. Here are the steps to get you started:

### Step 1: Install Go 1.19

1. **Download Go 1.19**:
   - Visit the [official Go downloads page](https://golang.org/dl/).
   - Download the Go 1.19 installer for your operating system (Windows, macOS, or Linux).

2. **Install Go**:
   - Follow the installation instructions for your specific OS.
   - Verify the installation by running:
     ```sh
     go version
     ```
   - You should see something like `go version go1.19 <OS/ARCH>`.

### Step 2: Set Up Your Workspace

1. **Choose a Workspace Directory**:
   - Create a directory for your Go workspace. For example:
     ```sh
     mkdir -p ~/go/src/github.com/yourusername/helloworld
     cd ~/go/src/github.com/yourusername/helloworld
     ```

2. **Initialize a New Module**:
   - Go uses modules to manage dependencies. Initialize a new module in your project directory:
     ```sh
     go mod init github.com/yourusername/helloworld
     ```

### Step 3: Write Your First Go Program

1. **Create a Main Package**:
   - Create a new file called `main.go`:
     ```sh
     touch main.go
     ```

2. **Add Code to `main.go`**:
   - Open `main.go` in your favorite text editor and add the following code:
     ```go
     package main

     import "fmt"

     func main() {
         fmt.Println("Hello, World!")
     }
     ```

### Step 4: Run Your Program

1. **Build and Run**:
   - In the terminal, navigate to your project directory and run:
     ```sh
     go run main.go
     ```
   - You should see `Hello, World!` printed in your terminal.

### Step 5: Managing Dependencies

1. **Adding Dependencies**:
   - To add a dependency to your project, use the `go get` command. For example:
     ```sh
     go get github.com/gin-gonic/gin
     ```
   - This will add the Gin web framework to your project and update your `go.mod` file.

2. **Updating Dependencies**:
   - To update dependencies to their latest versions, run:
     ```sh
     go get -u ./...
     ```

3. **Tidying Up**:
   - Clean up your `go.mod` and `go.sum` files by running:
     ```sh
     go mod tidy
     ```

### Step 6: Building and Testing

1. **Build Your Program**:
   - To build your program into an executable, run:
     ```sh
     go build
     ```
   - This will create an executable file in your project directory.

2. **Run Tests**:
   - To run tests, create a test file (e.g., `main_test.go`) and use the `testing` package:
     ```go
     package main

     import "testing"

     func TestMain(t *testing.T) {
         // Test code here
     }
     ```
   - Run tests with:
     ```sh
     go test ./...
     ```

### Step 7: Organizing Your Project

1. **Project Structure**:
   - As your project grows, organize your code into packages. A common structure is:
     ```
     yourproject/
     ├── go.mod
     ├── go.sum
     ├── main.go
     ├── cmd/
     │   └── yourapp/
     │       └── main.go
     ├── pkg/
     │   └── somepackage/
     │       └── somefile.go
     └── internal/
         └── anotherpackage/
             └── anotherfile.go
     ```

### Summary

Starting a project in Go 1.19 involves setting up your development environment, initializing a module, writing your first Go program, managing dependencies, and organizing your project structure. By following these steps, you'll be well on your way to developing applications in Go.

Common command to use
```
// build and run
go run main.go

// add dependency
go get `<github url of dependency>`

// this command is to auto map module into checksome and check error
go mod tidy

// build project
go build


```