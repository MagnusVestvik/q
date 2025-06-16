# Q

A command-line tool for displaying directory contents in a beautiful box format.

## Features

- Displays directory contents in a formatted box
- Shows file/directory name, type, and size
- Human-readable file sizes
- Customizable display configuration

## Installation

```bash
go install github.com/svoosh/q@latest
```

## Usage

```bash
q [directory]
```

If no directory is specified, the current directory will be used.


## Installation
```bash
go install github.com/MagnusVestvik/q/cmd/q@latest
```

## Project Structure

This project follows the standard Go project layout:

```
.
├── cmd/            # Main applications
│   └── q/         # The main application
├── pkg/           # Library code that's ok to use by external applications
│   ├── domain/    # Domain-specific code
│   └── logic/     # Business logic
├── internal/      # Private application and library code
├── docs/          # Documentation
├── scripts/       # Build and deployment scripts
├── build/         # Packaging and CI
│   ├── package/   # Packaging scripts
│   └── ci/        # CI configuration
└── test/          # Additional test files
```

## Development

1. Clone the repository:
```bash
git clone https://github.com/svoosh/q.git
cd q
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build -o bin/q cmd/q/main.go
```
