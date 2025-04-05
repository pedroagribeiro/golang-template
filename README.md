![build branch master](https://github.com/pedroagribeiro/golang-template/actions/workflows/build.yml/badge.svg)
![tests branch master](https://github.com/pedroagribeiro/golang-template/actions/workflows/tests.yml/badge.svg)

<div align="center">
    <img src="src/resources/golang-template.png" alt="golang-scaffold" width="350px">
    <h1>Name of your application</h1>
</div>


## 🚀 Getting started

For data persistence PostgreSQL is used. Therefore, you should either have a
local docker container running the database and guarantee that you have a stable
connection to the remote instance.

This project uses settings configured in environment variables defined in the
`.env` file. Use the `.env.sample` as a starting point.

```bash
cp .env.sample .env
```

In order to get them exported is recommended to setup
[direnv](https://github.com/direnv/direnv) for a terminal based workflow.

## 👨‍💻 Development Environment

The PostgreSQL database can be populated with fake information (used for testing
purposes) by running the following script.

```bash
# Start the development environment
./bin/dev/start

# Stop the fake data database
./bin/dev/stop

# Destroy the development environment (kill & delete containers)
./bin/dev/destroy
```

## 📥 Prerequisites

The following software is required to be installed on your system:

- [GoLang](https://go.dev/)
- [Bruno](https://www.usebruno.com/)
- [Docker](https://www.docker.com/)

## 🔨 Development

These are the commands that you can use to boost your productivity while
developing:

**Compile and start a development server** - you can change the port if you want
to:

```bash
./bin/go/server [port]
```

**Run unit tests**

```bash
./bin/go/test
```

**Format your code**

```bash
./bin/go/format
```

**Lint your code**

```bash
./bin/go/lint
```

**Run api tests**
```bash
./bin/api/test
```

## 📚 Documentation

You can access the api's documentation while the server is up at:

- **http://<server_address>:<port>/swagger/index.html**

## 📦 Deployment

You can deploy a new version of the application with:

```bash
./bin/deploy
```