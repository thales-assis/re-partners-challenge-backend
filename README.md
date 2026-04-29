# RE Partners - Challenge - Backend

Challenge for a Backend Software Engineer job.

## Run the Application

To run the Application, simply type `make run` in the terminal (you need to have Docker installed on your machine). After this, you can access: http://localhost:8080 and try and to calculate some packs and items.

## Architecture

For this test, I decided to create a Clean Architecture using these layers:

- HTTP Router (framework)
- Handler
- Use Case
- Domain -> Service & Entity
- Repository

The main application layers (Use Case, Domain) don't know any type of framework that can be used.

If we decided to change de HTTP Framework Server or some ORM for database, the business logic won't change,
it will remain the same. We'll just need to change the application's input and output shells/ports.

## Error Handler

I decided to create some middlewares to catch the errors. This allows our code to be decoupled.

For example, we just write `if err != nil { return err }` without needing to know what type of error we are returning.

The responsibility for knowing lies in the application's shell, as it will inform an appropriate response to the client.

## Libraries

The following libraries were used to develop this application:

Zap - Logger:
- `https://github.com/uber-go/zap`

Viper - Config File Reader:
- `https://github.com/spf13/viper`

Bunrouter - HTTP Framework:
- `https://github.com/uptrace/bunrouter`

Air - Live reload for Go apps:
- `https://github.com/air-verse/air`

Wire - Dependency Injection:
- `https://github.com/google/wire`


After installing the necessary libraries, run these commands:

```bash
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin
```

This will ensure that commands via Makefile can be executed correctly.

## Endpoints

For this test the endpoints are `/packs` (GET, PUT) and `/calculator` (POST).

`GET - /packs` -> Will return all pack sizes in Database.

`PUT - /packs` -> Will allow to update the pack sizes.

`POST - /calculator` -> Will calculate the necessary distribution for the packs (given the amount of items).

You can also check the routes by accessing the application's Swagger in development mode. Execute `make watch` and access:

[Swagger Documentation](http://localhost:3000/re-partners/shipping-service-api/v1/swagger/index.html).

## Running Tests

To run tests, run the following command:

```bash
make tests
```

The tests are configured to run in parallel.

The `-p` flag is a build flag that controls how many packages are compiled and tested at the same time.

The `-parallel` flag is a test binary flag that controls how many individual test functions within a single package can run simultaneously.

## Contact

I am available for comments, criticism, suggestions for improvement, exchange of knowledge, experiences, etc.

**E-mail:** thaleslima-19@hotmail.com

**LinkedIn:** https://www.linkedin.com/in/thales-lima-de-assis/