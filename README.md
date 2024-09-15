# ethgen - Ethereum Account Generator
ethgen is a command-line tool to generate Ethereum accounts, including the private key, public key, and Ethereum address. It supports output in both plain text and JSON formats, making it a convenient tool for developers, testers, and Ethereum users.

## Features
- Generate Ethereum private key, public key, and address.
- Output in either plain text or JSON formats.
- Simple and user-friendly CLI interface.
- Secure key generation using the go-ethereum library.
  
## Installation
Prerequisites
Make sure you have Go installed on your machine. You can download and install Go from the official Go website.

### Install via go get:
```
go get github.com/devlongs/ethgen
```

### Build from source:
1. Clone the repository:
```
git clone https://github.com/devlongs/ethgen.git
```
2. Change to the project directory:
```
cd ethgen
```
3.Build the project:
```
go build -o ethgen
```
4. Run the binary:
```
./ethgen
```

## Usage
### Generate an Ethereum account
To generate an Ethereum account and display the private key, public key, and address in plain text (default):
```
./ethgen
```

### Generate an Ethereum account with JSON output
To generate an Ethereum account and display the private key, public key, and address in JSON format:
```
./ethgen --format json
```

### Help
To see all available commands and options:
```
./ethgen --help
```

## Example Output
### Plain Text Output
```txt
Private Key: 0x4c0883a69102937d6231471b5dbb6204fe512961708279a49887f8b8693b7767
Public Key: 0x5f3605f5d5b5d6b5dc5f2a92f52a9e29e225bfb9203077b8c7635e62b4adba9c
Ethereum Address: 0x96216849c49358B10257cb55b28eA603c874b05E
```

### JSON Output
```json
{
  "private_key": "0x4c0883a69102937d6231471b5dbb6204fe512961708279a49887f8b8693b7767",
  "public_key": "0x5f3605f5d5b5d6b5dc5f2a92f52a9e29e225bfb9203077b8c7635e62b4adba9c",
  "address": "0x96216849c49358B10257cb55b28eA603c874b05E"
}
```

## Important Notes
- Security: Keep your private key secure and never share it publicly. The private key grants full control over your Ethereum address and its associated assets.
- For development purposes: This tool is intended primarily for quick account generation in development environments. Use strong encryption methods to store private keys securely if used in production.
  

## Contributing
Contributions are welcome! If you'd like to contribute, feel free to open issues or pull requests on the GitHub repository.
