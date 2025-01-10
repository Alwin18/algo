Powerful CLI tool designed to streamline the process of creating Go projects with standardized structure.

## Why Choose Algo?
- Easy Setup and Installation: Algo simplifies the setup process, making it a breeze to install and get started with your Go projects.

- Pre-established Go Project Structure: Save time and effort by having the entire Go project structure set up automatically. No need to worry about directory layouts or configuration files.

- HTTP Server Configuration Made Easy: Whether you prefer Go's standard library HTTP package, Chi, Gin, Fiber, HttpRouter, Gorilla/mux or Echo, Algo caters to your server setup needs.

- Focus on Your Application Code: With Algo handling the project scaffolding, you can dedicate more time and energy to developing your application logic.

## Project Structure
Project structure created by Algo
```
.root
├── .github
│   └── workflows
│       ├── go-test.yml
│       └── release.yml
├── cmd
│   ├── api
│   │   └── main.go
│   └── web
├── internal
│   ├── config
│   │   ├── app.go
│   │   ├── config.go
│   │   ├── fiber.go
│   │   ├── database.go
│   │   ├── logrus.go
│   │   └── validator.go
│   ├── domain
│   │   ├── User.go
│   │   └── Role.go
│   ├── repository
│   │   ├── user.repository.go
│   │   └── role.repository.go
│   ├── service
│   │   ├── user.service.go
│   │   └── role.service.go
│   ├── dto
│   │   ├── user.dto.go
│   │   └── role.dto.go
│   └── routes
│       ├── handler
│       │   ├── user.handler.go
│       │   └── role.handler.go
│       ├── response.go
│       └── routes.go
├── pkg
│   └── middleware
│       └── auth.go
├── utils
├── .env
├── .env.example
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── JenkinsFile
├── Makefile
└── Readme.md
```
