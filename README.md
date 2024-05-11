# Inventory API Server

A REST API for interacting with an inventory tracking DB

# Getting Started

1. `docker build -t inventory_api_server .`
2. `docker compose up -d`

# Testing

`go test ./test`

# TODO

- Check API input
    - Prepared statements
    - Sanitization
    - Escape chars
- Authentication
    - API keys / Bearer