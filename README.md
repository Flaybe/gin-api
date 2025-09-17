# gin-api


## Run the server
```bash
go run main.go



curl -X POST http://localhost:8080/ask \
  -H "Content-Type: application/json" \
  -d '{"question": "Test question"}'
