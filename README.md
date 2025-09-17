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
