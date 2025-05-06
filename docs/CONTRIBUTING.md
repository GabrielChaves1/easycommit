# Contributing to EasyCommit

Thank you for your interest in contributing to EasyCommit! This document provides guidelines and instructions for contributing to the project.

## Development Setup

### Prerequisites

- Go 1.20 or higher
- Git

### Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
   ```bash
   git clone https://github.com/yourusername/easycommit.git
   cd easycommit
   ```
3. Add the original repository as an upstream remote
   ```bash
   git remote add upstream https://github.com/GabrielChaves1/easycommit.git
   ```
4. Install dependencies
   ```bash
   go mod download
   ```

## Code Style Guidelines

EasyCommit follows standard Go coding conventions:

- Run `go fmt` before committing to ensure consistent formatting
- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `golint` and `go vet` to check for common issues
- Write meaningful comments, especially for exported functions and types
- Keep functions small and focused on a single responsibility
- Use meaningful variable and function names that describe their purpose

## Contribution Workflow

1. Create a new branch for your feature or bugfix

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes, following the code style guidelines

5. Commit your changes with a clear and descriptive commit message

   ```bash
   git commit -m "Add support for new feature X"
   ```

6. Push to your fork

   ```bash
   git push origin feature/your-feature-name
   ```

7. Create a Pull Request against the main repository

## Pull Request Guidelines

- Provide a clear description of the problem you're solving
- Update documentation if necessary
- Add or update tests as appropriate
- Keep PRs focused on a single issue/feature to make them easier to review

## Reporting Issues

When reporting issues, please include:

- A clear description of the problem
- Steps to reproduce
- Expected vs. actual behavior
- Version of EasyCommit you're using
- Go version and OS
- Any relevant logs or error messages

## License

By contributing to EasyCommit, you agree that your contributions will be licensed under the project's MIT license.