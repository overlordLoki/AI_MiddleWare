
# AI Middleware

## Overview

AI Middleware is a Go application designed to interact with an AI service. It provides an HTTP API that forwards chat messages to an AI service and retrieves responses.


## Setup

### Prerequisites

- Go 1.16 or later
- An AI service running at `http://localhost:11434/api/chat`

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/AI_MiddleWare.git
   cd AI_MiddleWare
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the application:

   ```bash
   go build -o ai_middleware main.go
   ```

## Running the Application

Start the server:

```bash
./ai_middleware
```

The server will run on `localhost:8085`.

## API Endpoints

### POST /chat

Send a chat request to the AI service.

**Request Body:**

```json
{
  "messages": [
    {
      "role": "user",
      "content": "your message here"
    }
  ]
}
```

**Response:**

```json
{
  "response": "AI response here"
}
```

### GET /

A simple endpoint to check if the server is running.

**Response:**

```json
{
  "message": "hello world"
}
```

## Testing

Run the unit tests with:

```bash
go test ./...
```

To test the `/chat` endpoint, use the `test.http` file with an HTTP client like Postman or `curl`.

## Deployment

To deploy this application to a home server:

1. Build the application as described above.
2. Transfer the binary to your home server.
3. Run the application on the server.

