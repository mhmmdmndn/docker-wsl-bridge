# Docker WSL Wrapper

A Go-based utility that provides seamless integration between Windows and Docker commands through Windows Subsystem for Linux (WSL). This tool allows Windows users to execute Docker and Docker Compose commands natively while the actual Docker engine runs in WSL.

## Features

- Transparent command forwarding from Windows to WSL
- Support for both Docker and Docker Compose commands
- Configurable WSL distribution through environment variables
- Proper handling of input/output streams
- Error handling and appropriate exit codes

## Prerequisites

- Windows 10 or later
- Windows Subsystem for Linux (WSL) installed
- Docker installed in your WSL distribution
- Go 1.16 or later (for building from source)

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/mhmmdmndn/docker-wsl-wrapper.git
   cd docker-wsl-wrapper
   ```

2. Build the executables:
   ```bash
   go build -o docker.exe docker-wrapper.go
   go build -o docker-compose.exe docker-composer-wrapper.go
   ```

3. Add the directory containing the executables to your system's PATH

## Usage

Once installed, you can use the Docker commands directly from Windows Command Prompt or PowerShell:

```bash
# Docker commands
docker ps
docker images
docker run hello-world

# Docker Compose commands
docker-compose up
docker-compose down
```

### Configuration

By default, the wrapper uses the "Ubuntu" WSL distribution. You can change this by setting the `WSL_DISTRO` environment variable. Here's how to set it permanently:

#### Method 1: Using Windows GUI

1. Press `Win + X` and select "System"
2. Click on "Advanced system settings"
3. Click on "Environment Variables"
4. Under "User variables" or "System variables":
   - Click "New"
   - Variable name: `WSL_DISTRO`
   - Variable value: Your WSL distribution name (e.g., `Debian`)
5. Click "OK" on all windows

#### Method 2: Using Command Line

1. To set for current user (PowerShell, administrator):
```powershell
[System.Environment]::SetEnvironmentVariable('WSL_DISTRO', 'Debian', 'User')
```

2. To set system-wide (PowerShell, administrator):
```powershell
[System.Environment]::SetEnvironmentVariable('WSL_DISTRO', 'Debian', 'Machine')
```

3. Temporary setting (current session only):
```powershell
set WSL_DISTRO=Debian
```

Note: After setting the environment variable, you'll need to restart your terminal or applications to apply the changes.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to Microsoft for WSL
- Thanks to Docker for their amazing containerization technology