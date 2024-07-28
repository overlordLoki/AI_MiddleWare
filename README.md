
# AI Middleware Server

## Overview

This project sets up a middleware server to interact with the LLM API.

## Running the Server

To start the server, run the following script:

```bash
./build.sh
```

This will start the server on port `9090`. You can change the port by modifying the `main.go` file.

## Testing the Server

You can test the server using `curl` or Postman. Here’s an example `curl` command:

```bash
curl -X POST http://localhost:9090/api/chat -H "Content-Type: application/json" -d '{"messages":[{"role":"user","content":"Hello"}]}'
```

This should return a response from the LLM API.