![GitHub](https://img.shields.io/github/license/costaconrado/services-csm?style=for-the-badge)
## Overview

The purpose of the template is, using the principles of Robert Martin (aka Uncle Bob):
- Show how to organize a project so logic keep extensible and clean
- Show how to use repository pattern to access external data

## Current Project Status
- Backend structured
- Repository pattern structured

## Future Updates
- Add a React + PatternFly frontend provided from static files
- Add integration with OpenShfift resources

## Content
- [Quick start](#quick-start)
- [Project structure](#project-structure)
- [Similar projects](#similar-projects)
- [Useful links](#useful-links)

## Quick start
Local development:
```sh
$ cp .env.example .env
$ make run
```

## Project structure
### `cmd/app/main.go`
First file to run, used to configuration and logger initialization.

### `internal/app/app.go`
called from `cmd/app/main.go` is used to initialize the application itself.

### `config`
Configuration. First, `config.yml` is read, then environment variables overwrite the yaml config if they match.
The config structure is in the `config.go`.
The `env-required: true` tag obliges you to specify a value (either in yaml, or in environment variables).


### `docs`
Swagger documentation auto-generated by [swag](https://github.com/swaggo/swag) library, generated from the comments in `internal/controller/http/v1/router.go`.

### `internal/controller`
Server handler layer (MVC controllers). The template shows 2 servers:
- REST http (Gin framework)

Server routers are written in the same style:
- Handlers are grouped by area of application (by a common basis)
- For each group, its own router structure is created, the methods of which process paths
- The structure of the business logic is injected into the router structure, which will be called by the handlers

### `internal/entity`
Entities of business logic that can be used from any layer and should not depends on any external layer.
There can also be methods, for example, for validation.

### `internal/usecase`
Business logic.
- Methods are grouped in folders by area of application.

#### `internal/usecase/<area>/repo`
A repository is an abstract storage (database) that business logic works with.

It will also allow us to do auto-generation of mocks (for example with [mockery](https://github.com/vektra/mockery)) and easily write unit tests.


## Similar projects
- [https://github.com/evrone/go-clean-template](https://github.com/evrone/go-clean-template)
- [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [https://github.com/zhashkevych/courses-backend](https://github.com/zhashkevych/courses-backend)

## Useful links
- [The Clean Architecture article](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Twelve factors](https://12factor.net/)
