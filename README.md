# Pincracker

A concurrent Go-based utility to crack 4-digit PIN endpoints. Initally designed to speed up HTTP requests for the HackTheBox Brute Force Attacks Module.

## Features
- **Native Concurrency:** Built using Goroutines for faster execution times.
- **Flexible Inputs:** Automatically detects both `IP:PORT` and `IP PORT` positional argument formats.

---

## Quickstart

### Prerequisites
Make sure you have [Go installed](https://go.dev/doc/install) on your system.

### 1. Installation & Compilation
Clone this repository (or download the `main.go` file) and compile it locally:

```bash
# Initialize the module
go mod init pincracker

# Build the custom binary executable
go build -o pincracker main.go
```

### 2. Make it Global (Optional)
To run the command from any directory without typing `./`, add it to your path. Some examples are:

**Linux / macOS:**
```bash
sudo mv pincracker /usr/local/bin/
```

**Windows:**
Move pincracker.exe to a dedicated tools folder (e.g., C:\Tools) and add that folder to your PATH.

### 3. Usage
The script automatically parses your target input format:
```Bash 
Option A: IP and Port together
pincracker 10.10.11.234:8080

# Option B: IP and Port separated
pincracker 10.10.11.234 8080
```

## Benchmark: Go vs Python
Execution time comparison when brute-forcing a full 4-digit PIN (0000-9999) against the same lab environment:

### Python:
Note that IP and Port are hardcoded in the script that Hack The Box grants us.

```bash
➜ python pincracker.py
<SNIP>
Correct PIN found: XXXX
Flag: XXXXXXXXXXXXXXXXXXXXX
Bruteforcing attack took: 641.35 seconds to complete
```

### Go:
```bash
➜ pincracker 154.57.164.74 31796
[*] Scanning target -> IP: 154.57.164.74 | Port: 31796

[+] Pin was found: XXXX
[+] Flag: XXXXXXXXXXXXXXXXXXXXX
[*] Total time: 16.759814634s
```
