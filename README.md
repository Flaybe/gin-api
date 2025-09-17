# gin-api
Simple gin REST-api that has a enpoint /ask which uses a ai stub for responding 

## Run the server
```bash
go run cmd/server/main.go
```

## Simple test for /ask
```bash
curl -X POST http://localhost:8080/ask \
  -H "Content-Type: application/json" \
  -d '{"question": "Test question"}'
```

## Heatlh check
```bash
curl http://localhost:8080/health
```
