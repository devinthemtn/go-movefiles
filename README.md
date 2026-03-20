# go-movefiles

An interactive Go-based CLI tool to quickly switch between different client configurations in a project.

## 🚀 Overview

`go-movefiles` simplifies the process of managing and swapping configuration files for multiple clients. It provides an interactive menu to select a client and then automatically copies the corresponding configuration files from a client-specific subdirectory into the main configuration directory.

## 🛠️ How it Works

The tool looks for configuration files in `src/Config/[ClientDir]` and copies them to the base `src/Config/` directory.

### Managed Files:
- `_colors.scss`
- `clientConfig.js`
- `clientStyleConfig.scss`
- `configData.js`

### Supported Clients:
- **Dev** (Subdirectory: `Dev`)
- **Salish** (Subdirectory: `Salish`)
- **Blackfoot** (Subdirectory: `BF`)
- **Colville-Tit** (Subdirectory: `Colville-tit`)
- **Colville-nse** (Subdirectory: `Colville-nse`)
- **Colville-nxa** (Subdirectory: `Colville-nxa`)

## 📦 Prerequisites

- [Go](https://go.dev/doc/install) (version 1.24 or higher recommended as per `go.mod`)

## 🏃 Usage

1.  Clone the repository and navigate to the project root.
2.  Run the application:
    ```bash
    go run main.go
    ```
3.  Use the interactive menu to select the client configuration you want to activate.

## 🏗️ Project Structure

```
.
├── main.go            # CLI logic
└── src/
    └── Config/        # Target directory for configuration files
        ├── Dev/       # Source configuration for 'Dev'
        ├── Salish/    # Source configuration for 'Salish'
        └── ...        # Other client-specific folders
```

## 🧩 Key Dependencies

- [huh](https://github.com/charmbracelet/huh): Used for building the interactive terminal selection form.
