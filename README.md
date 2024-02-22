# Go Web Service

This is a sample server for a Go web service.

## Getting Started
### Prerequisites
- Go 1.21 or higher
- PostgreSQL
- .env containing : PORT, DB_STRING, SECRET_KEY

### API-Definition
- Redirects all endpoint of /swagger/ to swagger/index.html (httpSwagger.handler of `go get github.com/swaggo/http-swagger/v2`)
#### How it was created
- Generates API definition from Annotations using swaggo/swag/cmd/swag@latest (`go install github.com/swaggo/swag/cmd/swag@latest`)
```
$ > swag init ./ --parseDependency

<!-- This will create Swagger API definition in docs/swagger.json -->
<!-- flag parseDependency is used as some endpoints types are not go-predeclared types, they are customized types from sql library for Null Values Detection -->
```

### API-Testing
- Manual testing can be done through swagger 2.0 API documents.

### Design Patterns
- Chain of Responsibility : Middleware
- Dependency Injection : Singleton database connection + Pgxpool concurrency supported