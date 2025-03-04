# 🚀 Solutions to LeetCode Problems in Go

This repository contains solutions to **LeetCode** problems implemented in **Go**. Each problem is fully independent, with its own module, main entry point, and unit tests. The goal is to improve skills in algorithms, data structures, and best practices for writing efficient and maintainable code.

## 📌 Repository Structure

Each LeetCode problem is organized into its own folder as a separate Go module with the following structure:

```
📂 problem_name/
 ├── go.mod            # Independent Go module
 ├── main.go           # Entry point for execution
 ├── solution.go       # Go implementation
 ├── solution_test.go  # Unit tests in Go
 ├── readme.md         # Problem description and solution explanation
```

### 📂 Example:
```
📂 88_merged_sorted_array/
 ├── go.mod
 ├── main.go
 ├── solution.go
 ├── solution_test.go
 ├── readme.md
```

## 🛠️ Technologies and Tools

- Language: **Go (Golang)**
- Each problem is an independent **Go module** (`go.mod` per problem)
- Auto-formatting with `gofmt`
- Linter: `golangci-lint`
- Unit testing with `go test`
- `pre-commit` hooks to ensure code quality before commits

## 🚀 How to Use

1. Clone the repository:
   ```sh
   git clone https://github.com/MartinFedericoAcevedo/leetcode-solutions.git
   cd leetcode-solutions
   ```

2. Navigate to a specific problem folder:
   ```sh
   cd merged_sorted_array_88
   ```

3. Initialize dependencies (if needed):
   ```sh
   go mod tidy
   ```

4. Run the solution:
   ```sh
   go run .
   ```

5. Run unit tests:
   ```sh
   go test
   ```

## 📌 Project Goals

- Solve algorithmic problems in a fully modular and independent way.
- Apply design patterns and best practices in Go.
- Improve performance and optimization of solutions.
- Maintain unit tests for each solution.

## 🤝 Contributions

If you want to contribute with new solutions or improvements, feel free to submit a **pull request**! 🎯

## 📜 License

This repository is distributed under the **MIT License**. You are free to use it for learning and personal development.
