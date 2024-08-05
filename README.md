# Port Scanner

A simple command-line tool to scan ports on a specified host, written in Go.

## Description

This port scanner checks whether the specified TCP ports on a given host are open or closed. It uses the `net` package in Go to attempt TCP connections and reports the status of each port.

## Features

- **Host Scanning**: Specify the host to scan.
- **Port Scanning**: Provide a comma-separated list of ports to check.
- **Timeout Handling**: Set a timeout for each connection attempt.

## Installation

To build the project, you'll need Go installed on your system. Clone this repository and use the provided Makefile to build the binary.

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yashasolutions/port_scanner.git
   cd port_scanner
   ```

2. **Build the project**
```sh
make build
```

3. **Basic Usage**

```sh
./port_scanner -host=<HOST> -ports=<PORTS>
```

- <HOST>: The host to scan (e.g., example.com or 192.168.1.1).
- <PORTS>: A comma-separated list of ports to scan (e.g., 80,443,8080).


Example : 

```sh
./port_scanner -host=example.com -ports=22,80,443
```
