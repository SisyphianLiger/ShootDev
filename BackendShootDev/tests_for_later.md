# Basic Tests for first Handler

Going to be used to be setup with the front end but for now:

```bash
curl -v -X POST http://localhost:8080/overthewire -H "Content-Type: application/json" -d '{"data": "test message"}' 
curl -v -X POST http://localhost:8080/overthewire -H "Content-Type: application/json" -d '{"data": ""}'
curl -v -X POST http://localhost:8080/overthewire -H "Content-Type: application/json" -d 'bad message'
```

Should output as followed

1. 200
2. 204
3. 400
