# Distributed Log Query

## Overview
A beginner-friendly, powerful command-line tool written in Go for querying logs across distributed systems efficiently. This project enables searching and analyzing log files on multiple machines from a single interface, using TCP socket programming for communication.

## Features
- Query logs across multiple servers simultaneously
- Filter results using grep-like syntax
- Real-time streaming of results
- Concurrent processing of queries using goroutines
- Containerized with Docker for easy deployment
- Minimal network and system overhead
- Support for custom log file locations

## Supported Libraries
- `net` - For TCP socket programming
- `shlex` - For parsing command-line syntax
- `sync` - For managing concurrency
- `io` - For handling data streams
- `exec` - For executing commands on remote servers
- `fmt` - For formatting output
- `strings` - For string manipulation

## Installation
```bash
# Clone the repository
git clone https://github.com/yourusername/distributed-log-query.git
cd distributed-log-query

# Build Client with Docker
docker build -t myapp-client -f Dockerfile.client .
docker run --rm -it myapp-client

# On a separate terminal, build Servers with Docker
docker build -t myapp-image -f Dockerfile.server .
docker compose up --build

```

## Usage
```bash

```

## Architecture
The system consists of:
- **TCP Client**: Connects to servers, sends queries, and displays results
- **TCP Server**: Listens for incoming connections, processes queries, and returns results
- **Query Processor**: Parses and executes grep commands on log files
- **Concurrency Manager**: Handles multiple connections using goroutines and sync primitives

## Configuration
Create a `docker-compose.yaml` file:
```yaml
version: '3'
services:
  server1:
    image: myapp-image
    ports:
      - "8081:8080"
    volumes:
      - ./log_files:/app/log_files
  server2:
    image: myapp-image
    ports:
      - "8082:8080"
```

## Contributing
Contributions welcome! Please check out our [contribution guidelines](CONTRIBUTING.md).

## License
MIT
# Distributed-Log-Querier
