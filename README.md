# explore-golang

This is a personnal side-project in order to discover, learn, understand and test the Go language.

> My personnal background:
> - I am a full-stack developer: front, back, ops/automation, database, OS/Unix
> - I used in a profesionnal context Java for 8 years, and JavaScript/Node/TypeScript for another 8 years

💡 I share some knowledge and thoughts on Mastodon:
- https://elk.zone/piaille.fr/@jbuget/110786473795556685

## Business use case

The main context in which I place this project is a system that manage accounts and authentification.

I feel this business context covers a wild scope of common features and problems met in a real world (profressionnal or open source).

This project also contains some miscelaneous scripts, tools or attempts that help me to play with Golang.

Things I want to discover:
- #language syntax and idioms
- #ecosystem main blocks (frameworks, libs, tools and services) such as : Web framewok, testing framework, cryptographic libs, logging libs, approches and mechanismes, etc.
- #community the main contributors (devs, writers, participants)
- #practices the good, bad, and ugly parts or technics in terms of: design, security, performance, automation, etc.


## Architecture / infrastructure

- a Web app / REST API in Golang with protected and public routes
- a PostgreSQL database
- maybe some other components such as: a Redis cache, a SMTP/mailing integration, CRON-like tasks, etc.

## Build (localhost)

Pre-requesites :
- Git
- Go CLI
- Docker engine with Compose
- Make

```shell
# Run database
$ docker compose up -d

# Download dependencies
$ go install

# Run app
$ go run main.go
```

## Usage

The example below is for an instance running on `localhost:80`:

```shell
# 1 - Create account
$ curl -v -X POST http://localhost/accounts -d '{"name"="Loulou","email"="loulou@example.org","password"="Abcd1234"}'

# 2 - Get a JWT token
$ curl POST http://localhost/token -d "email=loulou@example.org&password=Abcd1234"

# 3 - Access protected routes
$ curl http://localhost/admin -H "Authorization: Bearer xxx.yyy.zzz"
```


## Deploy

### On Scalingo (french PaaS)

**1/** Configure your account

You MUST have configure your account with a SSL key in order to push.

You MUST have the Scalingo CLI installed and configured (and well logged-in) on your computer. 

You MUST be a member to the Scalingo app `explore-golang`.

**2/** Configure your host

```shell
$ scalingo --app explore-golang git-setup
```

This command adds a Git remote in your `.git/config` file.

**3/** Push to Git remote managed on Scalingo

```shell
$ git push scalingo main
```

> 🔗 See https://explore-golang.osc-fr1.scalingo.io