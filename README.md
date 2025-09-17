# gin-api

A simple REST API built with [Gin](https://github.com/gin-gonic/gin).

The API provides two endpoints:

- **`/ask`** – accepts a question and returns a test answer from an AI stub.  
- **`/health`** – simple health check returning `{ "status": "ok" }`.  

---

## Run locally

### Start the server
```bash
go run cmd/server/main.go
```

By default, the server listens on port `:8080`.

---

## Usage

### Test `/ask`
```bash
curl -X POST http://localhost:8080/ask \
  -H "Content-Type: application/json" \
  -d '{"question": "Test question"}'
```

Example response:
```json
{
  "answer": "This is a test answer from the AI for the question: Test question",
  "source": "stubbed"
}
```

### Test `/health`
```bash
curl http://localhost:8080/health
```

Response:
```json
{ "status": "ok" }
```

---

## Run with Docker

### Build the image
```bash
docker build -t gin-api:local .
```

### Run the container
```bash
docker run --rm -p 8080:8080 \
  -e ENV=prod \
  -e SERVER_PORT=":8080" \
  gin-api:local
```

- `ENV=prod` enables stricter CORS configuration.  
- Leaving `ENV` empty runs the API in development mode (all origins allowed).  
